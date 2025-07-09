package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/stats"
	"time"

	"github.com/gofiber/fiber/v2"
)

func InventoryHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")
	if len(profileId) > 0 && profileId[0] == '/' {
		profileId = profileId[1:]
	}

	mowojang, err := api.ResolvePlayer(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to resolve player: %v", err),
		})
	}

	profile, err := api.GetProfile(mowojang.UUID, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	userProfileValue := profile.Members[mowojang.UUID]
	userProfile := &userProfileValue

	items, err := stats.GetItems(userProfile, profile.ProfileID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get items: %v", err),
		})
	}

	fmt.Printf("Returning /api/inventory/%s in %s\n", uuid, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"items": items,
	})
}
