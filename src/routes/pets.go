package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/stats"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PetsHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")

	profile, err := api.GetProfile(uuid, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	userProfileValue := profile.Members[uuid]
	userProfile := &userProfileValue

	pets, err := stats.GetPets(userProfile, profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get pets: %v", err),
		})
	}

	fmt.Printf("Returning /api/pets/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"pets": pets,
	})
}
