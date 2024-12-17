package routes

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app fiber.Router) {
	rootGroup := app.Group("/users")

	rootGroup.Patch("/password", handlers.UserPasswordUpdateHandler)

	RegisterAuthRoutes(rootGroup)
}
