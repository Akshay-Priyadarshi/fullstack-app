package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/repositories"
	repositoryerrors "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/repositories/repository_errors"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	serviceerrors "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services/service_errors"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services/service_errors/users"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/jwts"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/passwords"
)

type AuthService struct{}

func (a *AuthService) Login(loginReqDto *dtos.AuthLoginReqData) (*dtos.AuthResData, error) {
	userRepo := repositories.UserRepository{}
	userPtr, err := userRepo.GetByEmail(loginReqDto.Email)
	if err != nil {
		if errors.Is(err, repositoryerrors.ErrRepositoryNoRows) {
			serviceerrors.New(users.ErrUserEmailAlreadyExists{Email: loginReqDto.Email})
		}
		return nil, serviceerrors.New(errors.New("failed to get user by email"))
	}
	err = passwords.ComparePassword(loginReqDto.Password, userPtr.Password)
	if err != nil {
		return nil, serviceerrors.New(fmt.Errorf("invalid email password combination"))
	}
	tokenString, err := jwts.GenerateJWT(userPtr.Id.String(), time.Duration(time.Hour*1), server.AppServer.Config.JwtSecret)
	if err != nil {
		return nil, fmt.Errorf("auth service: %w", err)
	}
	loginResData := dtos.AuthResData{
		Id:    userPtr.Id,
		Email: userPtr.Email,
		Token: tokenString,
	}
	return &loginResData, nil
}

func (a *AuthService) Register(registerReqDto *dtos.AuthRegisterReqData) (*dtos.AuthResData, error) {
	userPtr, err := registerReqDto.ToUser()
	if err != nil {
		return nil, fmt.Errorf("auth service: %w", err)
	}
	userRepo := repositories.UserRepository{}
	err = userRepo.Create(userPtr)
	if err != nil {
		if errors.Is(err, repositoryerrors.ErrRepositoryDuplicateValue) {
			return nil, serviceerrors.New(users.ErrUserEmailAlreadyExists{Email: userPtr.Email})
		}
		return nil, serviceerrors.New(errors.New("failed to create user"))
	}
	tokenString, err := jwts.GenerateJWT(userPtr.Id.String(), time.Duration(time.Hour*1), server.AppServer.Config.JwtSecret)
	if err != nil {
		return nil, fmt.Errorf("auth service: %w", err)
	}
	registerResData := dtos.AuthResData{
		Id:    userPtr.Id,
		Email: userPtr.Email,
		Token: tokenString,
	}
	return &registerResData, nil
}
