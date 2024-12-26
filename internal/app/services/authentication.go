package services

import (
	"time"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/repositories"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/jwts"
)

type AuthService struct{}

func (a *AuthService) Login(loginReqDto *dtos.AuthLoginReqData) (*dtos.AuthResData, error) {
	userRepo := repositories.UserRepository{}
	userPtr, err := userRepo.GetByEmail(loginReqDto.Email)
	if err != nil {
		return nil, err
	}
	tokenString, err := jwts.GenerateJWT(userPtr.Id.String(), time.Duration(time.Hour*1), server.AppServer.Config.JwtSecret)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	userRepo := repositories.UserRepository{}
	err = userRepo.Create(userPtr)
	if err != nil {
		return nil, err
	}
	tokenString, err := jwts.GenerateJWT(userPtr.Id.String(), time.Duration(time.Hour*1), server.AppServer.Config.JwtSecret)
	if err != nil {
		return nil, err
	}
	registerResData := dtos.AuthResData{
		Id:    userPtr.Id,
		Email: userPtr.Email,
		Token: tokenString,
	}
	return &registerResData, nil
}
