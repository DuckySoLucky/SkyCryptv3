package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/stats"
	"time"

	"github.com/gofiber/fiber/v2"
)

func StatsHandler(c *fiber.Ctx) error {
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

	profiles, err := api.GetProfiles(mowojang.UUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	player, err := api.GetPlayer(mowojang.UUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get player: %v", err),
		})
	}

	profile, err := stats.GetProfile(profiles, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	profileMuseum, err := api.GetMuseum(profile.ProfileID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get museum: %v", err),
		})
	}

	members, err := stats.FormatMembers(profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to format members: %v", err),
		})
	}

	userProfileValue := profile.Members[mowojang.UUID]
	museum := profileMuseum[mowojang.UUID]
	userProfile := &userProfileValue

	stats.GetItems(userProfile, profile.ProfileID)

	fmt.Printf("Returning /api/stats/%s in %s\n", uuid, time.Since(timeNow))

	return c.JSON(fiber.Map{
		"username":          mowojang.Name,
		"uuid":              mowojang.UUID,
		"profile_id":        profile.ProfileID,
		"profile_cute_name": profile.CuteName,
		"selected":          profile.Selected,
		"profiles":          stats.FormatProfiles(profiles),
		"members":           members,
		"social":            player.SocialMedia.Links,
		"rank":              stats.GetRank(player),
		"skills":            stats.GetSkills(userProfile, profile, player),
		"skyblock_level":    stats.GetSkyBlockLevel(userProfile),
		"joined":            userProfile.Profile.FirstJoin,
		"purse":             userProfile.Currencies.CoinPurse,
		"bank":              profile.Banking.Balance,
		"personalBank":      userProfile.Profile.BankAccount,
		"fairySouls":        stats.GetFairySouls(userProfile, profile.GameMode),
		"apiSettings":       stats.GetAPISettings(userProfile, profile, museum),
	})
}
