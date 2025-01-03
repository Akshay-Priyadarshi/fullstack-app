package services

import (
	"fmt"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/repositories"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/passwords"
	"github.com/google/uuid"
)

type UserService struct{}

func (us *UserService) UpdatePassword(id string, userPasswordUpdateDto *dtos.UserPasswordUpdateReqData) (*dtos.UserResData, error) {
	userRepo := &repositories.UserRepository{}
	user, err := userRepo.GetById(uuid.MustParse(id))
	if err != nil {
		return nil, New400ServiceError(err, fmt.Sprintf("user with id '%s' is not found", id))
	}
	err = user.ComparePassword(userPasswordUpdateDto.OldPassword)
	if err != nil {
		return nil, New400ServiceError(err, "old password is incorrect")
	}
	hashedPassword, err := passwords.HashPassword(userPasswordUpdateDto.NewPassword)
	if err != nil {
		return nil, New500ServiceError(err)
	}
	err = userRepo.UpdatePassword(user.Id, hashedPassword)
	if err != nil {
		return nil, New400ServiceError(err, "failed to update password")
	}
	userResData := dtos.UserResData{
		Id:    user.Id.String(),
		Email: user.Email,
	}
	return &userResData, nil
}

func (us *UserService) GetById(id string) (*dtos.UserResData, error) {
	userRepo := &repositories.UserRepository{}
	parsedUUID := uuid.MustParse(id)
	user, err := userRepo.GetById(parsedUUID)
	if err != nil {
		return nil, New400ServiceError(err, fmt.Sprintf("user with id '%s' is not found", id))
	}
	userResData := dtos.UserResData{
		Id:    user.Id.String(),
		Email: user.Email,
	}
	return &userResData, nil
}
