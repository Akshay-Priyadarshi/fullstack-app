package dtos

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/passwords"
	"github.com/google/uuid"
)

// Register
type AuthRegisterReqData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (data *AuthRegisterReqData) Validate() error {
	err := server.AppServer.Validate.Struct(data)
	return err
}

func (data *AuthRegisterReqData) ToUser() (*models.User, error) {
	passwordHash, err := passwords.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	userPtr := &models.User{
		Identity: models.Identity[uuid.UUID]{Id: uuid.New()},
		Email:    data.Email,
		Password: passwordHash,
	}
	return userPtr, nil
}

type AuthResData struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

type AuthRegisterApiRes struct {
	Success        bool                    `json:"success"`
	Message        string                  `json:"message"`
	Data           *AuthResData            `json:"data"`
	StatusCode     int                     `json:"statusCode"`
	AdditionalInfo *map[string]interface{} `json:"additionalInfo"`
}

// Login
type AuthLoginReqData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (data *AuthLoginReqData) Validate() error {
	err := server.AppServer.Validate.Struct(data)
	return err
}

type AuthLoginApiRes struct {
	Success        bool                    `json:"success"`
	Message        string                  `json:"message"`
	Data           *AuthResData            `json:"data"`
	StatusCode     int                     `json:"statusCode"`
	AdditionalInfo *map[string]interface{} `json:"additionalInfo"`
}
