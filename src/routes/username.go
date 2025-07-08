package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/constants"
	"skycrypt/src/utility"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UsernameHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	if !utility.IsUUID(uuid) {
		c.Status(400)
		return c.JSON(constants.InvalidUserError)
	}

	username, err := api.GetUsername(uuid)
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
