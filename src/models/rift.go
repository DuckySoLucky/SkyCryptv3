package models

type RiftOutput struct {
	Visits     int                  `json:"visits"`
	Motes      RiftMotesOutput      `json:"motes"`
	Enigma     RiftEnigmaOutput     `json:"enigma"`
	Castle     RiftCastleOutput     `json:"castle"`
	Porhtals   RiftPortalsOutput    `json:"portals"`
	Timecharms RiftTimecharmsOutput `json:"timecharms"`
	Armor      ArmorResult          `json:"armor"`
	Equipment  EquipmentResult      `json:"equipment"`
}

type RiftMotesOutput struct {
	Purse    int `json:"purse"`
	Lifetime int `json:"lifetime"`
	Orbs     int `json:"orbs"`
}

type RiftEnigmaOutput struct {
	Souls      int `json:"souls"`
	TotalSouls int `json:"totalSouls"`
}

type RiftCastleOutput struct {
	GrubberStacks int `json:"grubberStacks"`
	MaxBurgers    int `json:"maxBurgers"`
}

type RiftPortalsOutput struct {
	PorhtalsFound int           `json:"portalsFound"`
	Porhtals      []RiftPorhtal `json:"portals"`
}

type RiftPorhtal struct {
	Name     string `json:"name"`
	Texture  string `json:"texture"`
	Unlocked bool   `json:"unlocked"`
}

type RiftTimecharmsOutput struct {
	TimecharmsFound int              `json:"timecharmsFound"`
	Timecharms      []RiftTimecharms `json:"timecharms"`
}

type RiftTimecharms struct {
	Name       string `json:"name"`
	Texture    string `json:"texture"`
	Unlocked   bool   `json:"unlocked"`
	UnlockedAt int64  `json:"unlockedAt"`
}
