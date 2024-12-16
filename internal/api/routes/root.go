package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRootRoutes(app *fiber.App) {
	rootGroup := app.Group("/api")

	RegisterAuthRoutes(rootGroup)
	RegisterUserRoutes(rootGroup)

	rootGroup.All("/*", func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
