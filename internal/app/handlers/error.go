package handlers

import (
	"log"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	serviceerrors "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services/service_errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Printf("Error: %v", err)
	if apiError, ok := err.(*models.ApiError); ok {
		apiErrorServerResponse := apiError.GetApiResponse()
		return c.Status(apiErrorServerResponse.StatusCode).JSON(apiErrorServerResponse)
	} else if serviceError, ok := err.(serviceerrors.ServiceError); ok {
		apiError := models.NewBadRequestError(serviceError.Error())
		apiErrorServerResponse := apiError.GetApiResponse()
		return c.Status(apiErrorServerResponse.StatusCode).JSON(apiErrorServerResponse)
	} else if fiberError, ok := err.(*fiber.Error); ok {
		fiberErrorServerResponse := models.NewApiResponse[any](fiberError.Message, fiberError.Code, nil, nil)
		return c.Status(fiberErrorServerResponse.StatusCode).JSON(fiberErrorServerResponse)
	} else {
		defaultError := models.NewApiError(fiber.ErrInternalServerError.Message, fiber.StatusInternalServerError, nil)
		defaultErrorApiResponse := defaultError.GetApiResponse()
		return c.Status(defaultErrorApiResponse.StatusCode).JSON(defaultErrorApiResponse)
	}
}
