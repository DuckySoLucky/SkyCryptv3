package handlers

import "github.com/gofiber/fiber/v2"

// Pre-allocate response to avoid JSON marshaling overhead
var helloResponse = fiber.Map{
	"message": "Hello World",
	"status":  "success",
}

// HelloHandler handles the hello endpoint
func HelloHandler(c *fiber.Ctx) error {
	return c.JSON(helloResponse)
}
