package routes

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(rootRouter fiber.Router, path string) {
	authGroup := rootRouter.Group("/auth")

	authGroup.Post("/login", middlewares.ReqBodyValidator[dtos.AuthLoginReqData](), handlers.AuthLoginHandler)

	authGroup.Post("/register", middlewares.ReqBodyValidator[dtos.AuthRegisterReqData](), handlers.AuthRegisterHandler)
}
