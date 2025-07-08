package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/constants"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UUIDHandler(c *fiber.Ctx) error {
	timeNow := time.Now()
	username := c.Params("username")
	if username == "" || len(username) < 3 || len(username) > 16 {
		c.Status(400)
		return c.JSON(constants.InvalidUserError)
	}

	uuid, err := api.GetUUID(username)
	if err != nil {
		c.Status(400)
		return c.JSON(constants.InvalidUserError)
	}

	fmt.Printf("Returning /api/username/%s in %s\n", username, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"username": username,
		"uuid":     uuid,
	})
}
