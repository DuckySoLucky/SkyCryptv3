package stats

import (
	"fmt"
	"skycrypt/src/constants"
	"skycrypt/src/models"
)

type SlayersOutput struct {
	Data                  map[string]SlayerData `json:"data"`
	TotalSlayerExperience int                   `json:"totalSlayerExp"`
	Stats                 map[string]float64    `json:"stats"`
}

type SlayerData struct {
	Name    string         `json:"name"`
	Texture string         `json:"texture"`
	Kills   map[string]int `json:"kills"`
	Level   SlayerLevel    `json:"level"`
}

type SlayerLevel struct {
	Experience        int  `json:"xp"`
	ExperienceForNext int  `json:"xpForNext"`
	Level             int  `json:"level"`
	MaxLevel          int  `json:"maxLevel"`
	Maxed             bool `json:"maxed"`
}

func getSlayerKills(slayerData models.SlayerBoss) map[string]int {
	tiers := []int{
		slayerData.BossKillsTier0,
		slayerData.BossKillsTier1,
		slayerData.BossKillsTier2,
		slayerData.BossKillsTier3,
		slayerData.BossKillsTier4,
	}

	total := 0
	kills := make(map[string]int)
	for i, count := range tiers {
		tier := i + 1
		kills[fmt.Sprintf("%d", tier)] = count
		total += count
	}

	kills["total"] = total

	return kills
}

func getSlayerLevel(experience int, slayerId string) SlayerLevel {
	if constants.SLAYER_INFO[slayerId].Levelling == nil {
		return SlayerLevel{
			Experience:        0,
			ExperienceForNext: 0,
			Level:             0,
			MaxLevel:          0,
			Maxed:             false,
		}
	}

	level := 0
	maxLevel := len(constants.SLAYER_INFO[slayerId].Levelling)
	for i := 1; i <= maxLevel; i++ {
		if experience < constants.SLAYER_INFO[slayerId].Levelling[i] {
			break
		}

		level = i
	}

	experienceForNext := 0
	if level < maxLevel {
		experienceForNext = constants.SLAYER_INFO[slayerId].Levelling[level+1]
	} else {
		experienceForNext = 0
	}

	return SlayerLevel{
		Experience:        experience,
		ExperienceForNext: experienceForNext,
		Level:             level,
		MaxLevel:          maxLevel,
		Maxed:             level == maxLevel,
	}
}

func GetSlayers(userProfile *models.Member) SlayersOutput {
	output := SlayersOutput{
		Data: make(map[string]SlayerData),
	}

	totalExperience := 0
	for slayerId, slayerData := range userProfile.Slayes.SlayerBosses {
		output.Data[slayerId] = SlayerData{
			Name:    constants.SLAYER_INFO[slayerId].Name,
			Texture: constants.SLAYER_INFO[slayerId].Head,
			Kills:   getSlayerKills(slayerData),
			Level:   getSlayerLevel(int(slayerData.Experience), slayerId),
		}

		totalExperience += int(slayerData.Experience)

		statsBonus := constants.STATS_BONUS[fmt.Sprintf("slayer_%s", slayerId)]
		if statsBonus == nil {
			continue
		}

		stats := constants.GetBonusStats(output.Data[slayerId].Level.Level, statsBonus)
		if output.Stats == nil {
			output.Stats = make(map[string]float64)
		}

		for stat, value := range stats {
			if _, exists := output.Stats[stat]; !exists {
				output.Stats[stat] = 0
			}
			
			output.Stats[stat] += value
		}
	}

	output.TotalSlayerExperience = totalExperience

	return output
}
