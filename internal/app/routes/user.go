package routes

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app fiber.Router, path string) {
	rootGroup := app.Group(path)

	rootGroup.Patch(
		"/password",
		middlewares.JwtExtractor(),
		middlewares.BodyValidator[dtos.UserPasswordUpdateReqData](),
		handlers.HandleUserPasswordUpdate,
	)

	rootGroup.Get(
		"/:id",
		middlewares.JwtExtractor(),
		handlers.HandleUserGetById,
	)
}
