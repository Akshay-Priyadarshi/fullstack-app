package middlewares

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos/abstractions"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

func ReqBodyValidator[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto T
		err := c.BodyParser(&dto)
		if err != nil {
			return models.New400ApiError("Invalid request body")
		}
		if validatable, ok := any(&dto).(abstractions.Validatable); ok {
			if err := validatable.Validate(); err != nil {
				return models.New400ApiError(err.Error())
			}
		}
		c.Locals("validatedDto", &dto)
		return c.Next()
	}
}
