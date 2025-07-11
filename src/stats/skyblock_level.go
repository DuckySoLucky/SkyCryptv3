package stats

import (
	"skycrypt/src/models"
	stats "skycrypt/src/stats/leveling"
)

func GetSkyBlockLevel(userProfile *models.Member) models.Skill {

	return stats.GetLevelByXp(userProfile.Leveling.Experience, &stats.ExtraSkillData{Type: "skyblock_level"})
}
