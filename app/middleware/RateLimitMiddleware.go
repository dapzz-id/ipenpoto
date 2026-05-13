package middleware

import (
	"fmt"
	"strconv"
	"time"

	"ipenpoto/config"

	"github.com/gofiber/fiber/v2"
)

func RateLimit(maxRequests int, windowSeconds int64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		key := fmt.Sprintf("ratelimit:%s", ip)

		val, err := config.RDB.Get(config.Ctx, key).Result()
		var count int

		if err != nil {
			count = 0
		} else {
			count, _ = strconv.Atoi(val)
		}

		if count >= maxRequests {
			return c.Status(429).JSON(fiber.Map{
				"message": "Too many requests - Rate limit exceeded",
			})
		}

		config.RDB.Incr(config.Ctx, key)
		if count == 0 {
			config.RDB.Expire(config.Ctx, key, time.Duration(windowSeconds)*time.Second)
		}

		c.Set("X-RateLimit-Remaining", fmt.Sprintf("%d", maxRequests-count-1))
		c.Set("X-RateLimit-Limit", fmt.Sprintf("%d", maxRequests))

		return c.Next()
	}
}
