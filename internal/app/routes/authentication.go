package routes

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/handlers"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(rootRouter fiber.Router, path string) {
	authGroup := rootRouter.Group("/auth")

	authGroup.Post(
		"/login",
		middlewares.BodyValidator[*dtos.AuthLoginReqData](),
		handlers.HandleAuthLogin,
	)

	authGroup.Post(
		"/register",
		middlewares.BodyValidator[*dtos.AuthRegisterReqData](),
		handlers.HandleAuthRegister,
	)
}
