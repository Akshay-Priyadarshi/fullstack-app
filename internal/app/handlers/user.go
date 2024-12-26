package handlers

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/passwords"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserPasswordUpdateHandler(c *fiber.Ctx) error {
	var UserPasswordUpdateReqData dtos.UserPasswordUpdateReqData
	if err := c.BodyParser(&UserPasswordUpdateReqData); err != nil {
		return models.NewApiError("Invalid request body", fiber.StatusBadRequest, nil)
	}
	if err := UserPasswordUpdateReqData.Validate(); err != nil {
		return models.NewApiError(err.Error(), fiber.StatusBadRequest, nil)
	}
	hash, _ := passwords.HashPassword("password")
	loggedInUser := models.User{
		Identity: models.Identity[uuid.UUID]{Id: uuid.New()},
		Email:    "akshay@gmail.com",
		Password: hash,
	}
	if err := loggedInUser.UpdatePassword(UserPasswordUpdateReqData.OldPassword, UserPasswordUpdateReqData.NewPassword); err != nil {
		return models.NewApiError(err.Error(), fiber.StatusBadRequest, nil)
	}
	UserResData := dtos.UserResData{
		Id:    loggedInUser.Identity.Id.String(),
		Email: loggedInUser.Email,
	}
	apiResponse := models.NewApiResponse[dtos.UserResData](
		"password updated successfully",
		fiber.StatusOK,
		&UserResData,
		nil,
	)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}
