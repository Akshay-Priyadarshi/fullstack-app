package middlewares

import (
	"strings"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/services"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/jwts"
	"github.com/gofiber/fiber/v2"
)

func JwtExtractor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return services.New401ServiceError(nil, "missing authorization header")
		}
		authToken := strings.Split(authHeader, " ")[1]
		jwtClaims, err := jwts.VerifyJWT(authToken, server.AppServer.Config.JwtSecret)
		if err != nil {
			return err
		}
		sub, err := jwtClaims.GetSubject()
		if err != nil {
			return services.New401ServiceError(err, "failed to get subject from token")
		}
		userService := services.UserService{}
		userResData, err := userService.GetById(sub)
		if err != nil {
			return services.New401ServiceError(err, "user not found")
		}
		c.Locals("authUser", userResData)
		return c.Next()
	}
}
