package stats

import (
	notenoughupdates "skycrypt/src/NotEnoughUpdates"
	"skycrypt/src/models"
	neu "skycrypt/src/models/NEU"
	"slices"
	"strconv"
)

type BestiaryOutput struct {
	Level             float64                           `json:"level"`
	MaxLevel          float64                           `json:"maxLevel"`
	FamiliesUnlocked  int                               `json:"familiesUnlocked"`
	FamiliesCompleted int                               `json:"familiesCompleted"`
	FamilyTiers       int                               `json:"familyTiers"`
	MaxFamilyTiers    int                               `json:"maxFamilyTiers"`
	TotalFamilies     int                               `json:"totalFamilies"`
	Categories        map[string]BestiaryCategoryOutput `json:"categories"`
}

type BestiaryCategoryOutput struct {
	Name         string              `json:"name"`
	Texture      string              `json:"texture"`
	Mobs         []BestiaryMobOutput `json:"mobs"`
	MobsUnlocked int                 `json:"mobsUnlocked"`
	MobsMaxed    int                 `json:"mobsMaxed"`
}

type BestiaryMobOutput struct {
	Name          string `json:"name"`
	Texture       string `json:"texture"`
	Kills         int    `json:"kills"`
	NextTierKills int    `json:"nextTierKills"`
	MaxKills      int    `json:"maxKills"`
	Tier          int    `json:"tier"`
	MaxTier       int    `json:"maxTier"`
}

func getCategoryMobs(userProfile *models.Member, mobs []neu.BestiaryMob) []BestiaryMobOutput {
	mobOutputs := make([]BestiaryMobOutput, 0)

	bestiaryKills := userProfile.Bestiary.Kills
	brackets := notenoughupdates.NEUConstants.Bestiary.Brackets
	for _, mobData := range mobs {
		kills := 0
		for _, mobName := range mobData.Mobs {
			if killCount, exists := bestiaryKills[mobName]; exists {
				kills += killCount
			}
		}

		mobBracket := brackets[strconv.Itoa(mobData.Bracket)]
		nextTierKills := 0
		tier := 0
		maxTier := slices.Index(mobBracket, mobData.Cap)
		if maxTier == -1 {
			maxTier = len(mobBracket) - 1
		}

		for i, bracketKills := range mobBracket {
			if kills < bracketKills && bracketKills <= mobData.Cap {
				nextTierKills = bracketKills
				tier = i
				break
			}

			if i == len(mobBracket)-1 && kills >= bracketKills {
				tier = maxTier
				nextTierKills = 0
			}
		}

		mobOutput := BestiaryMobOutput{
			Name:          mobData.Name,
			Texture:       mobData.Texture,
			Kills:         kills,
			NextTierKills: nextTierKills,
			MaxKills:      mobData.Cap,
			Tier:          tier,
			MaxTier:       maxTier,
		}

		mobOutputs = append(mobOutputs, mobOutput)
	}

	return mobOutputs
}

func GetBestiary(userProfile *models.Member) *BestiaryOutput {
	output := &BestiaryOutput{
		Categories:        make(map[string]BestiaryCategoryOutput),
		FamiliesCompleted: 0,
		FamiliesUnlocked:  0,
		FamilyTiers:       0,
		MaxFamilyTiers:    0,
		TotalFamilies:     0,
	}

	for categoryId, categoryData := range notenoughupdates.NEUConstants.Bestiary.Islands {
		categoryData := BestiaryCategoryOutput{
			Name:    categoryData.Name,
			Texture: categoryData.Texture,
			Mobs:    getCategoryMobs(userProfile, categoryData.Mobs),
		}

		categoryData.MobsUnlocked = 0
		categoryData.MobsMaxed = 0
		tiers, maxTiers := 0, 0
		for _, mob := range categoryData.Mobs {
			tiers += mob.Tier
			maxTiers += mob.MaxTier
			if mob.Kills > 0 {
				categoryData.MobsUnlocked++
			}

			if mob.Kills >= mob.MaxKills {
				categoryData.MobsMaxed++
			}

		}

		output.FamiliesUnlocked += categoryData.MobsUnlocked
		output.FamiliesCompleted += categoryData.MobsMaxed

		output.FamilyTiers += tiers
		output.MaxFamilyTiers += maxTiers
		output.TotalFamilies += len(categoryData.Mobs)

		output.Categories[categoryId] = categoryData
	}

	output.Level = float64(output.FamilyTiers) / 10.0
	output.MaxLevel = float64(output.MaxFamilyTiers) / 10.0

	return output
}
