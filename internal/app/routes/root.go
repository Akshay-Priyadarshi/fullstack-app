package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func RegisterRoutes(app *fiber.App, path string) {
	rootGroup := app.Group(path)

	rootGroup.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/live",
		ReadinessEndpoint: "/ready",
	}))

	RegisterAuthRoutes(rootGroup, "/auth")
	RegisterUserRoutes(rootGroup, "/users")

	rootGroup.All("*", func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
