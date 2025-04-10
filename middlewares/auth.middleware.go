package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/task-manager/services"
	"github.com/task-manager/types"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing Authorization header.",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid Authorization header format.",
			})
		}

		tokenString := parts[1]

		claims, err := services.VerifyToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired token.",
			})
		}

		user := types.CurrentUser{
			ID:    uint(claims["id"].(float64)),
			Email: claims["email"].(string),
		}

		c.Locals("currentUser", user)

		return c.Next()
	}
}
