package routes

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(rootRouter fiber.Router) {
	authGroup := rootRouter.Group("/authentication")

	authGroup.Post("/login", handlers.LoginHandler)

	authGroup.Post("/register", handlers.RegisterHandler)
}
