package models

type MinionsOutput struct {
	Minions      map[string]MinionCategory `json:"minions"`
	TotalMinions int                       `json:"totalMinions"`
	MaxedMinions int                       `json:"maxedMinions"`
	TotalTiers   int                       `json:"totalTiers"`
	MaxedTiers   int                       `json:"maxedTiers"`
	MinionSlots  *MinionSlotsOutput        `json:"minionsSlots"`
}

type MinionCategory struct {
	Minions      []Minion `json:"minions"`
	Texture      string   `json:"texture"`
	TotalMinions int      `json:"totalMinions"`
	MaxedMinions int      `json:"maxedMinions"`
	TotalTiers   int      `json:"totalTiers"`
	MaxedTiers   int      `json:"maxedTiers"`
}

type Minion struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
	MaxTier int    `json:"maxTier"`
	Tiers   []int  `json:"tiers"`
}

type MinionSlotsOutput struct {
	BonusSlots int `json:"bonusSlots"`
	Current    int `json:"current"`
	Next       int `json:"next"`
}
