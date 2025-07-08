package routes

import (
	"fmt"
	"skycrypt/src/api"
	"skycrypt/src/models"
	"skycrypt/src/stats"
	"sync"
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

	var (
		profiles    *models.HypixelProfilesResponse
		player      *models.Player
		profilesErr error
		playerErr   error
		wg          sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		profiles, profilesErr = api.GetProfiles(mowojang.UUID)
	}()

	go func() {
		defer wg.Done()
		player, playerErr = api.GetPlayer(mowojang.UUID)
	}()

	wg.Wait()

	if profilesErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", profilesErr),
		})
	}

	if playerErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get player: %v", playerErr),
		})
	}

	profile, err := stats.GetProfile(profiles, profileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get profile: %v", err),
		})
	}

	var (
		profileMuseum *map[string]models.Museum
		members       []*models.MemberStats
		museumErr     error
		membersErr    error
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		profileMuseum, museumErr = api.GetMuseum(profile.ProfileID)
	}()

	go func() {
		defer wg.Done()
		members, membersErr = stats.FormatMembers(profile)
	}()

	wg.Wait()

	if museumErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get museum: %v", museumErr),
		})
	}

	if membersErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to format members: %v", membersErr),
		})
	}

	userProfileValue := profile.Members[mowojang.UUID]
	museum := (*profileMuseum)[mowojang.UUID]
	userProfile := &userProfileValue

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
		"museum":            museum.Value,
	})
}
