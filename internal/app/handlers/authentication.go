package handlers

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/responses"
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
func HandleAuthRegister(c *fiber.Ctx) error {
	registerReqDto := c.Locals("validatedDto").(*dtos.AuthRegisterReqData)
	authService := services.AuthService{}
	authResData, err := authService.Register(registerReqDto)
	if err != nil {
		return err
	}
	apiResponse := responses.NewOkResponse("user registered successfully", authResData)
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
func HandleAuthLogin(c *fiber.Ctx) error {
	loginReqDto := c.Locals("validatedDto").(*dtos.AuthLoginReqData)
	authService := services.AuthService{}
	authResData, err := authService.Login(loginReqDto)
	if err != nil {
		return err
	}
	apiResponse := responses.NewOkResponse("user logged in successfully", authResData)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}
