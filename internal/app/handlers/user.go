package handlers

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/responses"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services"
	"github.com/gofiber/fiber/v2"
)

func HandleUserPasswordUpdate(c *fiber.Ctx) error {
	authUser := c.Locals("authUser").(*dtos.UserResData)
	passwordUpdateReqDto := c.Locals("validatedDto").(*dtos.UserPasswordUpdateReqData)
	userService := services.UserService{}
	userResData, err := userService.UpdatePassword(authUser.Id, passwordUpdateReqDto)
	if err != nil {
		return err
	}
	apiResponse := responses.NewOkResponse(
		"password updated successfully",
		userResData,
	)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}

func HandleUserGetById(c *fiber.Ctx) error {
	userId := c.Params("id")
	userService := services.UserService{}
	userResData, err := userService.GetById(userId)
	if err != nil {
		return err
	}
	apiResponse := responses.NewOkResponse(
		"user fetched successfully",
		userResData,
	)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}
