package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/models"
	"skycrypt/src/stats"
	statsItems "skycrypt/src/stats/items"
	"strings"

	"time"

	"github.com/gofiber/fiber/v2"
)

func InventoryHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")
	inventoryId := c.Params("inventoryId")
	profile, err := api.GetProfile(uuid, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	userProfileValue := profile.Members[uuid]
	userProfile := &userProfileValue

	items, err := stats.GetItems(userProfile, profile.ProfileID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get items: %v", err),
		})
	}

	if inventoryId == "search" {
		searchString := c.Params("search")
		output := []models.Item{}
		for _, inventory := range items {
			for _, item := range inventory {
				if item.Tag == nil || item.Tag.Display.Name == "" {
					continue
				}

				if strings.Contains(strings.ToLower(item.Tag.Display.Name), searchString) || strings.Contains(strings.Join(item.Tag.Display.Lore, " "), searchString) {
					output = append(output, item)
				}
			}
		}

		fmt.Printf("Returning /api/inventory/%s/%s/%s in %s\n", uuid, inventoryId, searchString, time.Since(timeNow))

		return c.JSON(fiber.Map{
			"items": output,
		})

	}

	fmt.Printf("Returning /api/inventory/%s/%s in %s\n", uuid, inventoryId, time.Since(timeNow))

	itemSlice := items[inventoryId]
	return c.JSON(fiber.Map{
		"items": statsItems.ProcessItems(&itemSlice, inventoryId),
	})
}
