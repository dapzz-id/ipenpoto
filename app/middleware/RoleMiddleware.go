package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("role").(string)
		if !ok || userRole == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized - Invalid role",
			})
		}

		for _, role := range roles {
			if role == userRole {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden - Insufficient permissions",
		})
	}
}
