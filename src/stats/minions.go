package stats

import (
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"slices"
	"strings"
)

type MinionsOutput struct {
	Minions      map[string]MinionCategory `json:"minions"`
	TotalMinions int                       `json:"totalMinions"`
	MaxedMinions int                       `json:"maxedMinions"`
	TotalTiers   int                       `json:"totalTiers"`
	MaxedTiers   int                       `json:"maxedTiers"`
	MinionSlots  *MinionSlotsOutput        `json:"minionSlots"`
}

type MinionCategory struct {
	Minions      map[string]Minion `json:"minions"`
	Texture      string            `json:"texture"`
	TotalMinions int               `json:"totalMinions"`
	MaxedMinions int               `json:"maxedMinions"`
	TotalTiers   int               `json:"totalTiers"`
	MaxedTiers   int               `json:"maxedTiers"`
}

type Minion struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
	MaxTier int    `json:"maxTier"`
	Tiers   []int  `json:"tiers"`
}

type MinionSlotsOutput struct {
	BonusSlots int `json:"bonusSlots"`
	Current    int `json:"current"`
	Next       int `json:"next"`
}

func getMinionSlots(profile *models.Profile, tiers int) *MinionSlotsOutput {
	keys := make([]int, 0, len(constants.MINION_SLOTS))
	for k := range constants.MINION_SLOTS {
		keys = append(keys, k)
	}
	utility.SortInts(keys)

	highestKey := keys[0]
	nextTier := 0
	for _, key := range keys {
		if tiers >= key {
			highestKey = key
			nextTier = keys[slices.Index(keys, key)+1]
		} else {
			break
		}
	}

	bonusSlots := 0
	for _, upgrade := range profile.CommunityUpgrades.UpgradeStates {
		if upgrade.Upgrade == "minion_slots" {
			bonusSlots++
		}
	}

	return &MinionSlotsOutput{
		BonusSlots: bonusSlots,
		Current:    constants.MINION_SLOTS[highestKey],
		Next:       nextTier - tiers,
	}
}

func getCraftedMinions(profile *models.Profile) map[string][]int {
	craftedMinions := make(map[string][]int)
	for _, member := range profile.Members {
		for _, minion := range member.PlayerData.Minions {
			parts := strings.Split(minion, "_")
			tierStr := parts[len(parts)-1]
			minionType := strings.Join(parts[:len(parts)-1], "_")

			tier, err := utility.ParseInt(tierStr)
			if err != nil {
				continue
			}

			craftedMinions[minionType] = append(craftedMinions[minionType], tier)
		}
	}

	for minionType := range craftedMinions {
		utility.SortInts(craftedMinions[minionType])
	}

	return craftedMinions
}

func GetMinions(profile *models.Profile) MinionsOutput {
	craftedMinions := getCraftedMinions(profile)

	output := MinionsOutput{
		Minions: make(map[string]MinionCategory),
	}

	for categoryId, categoryData := range constants.MINIONS {
		category := MinionCategory{
			Minions:      make(map[string]Minion),
			Texture:      constants.MINION_CATEGORY_ICONS[categoryId],
			TotalMinions: 0,
			MaxedMinions: 0,
			TotalTiers:   0,
			MaxedTiers:   0,
		}

		totalTiers := 0
		for minionId, minionData := range categoryData {
			name := minionData.Name
			if name == "" {
				name = utility.TitleCase(minionId)
			}

			maxTier := minionData.MaxTier
			if maxTier == 0 {
				maxTier = 11 // Default max tier if not specified
			}

			category.Minions[minionId] = Minion{
				Name:    name,
				Texture: minionData.Texture,
				MaxTier: maxTier,
				Tiers:   craftedMinions[minionId],
			}

			totalTiers += maxTier
			category.MaxedTiers += len(craftedMinions[minionId])
			if craftedMinions[minionId][len(craftedMinions[minionId])-1] == maxTier {
				category.MaxedMinions++
			}
		}

		category.TotalMinions = len(category.Minions)
		category.TotalTiers = totalTiers

		output.Minions[categoryId] = category

		output.TotalMinions += category.TotalMinions
		output.MaxedMinions += category.MaxedMinions
		output.TotalTiers += category.TotalTiers
		output.MaxedTiers += category.MaxedTiers
	}

	output.MinionSlots = getMinionSlots(profile, output.MaxedTiers)

	return output
}
