package routes

import (
	"fmt"
	"skycrypt/src/api"
	"time"

	"github.com/gofiber/fiber/v2"
)

func MuseumHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	profileId := c.Params("profileId")
	museum, err := api.GetMuseum(profileId)
	fmt.Print(err)
	if err != nil {
		return c.JSON(fiber.Map{})
	}

	fmt.Printf("Returning /api/museum/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"museum": museum,
	})
}
