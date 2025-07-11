package stats

import (
	"fmt"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	statsItems "skycrypt/src/stats/items"
	stats "skycrypt/src/stats/leveling"

	"skycrypt/src/utility"
	"slices"
	"strings"
)

type miningOutput struct {
	Level                  models.Skill            `json:"level"`
	PeakOfTheMountain      peakOfTheMountain       `json:"peak_of_the_mountain"`
	SelectedPickaxeAbility string                  `json:"selected_pickaxe_ability"`
	Tokens                 hotmTokens              `json:"tokens"`
	Commissions            commissions             `json:"commissions"`
	CrystalHollows         crystalHollows          `json:"crystalHollows"`
	Powder                 powderOutput            `json:"powder"`
	GlaciteTunnels         glaciteTunnels          `json:"glaciteTunnels"`
	Forge                  []forgeOutput           `json:"forge"`
	Tools                  models.SkillToolsResult `json:"tools"`
}

type peakOfTheMountain struct {
	Level    int `json:"level"`
	MaxLevel int `json:"max_level"`
}

type hotmTokens struct {
	Total     int `json:"total"`
	Spent     int `json:"spent"`
	Available int `json:"available"`
}

type commissions struct {
	Milestone   int `json:"milestone"`
	Completions int `json:"completions"`
}

type crystalHollows struct {
	CrystalHollowsLastAccess int64              `json:"crystalHollowsLastAccess"`
	NucleusRuns              int                `json:"nucleusRuns"`
	Progress                 crystalNucleusRuns `json:"progress"`
}

type crystalNucleusRuns struct {
	Crystals map[string]string `json:"crystals"`
	Parts    map[string]string `json:"parts"`
}

type glaciteTunnels struct {
	MineshaftsEntered int     `json:"mineshaftsEntered"`
	FossilDust        float64 `json:"fossilDust"`
	Corpses           Corpses `json:"corpses"`
	Fossils           Fossils `json:"fossils"`
}

type Corpses struct {
	Found   int      `json:"found"`
	Max     int      `json:"max"`
	Corpses []Corpse `json:"corpses"`
}

type Fossils struct {
	Found   int      `json:"found"`
	Max     int      `json:"max"`
	Fossils []Fossil `json:"fossils"`
}

type Corpse struct {
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	Texture string `json:"texture_path"`
}

type Fossil struct {
	Name    string `json:"name"`
	Found   bool   `json:"amount"`
	Texture string `json:"texture_path"`
}

func getPeakOfTheMountain(userProfile *models.Member) peakOfTheMountain {
	return peakOfTheMountain{
		Level:    userProfile.Mining.Nodes["special_0"],
		MaxLevel: constants.MAX_PEAK_OF_THE_MOUNTAIN_LEVEL,
	}
}

func getSelectedPickaxeAbility(userProfile *models.Member) string {
	if userProfile.Mining.SelectedPickaxeAbility == "" {
		return "None"
	}

	return userProfile.Mining.SelectedPickaxeAbility
}

func calcHotmTokens(hotmTier int, potmTier int) int {
	tokens := 0

	for tier := 1; tier <= hotmTier; tier++ {
		if reward, ok := constants.HOTM_REWARDS[tier]; ok {
			tokens += reward.TokenOfTheMountain
		}
	}

	for tier := 1; tier <= potmTier; tier++ {
		if reward, ok := constants.POTM_REWARDS[tier]; ok {
			tokens += reward.TokenOfTheMountain
		}
	}

	return tokens
}

func getHotmTokens(hotmLevel models.Skill, userProfile *models.Member) hotmTokens {
	potmTier := userProfile.Mining.Nodes["special_0"]
	hotmTokensAmount := calcHotmTokens(hotmLevel.Level, potmTier)
	return hotmTokens{
		Total:     hotmTokensAmount,
		Spent:     userProfile.Mining.TokensSpent,
		Available: hotmTokensAmount - userProfile.Mining.TokensSpent,
	}
}

func getCommissions(userProfile *models.Member, player *models.Player) commissions {
	var milestone = 0
	for _, tutorial := range userProfile.Objectives.Tutorial {
		if strings.HasPrefix(tutorial, "commission_milestone_reward_mining_xp_tier_") {
			tier := strings.Split(tutorial, "_")
			if len(tier) > 0 {
				lastElement := tier[len(tier)-1]
				parsedTier, err := utility.ParseInt(lastElement)
				if err == nil {
					if parsedTier > milestone {
						milestone = parsedTier

					}
				}
			}

		}
	}

	return commissions{
		Milestone:   milestone,
		Completions: player.Achievements.HotMCommissions,
	}
}

func getCrystalHollows(userProfile *models.Member) crystalHollows {
	totalRuns := 0
	for _, crystalData := range userProfile.Mining.Crystals {
		if crystalData.TotalPlaced > totalRuns {
			totalRuns = crystalData.TotalPlaced
		}
	}

	crystalHollows := crystalHollows{
		CrystalHollowsLastAccess: userProfile.Mining.GreaterMinesLastAccess,
		NucleusRuns:              totalRuns,
		Progress: crystalNucleusRuns{
			Crystals: make(map[string]string),
			Parts:    make(map[string]string),
		},
	}

	for _, crystal := range constants.GEMSTONE_CRYSTALS {
		crystalKey := crystal + "_crystal"

		if crystalData, exists := userProfile.Mining.Crystals[crystalKey]; exists {
			crystalHollows.Progress.Crystals[crystal] = crystalData.State
		} else {
			crystalHollows.Progress.Crystals[crystal] = "NOT_FOUND"
		}
	}

	for _, part := range constants.PRECURSOR_PARTS {
		partKey := strings.ToUpper(part)
		partsDelivered := userProfile.Mining.Biomes.Precursor.PartsDelivered

		if slices.Contains(partsDelivered, partKey) {
			crystalHollows.Progress.Parts[part] = "DELIVERED"
		} else {
			crystalHollows.Progress.Parts[part] = "NOT_DELIVERED"
		}
	}

	return crystalHollows
}

