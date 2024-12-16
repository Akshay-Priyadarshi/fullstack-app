package dtos

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/api/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/api/services"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Register
type AuthenticationRegisterRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (data *AuthenticationRegisterRequestData) Validate() error {
	validate := validator.New()
	return validate.Struct(data)
}

func (data *AuthenticationRegisterRequestData) ToUser() *models.User {
	passwordHash, err := services.HashPassword(data.Password)
	if err != nil {
		panic(err)
	}
	return &models.User{
		Identity: models.Identity[uuid.UUID]{Id: uuid.New()},
		Email:    data.Email,
		Password: passwordHash,
	}
}

type AuthenticationRegisterResponseData struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

// Login
type AuthenticationLoginRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationLoginResponseData struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}
