package routes

import (
	"fmt"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	stats "skycrypt/src/stats"
	statsItems "skycrypt/src/stats/items"

	"time"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func GearHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	var items map[string][]models.Item

	// uuid := c.Params("uuid")
	profileId := c.Params("profileId")

	cache, err := redis.Get(fmt.Sprintf("items:%s", profileId))
	if err == nil && cache != "" {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal([]byte(cache), &items)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to parse items: %v", err),
			})
		}
	}

	if items == nil || items["armor"] == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("No items found for profile %s", profileId),
		})
	}

	var processedItems = make(map[string][]models.ProcessedItem)
	inventoryKeys := []string{"armor", "equipment", "wardrobe", "inventory", "enderchest", "backpack"}
	for _, inventoryId := range inventoryKeys {
		inventoryData := items[inventoryId]
		if len(inventoryData) == 0 {
			continue
		}

		processedItems[inventoryId] = statsItems.ProcessItems(&inventoryData, inventoryId)
	}

	allItems := make([]models.ProcessedItem, 0)
	for _, inventoryId := range inventoryKeys {
		if processedItems[inventoryId] != nil {
			allItems = append(allItems, processedItems[inventoryId]...)
		}
	}

	fmt.Printf("Returning /api/gear/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"gear": stats.GetGear(processedItems, allItems),
	})
}
