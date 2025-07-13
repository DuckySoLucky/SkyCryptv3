package models

type GetMissingAccessoresOutput struct {
	Stats               map[string]float64    `json:"stats"`
	Enrichments         map[string]int        `json:"enrichments"`
	Unique              int                   `json:"unique"`
	Total               int                   `json:"total"`
	Recombobulated      int                   `json:"recombobulated"`
	TotalRecombobulated int                   `json:"total_recombobulated"`
	SelectedPower       string                `json:"selected_power"`
	MagicalPower        GetMagicalPowerOutput `json:"magicalPower"`
	Accessories         []InsertAccessory     `json:"accessories"`
	Missing             []ProcessedItem       `json:"missing"`
	Upgrades            []ProcessedItem       `json:"upgrades"`
}

type MissingOutput struct {
	Upgrades     []ProcessedItem `json:"upgrades"`
	Other        []ProcessedItem `json:"other"`
	AccessoryIds []AccessoryIds  `json:"accessoryIds"`
}

type GetMagicalPowerOutput struct {
	Total       int `json:"total"`
	Accessories int `json:"accessories"`
	Abiphone    int `json:"abiphone"`
	RiftPrism   int `json:"riftPrism"`
	Hegemony    struct {
		Rarity string `json:"rarity"`
		Amount int    `json:"amount"`
	} `json:"hegemony"`
	Rarities map[string]struct {
		Rarity int `json:"rarity"`
		Amount int `json:"magicalPower"`
	} `json:"rarities"`
}

type InsertAccessory struct {
	ProcessedItem
	Id         string `json:"id"`
	Rarity     string `json:"rarity"`
	IsInactive bool   `json:"isInactive,omitempty"`
}

type AccessoryIds struct {
	Id     string `json:"id"`
	Rarity string `json:"rarity"`
}

type AccessoriesOutput struct {
	Accessories  []InsertAccessory `json:"accessories"`
	AccessoryIds []AccessoryIds    `json:"accessoryIds"`
}
