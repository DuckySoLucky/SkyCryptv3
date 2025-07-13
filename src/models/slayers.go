package models

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
