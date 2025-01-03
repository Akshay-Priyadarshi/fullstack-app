package main

import (
	"encoding/json"
	"flag"
	"os"

	_ "github.com/Akshay-Priyadarshi/fullstack-app/api/openapi"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/constants"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/routes"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server/initialisations"
	"github.com/Akshay-Priyadarshi/fullstack-app/web"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	GoEnvFromEnvironment := os.Getenv("GO_ENV")
	if GoEnvFromEnvironment == "" {
		os.Setenv("GO_ENV", constants.EnvProduction)
	}
	if GoEnvFromEnvironment == constants.EnvDevelopment {
		godotenv.Load(".env")
	}
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
	env := flag.String("env", os.Getenv("GO_ENV"), "Set the environment (development, staging, production)")
	flag.Parse()
	logger := initialisations.InitLogger(*env)

	fiberApp := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: handlers.HandleError,
	})

	appConfig := &server.Config{
		Port:      os.Getenv("PORT"),
		DBString:  os.Getenv("DB_CONN_STRING"),
		ApiPath:   os.Getenv("API_PATH"),
		WebPath:   os.Getenv("WEB_PATH"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		Env:       *env,
	}

	server.AppServer = server.New(fiberApp, appConfig, logger)

	server.AppServer.Configure()

	// Register api routes
	routes.RegisterRoutes(server.AppServer.App, server.AppServer.Config.ApiPath)

	// Register web routes
	web.RegisterRoutes(server.AppServer.App, server.AppServer.Config.WebPath)

	server.AppServer.Start()
}
