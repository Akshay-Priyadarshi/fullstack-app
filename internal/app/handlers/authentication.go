package handlers

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services"
	"github.com/gofiber/fiber/v2"
)

// RegisterHandler godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dtos.AuthRegisterReqData true "User registration details"
// @Success 200 {object} dtos.AuthRegisterApiRes
// @Router /auth/register [post]
func AuthRegisterHandler(c *fiber.Ctx) error {
	registerReqDto := c.Locals("validatedDto").(*dtos.AuthRegisterReqData)
	authService := services.AuthService{}
	authResDataPtr, err := authService.Register(registerReqDto)
	if err != nil {
		return models.NewBadRequestError(err.Error())
	}
	apiResponse := models.NewOkApiResponse("user registered successfully", authResDataPtr)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}

// LoginHandler godoc
// @Summary Login a user
// @Description Login a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dtos.AuthLoginReqData true "User login details"
// @Success 200 {object} dtos.AuthLoginApiRes
// @Router /auth/login [post]
func AuthLoginHandler(c *fiber.Ctx) error {
	loginReqDto := c.Locals("validatedDto").(*dtos.AuthLoginReqData)
	authService := services.AuthService{}
	authResDataPtr, err := authService.Login(loginReqDto)
	if err != nil {
		return models.NewBadRequestError(err.Error())
	}
	apiResponse := models.NewOkApiResponse("user logged in successfully", authResDataPtr)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}
