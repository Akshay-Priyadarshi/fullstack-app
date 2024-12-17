package dtos

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/pkg/passwords"
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
	passwordHash, err := passwords.HashPassword(data.Password)
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

type AuthenticationRegisterApiResponse struct {
	Success        bool                                `json:"success"`
	Message        string                              `json:"message"`
	Data           *AuthenticationRegisterResponseData `json:"data"`
	StatusCode     int                                 `json:"statusCode"`
	AdditionalInfo *map[string]interface{}             `json:"additionalInfo"`
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

type AuthenticationLoginApiResponse struct {
	Success        bool                             `json:"success"`
	Message        string                           `json:"message"`
	Data           *AuthenticationLoginResponseData `json:"data"`
	StatusCode     int                              `json:"statusCode"`
	AdditionalInfo *map[string]interface{}          `json:"additionalInfo"`
}
