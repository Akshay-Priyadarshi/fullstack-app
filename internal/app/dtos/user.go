package dtos

import "github.com/go-playground/validator/v10"

type UserPasswordUpdateReqData struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

func (data *UserPasswordUpdateReqData) Validate() error {
	validate := validator.New()
	return validate.Struct(data)
}

type UserResData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
