package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SecurityLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		statusCode := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()
		ip := c.IP()
		userID := c.Locals("user_id")

		if statusCode >= 400 {
			log.Printf(
				"[SECURITY] %s %s - Status:%d IP:%s UserID:%v Duration:%v",
				method,
				path,
				statusCode,
				ip,
				userID,
				duration,
			)
		}

		if statusCode == 401 || statusCode == 403 {
			log.Printf(
				"[AUTH_FAILURE] %s %s - Status:%d IP:%s UserID:%v",
				method,
				path,
				statusCode,
				ip,
				userID,
			)
		}

		if statusCode == 429 {
			log.Printf("[RATE_LIMIT] IP:%s Path:%s", ip, path)
		}

		fmt.Printf(
			"[%s] %s %s %d %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			statusCode,
			duration,
		)

		return nil
	}
}
