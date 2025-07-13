package models

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
