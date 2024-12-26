package main

import (
	"encoding/json"
	"os"

	_ "github.com/Akshay-Priyadarshi/fullstack-app/api/openapi"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/routes"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/Akshay-Priyadarshi/fullstack-app/web"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

// @title Api Template
// @version 1.0
// @description This is a template for a RESTful API using Fiber.
// @BasePath /
// @schemes http
// @host localhost:8080
// @produce json
// @consumes json
func main() {
	fiberApp := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: handlers.ErrorHandler,
	})

	server.AppServer = server.New(fiberApp, &server.Config{
		Port:      os.Getenv("PORT"),
		DBString:  os.Getenv("DB_CONN_STRING"),
		ApiPath:   os.Getenv("API_PATH"),
		WebPath:   os.Getenv("WEB_PATH"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	})

	// Register api routes
	routes.RegisterRoutes(server.AppServer.App, server.AppServer.Config.ApiPath)

	// Register web routes
	web.RegisterRoutes(server.AppServer.App, server.AppServer.Config.WebPath)

	server.AppServer.Start()
}
