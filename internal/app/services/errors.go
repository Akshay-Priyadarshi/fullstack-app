package services

import (
	"fmt"
	"strings"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/responses"
	"github.com/gofiber/fiber/v2"
)

type ServiceError struct {
	InternalError  error
	StatusCode     int
	Message        string
	AdditionalInfo *map[string]interface{}
}

func (se *ServiceError) Error() string {
	return fmt.Sprintf("%s: %s", se.InternalError.Error(), se.Message)
}

func (se *ServiceError) ToApiResponse() responses.ApiResponse[any] {
	return responses.NewResponse[any](se.Message, se.StatusCode, nil, se.AdditionalInfo)
}

func NewServiceError(internalError error, statusCode int, message string, additionalInfo *map[string]interface{}) *ServiceError {
	return &ServiceError{
		InternalError:  internalError,
		Message:        strings.ToLower(message),
		StatusCode:     statusCode,
		AdditionalInfo: additionalInfo,
	}
}

func New400ServiceError(internalError error, message string) *ServiceError {
	code := fiber.ErrBadRequest.Code
	if message == "" {
		message = fiber.ErrBadRequest.Message
	}
	return NewServiceError(internalError, code, message, nil)
}

func New401ServiceError(internalError error, message string) *ServiceError {
	code := fiber.ErrUnauthorized.Code
	if message == "" {
		message = fiber.ErrUnauthorized.Message
	}
	return NewServiceError(internalError, code, message, nil)
}

func New403ServiceError(internalError error, message string) *ServiceError {
	code := fiber.ErrForbidden.Code
	if message == "" {
		message = fiber.ErrForbidden.Message
	}
	return NewServiceError(internalError, code, message, nil)
}

func New404ServiceError(internalError error, message string) *ServiceError {
	code := fiber.ErrNotFound.Code
	if message == "" {
		message = fiber.ErrNotFound.Message
	}
	return NewServiceError(internalError, code, message, nil)
}

func New422ServiceError(internalError error, message string, additionalInfo *map[string]any) *ServiceError {
	code := fiber.ErrUnprocessableEntity.Code
	if message == "" {
		message = fiber.ErrUnprocessableEntity.Message
	}
	return NewServiceError(internalError, code, message, additionalInfo)
}

func New500ServiceError(internalError error) *ServiceError {
	code := fiber.ErrInternalServerError.Code
	message := fiber.ErrInternalServerError.Message
	return NewServiceError(internalError, code, message, nil)
}
