package swagger

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models/dtos"
)

// Register
type AuthenticationRegisterApiResponse struct {
	Success        bool                                     `json:"success"`
	Message        string                                   `json:"message"`
	Data           *dtos.AuthenticationRegisterResponseData `json:"data"`
	StatusCode     int                                      `json:"statusCode"`
	AdditionalInfo *map[string]interface{}                  `json:"additionalInfo"`
}

// Login
type AuthenticationLoginApiResponse struct {
	Success        bool                                  `json:"success"`
	Message        string                                `json:"message"`
	Data           *dtos.AuthenticationLoginResponseData `json:"data"`
	StatusCode     int                                   `json:"statusCode"`
	AdditionalInfo *map[string]interface{}               `json:"additionalInfo"`
}
