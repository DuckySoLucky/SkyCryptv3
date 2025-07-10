package routes

import (
	"fmt"
	"skycrypt/src/api"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/stats"
	"time"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func AccessoriesHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	var items map[string][]models.Item

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

	if items == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("No items found for profile %s", profileId),
		})
	}

	fmt.Printf("Returning /api/accessories/%s in %s\n", profileId, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"accessories": stats.GetAccessories(userProfile, items),
	})
}
