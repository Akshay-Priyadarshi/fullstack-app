package models

import "github.com/gofiber/fiber/v2"

type ApiError struct {
	Message        string                  `json:"message"`
	StatusCode     int                     `json:"statusCode"`
	AdditionalInfo *map[string]interface{} `json:"additionalInfo"`
}

func (apiError *ApiError) Error() string {
	return apiError.Message
}

func (apiError *ApiError) GetApiResponse() ApiResponse[any] {
	return NewApiResponse[any](
		apiError.Message,
		apiError.StatusCode,
		nil,
		apiError.AdditionalInfo,
	)
}

func NewApiError(message string, statusCode int, additionalInfo *map[string]interface{}) *ApiError {
	if statusCode == 0 {
		statusCode = 500
	}
	if message == "" {
		message = "Internal Server Error"
	}
	return &ApiError{
		Message:        message,
		StatusCode:     statusCode,
		AdditionalInfo: additionalInfo,
	}
}

func New400ApiError(message string) *ApiError {
	return NewApiError(message, fiber.StatusBadRequest, nil)
}
