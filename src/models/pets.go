package models

type ProcessedPet struct {
	Type      string             `json:"type,omitempty"`
	Name      string             `json:"display_name,omitempty"`
	Rarity    string             `json:"rarity,omitempty"`
	Active    bool               `json:"active,omitempty"`
	Price     float64            `json:"price,omitempty"`
	Level     PetLevel           `json:"level"`
	Texture   string             `json:"texture_path,omitempty"`
	Lore      []string           `json:"lore,omitempty"`
	Stats     map[string]float64 `json:"stats,omitempty"`
	CandyUsed int                `json:"candyUsed,omitempty"`
	Skin      string             `json:"skin,omitempty"`
}

type PetLevel struct {
	Experience            int     `json:"xp,omitempty"`
	Level                 int     `json:"level,omitempty"`
	CurrentExperience     int     `json:"currentXp,omitempty"`
	ExperienceForNext     int     `json:"xpForNext,omitempty"`
	Progress              float64 `json:"progress,omitempty"`
	ExperienceForMaxLevel int     `json:"xpMaxLevel,omitempty"`
}
