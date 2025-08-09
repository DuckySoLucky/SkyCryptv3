package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/constants"
	"skycrypt/src/stats"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GardenHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	profileId := c.Params("profileId")
	garden, err := api.GetGarden(profileId)
	if err != nil {
		c.Status(400)
		return c.JSON(constants.InvalidUserError)
	}

	fmt.Printf("Returning /api/garden/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"garden": stats.GetGarden(garden),
	})
}
