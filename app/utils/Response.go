package utils

import "github.com/gofiber/fiber/v2"

func Success(
	c *fiber.Ctx,
	messageOrCode interface{},
	dataOrMessage interface{},
	optionalData ...interface{},
) error {
	var code int = 200
	var message string
	var data interface{}

	switch v := messageOrCode.(type) {
	case int:
		code = v
		if msg, ok := dataOrMessage.(string); ok {
			message = msg
		}
		if len(optionalData) > 0 {
			data = optionalData[0]
		}
	case string:
		message = v
		data = dataOrMessage
	}

	return c.Status(code).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func Error(
	c *fiber.Ctx,
	code int,
	message string,
	err interface{},
) error {

	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"error":   err,
	})
}
