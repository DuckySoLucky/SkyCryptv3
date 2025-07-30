package stats

import "skycrypt/src/models"

func GetAPISettings(userProfile *models.Member, profile *models.Profile, museum *models.Museum) map[string]bool {
	return map[string]bool{
		"skills":         userProfile.PlayerData.Experience != nil,
		"inventory":      userProfile.Inventory != nil,
		"personal_vault": userProfile.Inventory.PersonalVault.Data != "",
		"collections":    userProfile.Collections != nil,
		"banking":        profile.Banking.Balance != nil,
		"museum":         museum != nil,
	}
}
