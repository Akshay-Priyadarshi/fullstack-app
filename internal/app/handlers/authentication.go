package handlers

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models/dtos"
	"github.com/gofiber/fiber/v2"
)

// RegisterHandler godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dtos.AuthenticationRegisterRequestData true "User registration details"
// @Success 200 {object} dtos.AuthenticationRegisterApiResponse
// @Router /authentication/register [post]
func RegisterHandler(c *fiber.Ctx) error {
	// Create a registerRequestDto to parse the incoming JSON body
	var registerRequestDto dtos.AuthenticationRegisterRequestData

	if err := c.BodyParser(&registerRequestDto); err != nil {
		return models.NewApiError("Invalid request body", fiber.StatusBadRequest, nil)
	}

	if err := registerRequestDto.Validate(); err != nil {
		return models.NewApiError(err.Error(), fiber.StatusBadRequest, nil)
	}

	user := registerRequestDto.ToUser()

	registerResponseData := dtos.AuthenticationRegisterResponseData{
		Id:    user.Id,
		Email: user.Email,
	}

	// Return the user (ensure sensitive fields like passwords are excluded)
	apiResponse := models.NewApiResponse(
		"User created successfully",
		fiber.StatusOK,
		&registerResponseData,
		&map[string]interface{}{},
	)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}

// LoginHandler godoc
// @Summary Login a user
// @Description Login a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dtos.AuthenticationLoginRequestData true "User login details"
// @Success 200 {object} dtos.AuthenticationLoginApiResponse
// @Router /authentication/login [post]
func LoginHandler(c *fiber.Ctx) error {
	// Return login route string
	loginRoute := "Login route"
	apiResponse := models.NewApiResponse[string](
		"Login route",
		fiber.StatusOK,
		&loginRoute,
		&map[string]interface{}{},
	)
	return c.Status(apiResponse.StatusCode).JSON(apiResponse)
}
