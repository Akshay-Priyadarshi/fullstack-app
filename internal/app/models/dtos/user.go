package dtos

import "github.com/go-playground/validator/v10"

type UserPasswordUpdateRequestData struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

func (data *UserPasswordUpdateRequestData) Validate() error {
	validate := validator.New()
	return validate.Struct(data)
}

type UserResponseData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