type powderAmount struct {
	Spent     int `json:"spent"`
	Total     int `json:"total"`
	Available int `json:"available"`
}

type powderOutput struct {
	Mithril  powderAmount `json:"mithril"`
	Gemstone powderAmount `json:"gemstone"`
	Glacite  powderAmount `json:"glacite"`
}

func getPowderAmount(userProfile *models.Member, powderType string) powderAmount {
	spent := 0
	total := 0
	available := 0

	switch powderType {
	case "mithril":
		available = userProfile.Mining.PowderMithril
		spent = userProfile.Mining.PowderSpentMithril
		total = userProfile.Mining.PowderMithrilTotal
	case "gemstone":
		available = userProfile.Mining.PowderGemstone
		spent = userProfile.Mining.PowderSpentGemstone
		total = userProfile.Mining.PowderGemstoneTotal
	case "glacite":
		available = userProfile.Mining.PowderGlacite
		spent = userProfile.Mining.PowderSpentGlacite
		total = userProfile.Mining.PowderGlaciteTotal
	}

	return powderAmount{
		Spent:     spent,
		Total:     total,
		Available: available,
	}
}

func getPowder(userProfile *models.Member) powderOutput {
	return powderOutput{
		Mithril:  getPowderAmount(userProfile, "mithril"),
		Gemstone: getPowderAmount(userProfile, "gemstone"),
		Glacite:  getPowderAmount(userProfile, "glacite"),
	}

}

// ...existing code...
type forgeOutput struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Slot         int     `json:"slot"`
	StartingTime int64   `json:"startingTime"`
	EndingTime   int64   `json:"endingTime"`
	Duration     float64 `json:"duration"`
}

func getForge(userProfile *models.Member) []forgeOutput {
	output := []forgeOutput{}

	quickForgeLevel := userProfile.Mining.Nodes["forge_time"]
	var quickForge float64
	if quickForgeLevel > 0 {
		if quickForgeLevel <= 19 {
			quickForge = float64(100+quickForgeLevel*5) / 10.0
		} else {
			quickForge = 300.0 / 10.0
		}
	} else {
		quickForge = 0
	}

	for _, item := range userProfile.Forge.ForgeProcesses.Forge {
		forgeConst := constants.FORGE[item.Id]
		duration := float64(forgeConst.Duration) - float64(forgeConst.Duration)*quickForge
		endingTime := item.StartTime + int64(float64(forgeConst.Duration)-float64(forgeConst.Duration)*(quickForge/100))
		output = append(output, forgeOutput{
			ID:           item.Id,
			Name:         forgeConst.Name,
			Slot:         item.Slot,
			StartingTime: item.StartTime,
			EndingTime:   endingTime,
			Duration:     duration,
		})
	}

	return output
}

func getGlaciteTunnels(userProfile *models.Member) glaciteTunnels {
	output := glaciteTunnels{
		MineshaftsEntered: userProfile.GlaciteTunnels.MineshaftsEntered,
		FossilDust:        userProfile.GlaciteTunnels.FossilDust,
		Corpses:           Corpses{},
		Fossils:           Fossils{},
	}

	found := 0
	for corpseId, corpseTexture := range constants.CORPSES {
		found += userProfile.GlaciteTunnels.CorpsesLooted[corpseId]
		corpse := Corpse{
			Amount:  userProfile.GlaciteTunnels.CorpsesLooted[corpseId],
			Name:    utility.TitleCase(corpseId),
			Texture: corpseTexture,
		}

		output.Corpses.Corpses = append(output.Corpses.Corpses, corpse)
	}

	output.Corpses.Found = found
	output.Corpses.Max = len(constants.CORPSES)

	found = 0
	for _, fossil := range constants.FOSSILS {
		isFound := slices.Contains(userProfile.GlaciteTunnels.FossilsDonated, fossil)
		if isFound {
			found++
		}

		texture := fmt.Sprintf("/api/item/%s_FOSSIL", fossil)
		if fossil == "HELIX" {
			texture = fmt.Sprintf("/api/item/%s", fossil)
		}

		fossil := Fossil{
			Name:    utility.TitleCase(fossil),
			Texture: texture,
			Found:   isFound,
		}

		output.Fossils.Fossils = append(output.Fossils.Fossils, fossil)
	}

	output.Fossils.Found = found
	output.Fossils.Max = len(constants.FOSSILS)

	return output
}

func GetMining(userProfile *models.Member, player *models.Player, items []models.ProcessedItem) miningOutput {
	HOTMLevel := stats.GetLevelByXp(int(userProfile.Mining.Experience), &stats.ExtraSkillData{Type: "hotm"})

	return miningOutput{
		Level:                  HOTMLevel,
		PeakOfTheMountain:      getPeakOfTheMountain(userProfile),
		SelectedPickaxeAbility: getSelectedPickaxeAbility(userProfile),
		Tokens:                 getHotmTokens(HOTMLevel, userProfile),
		Commissions:            getCommissions(userProfile, player),
		CrystalHollows:         getCrystalHollows(userProfile),
		Powder:                 getPowder(userProfile),
		GlaciteTunnels:         getGlaciteTunnels(userProfile),
		Forge:                  getForge(userProfile),
		Tools:                  statsItems.GetSkillTools("mining", items),
	}
}
