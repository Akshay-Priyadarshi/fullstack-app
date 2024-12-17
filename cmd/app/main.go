package main

import (
	"encoding/json"

	_ "github.com/Akshay-Priyadarshi/fullstack-app/api/openapi"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/routes"
	"github.com/Akshay-Priyadarshi/fullstack-app/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Api Template
// @version 1.0
// @description This is a template for a RESTful API using Fiber.
// @BasePath /
// @schemes http
// @host localhost:8080
// @produce json
// @consumes json
func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: handlers.ErrorHandler,
	})

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Enable security headers
	app.Use(helmet.New())

	// Enable logging
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint:  "/live",
		ReadinessEndpoint: "/ready",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.RegisterRootRoutes(app)

	web.RegisterClientRoutes(app)

	port := "8080"
	println("Server is starting at: http://localhost:" + port)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
