package models

type HypixelGardenResponse struct {
	Success bool      `json:"success"`
	Cause   string    `json:"cause,omitempty"`
	Garden  GardenRaw `json:"garden"`
}

type GardenRaw struct {
	UnlockedPlotsIds []string `json:"unlocked_plots_ids"`
	CommissionData   struct {
		Visits           map[string]int `json:"visits"`
		Completed        map[string]int `json:"completed"`
		TotalCompleted   int            `json:"total_completed"`
		UniqueNpcsServed int            `json:"unique_npcs_served"`
	} `json:"commission_data"`
	ResourcesCollected map[string]float64 `json:"resources_collected"`
	ComposterData      struct {
		Upgrades map[string]int `json:"upgrades"`
	} `json:"composter_data"`
	Experience        float64        `json:"garden_experience"`
	SelectedBarnSkin  string         `json:"selected_barn_skin"`
	CropUpgradeLevels    map[string]int `json:"crop_upgrade_levels"`
	UnlockedBarnSkins []string       `json:"unlocked_barn_skins"`
}
