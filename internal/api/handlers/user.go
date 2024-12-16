package handlers

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/api/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/api/models/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserPasswordUpdateHandler(c *fiber.Ctx) error {
	var userPasswordUpdateRequestData dtos.UserPasswordUpdateRequestData

	if err := c.BodyParser(&userPasswordUpdateRequestData); err != nil {
		return models.NewApiError("Invalid request body", fiber.StatusBadRequest, nil)
	}

	if err := userPasswordUpdateRequestData.Validate(); err != nil {
		return models.NewApiError(err.Error(), fiber.StatusBadRequest, nil)
	}

	hash, _ := services.HashPassword("password")
	loggedInUser := models.User{
		Identity: models.Identity[uuid.UUID]{Id: uuid.New()},
		Email:    "akshay@gmail.com",
		Password: hash,
	}

	if err := loggedInUser.UpdatePassword(userPasswordUpdateRequestData.OldPassword, userPasswordUpdateRequestData.NewPassword); err != nil {
		return models.NewApiError(err.Error(), fiber.StatusBadRequest, nil)
	}

	userResponseData := dtos.UserResponseData{
		Id:    loggedInUser.Identity.Id.String(),
		Email: loggedInUser.Email,
	}

	apiResponse := models.NewApiResponse[dtos.UserResponseData](
		"password updated successfully",
		fiber.StatusOK,
		&userResponseData,
		nil,
	)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}
