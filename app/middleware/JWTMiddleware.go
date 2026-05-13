package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized - Missing token",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized - Invalid token format",
			})
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&jwt.MapClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized - Invalid token",
			})
		}

		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized - Invalid claims",
			})
		}

		userID, ok := (*claims)["user_id"]
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized - Missing user_id",
			})
		}

		role, ok := (*claims)["role"]
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized - Missing role",
			})
		}

		c.Locals("user_id", userID)
		c.Locals("role", role)

		return c.Next()
	}
}
