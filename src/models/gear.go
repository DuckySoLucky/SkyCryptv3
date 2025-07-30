package models

type WeaponsResult struct {
	Weapons               []ProcessedItem `json:"weapons"`
	HighestPriorityWeapon *ProcessedItem  `json:"highest_priority_weapon"`
}

type Gear struct {
	Armor     ArmorResult        `json:"armor"`
	Equipment EquipmentResult    `json:"equipment"`
	Wardrobe  [][]*ProcessedItem `json:"wardrobe"`
	Weapons   WeaponsResult      `json:"weapons"`
}
