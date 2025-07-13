package stats

import (
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"strings"
)

func GetKuudraCompletions(userProfile *models.Member) int {
	if userProfile.CrimsonIsle.Kuudra == nil {
		return 0
	}

	kills := 0
	for kuudraId, kuudrakills := range *userProfile.CrimsonIsle.Kuudra {
		if kuudraId == "total" || strings.HasPrefix(kuudraId, "highest") {
			continue
		}

		kills += kuudrakills * constants.KUUDRA_COMPLETIONS_MULTIPLIER[kuudraId]
	}

	return kills
}
