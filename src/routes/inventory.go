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
	if inventoryId == "museum" {
		profileMuseum, err := api.GetMuseum(profileId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to get museum: %v", err),
			})
		}

		museum := profileMuseum[uuid]
		if museum == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Sprintf("No museum data found for profile %s", profileId),
			})
		}

		fmt.Printf("Returning /api/inventory/%s/%s in %s\n", uuid, inventoryId, time.Since(timeNow))

		return c.JSON(fiber.Map{
			"items": statsItems.StripItems(statsItems.GetMuseum(museum)),
		})

	}

	profile, err := api.GetProfile(uuid, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	userProfileValue := profile.Members[uuid]
	userProfile := &userProfileValue

	if inventoryId == "search" {
		items, err := stats.GetItems(userProfile, profile.ProfileID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to get items: %v", err),
			})
		}

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

	itemSlice := stats.GetInventory(userProfile, inventoryId)
	output := statsItems.ProcessItems(&itemSlice, inventoryId)

	fmt.Printf("Returning /api/inventory/%s/%s in %s\n", uuid, inventoryId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"items": statsItems.StripItems(output),
	})
}
