package handlers

import (
	"errors"
	"log/slog"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/responses"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services"
	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, err error) error {
	slog.Error(err.Error())
	var serviceError *services.ServiceError
	var fiberError *fiber.Error
	switch {
	case errors.As(err, &serviceError):
		apiResponse := serviceError.ToApiResponse()
		return c.Status(apiResponse.StatusCode).JSON(apiResponse)
	case errors.As(err, &fiberError):
		apiResponse := responses.NewResponse[any](fiberError.Message, fiberError.Code, nil, nil)
		return c.Status(apiResponse.StatusCode).JSON(apiResponse)
	default:
		defaultError := services.New500ServiceError(nil)
		apiResponse := defaultError.ToApiResponse()
		return c.Status(apiResponse.StatusCode).JSON(apiResponse)
	}
}
