package services

import (
	"fmt"
	"time"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/repositories"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/jwts"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/passwords"
)

type AuthService struct{}

func (a *AuthService) Login(loginReqDto *dtos.AuthLoginReqData) (*dtos.AuthResData, error) {
	userRepo := repositories.UserRepository{}
	user, err := userRepo.GetByEmail(loginReqDto.Email)
	if err != nil {
		return nil, New400ServiceError(err, fmt.Sprintf("user with email '%s' does not exist, try registering", loginReqDto.Email))
	}
	err = passwords.ComparePassword(loginReqDto.Password, user.Password)
	if err != nil {
		return nil, New400ServiceError(err, "invalid email password combination")
	}
	tokenString, err := jwts.GenerateJWT(user.Id.String(), time.Duration(time.Hour*1), server.AppServer.Config.JwtSecret)
	if err != nil {
		return nil, New500ServiceError(err)
	}
	loginResData := dtos.AuthResData{
		Id:    user.Id,
		Email: user.Email,
		Token: tokenString,
	}
	return &loginResData, nil
}

func (a *AuthService) Register(registerReqDto *dtos.AuthRegisterReqData) (*dtos.AuthResData, error) {
	user, err := registerReqDto.ToUser()
	if err != nil {
		return nil, New500ServiceError(err)
	}
	userRepo := repositories.UserRepository{}
	err = userRepo.Create(user)
	if err != nil {
		return nil, New400ServiceError(err, fmt.Sprintf("user with email '%s' already exists", user.Email))
	}
	tokenString, err := jwts.GenerateJWT(user.Id.String(), time.Duration(time.Hour*1), server.AppServer.Config.JwtSecret)
	if err != nil {
		return nil, New500ServiceError(err)
	}
	registerResData := dtos.AuthResData{
		Id:    user.Id,
		Email: user.Email,
		Token: tokenString,
	}
	return &registerResData, nil
}
