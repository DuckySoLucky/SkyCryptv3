package models

import neu "skycrypt/src/models/NEU"

type NEUItem struct {
	MinecraftId string   `json:"itemid,omitempty"`
	Name        string   `json:"displayname,omitempty"`
	Damage      int      `json:"damage,omitempty"`
	Lore        []string `json:"lore,omitempty"`
	NEUId       string   `json:"internalname,omitempty"`
	NBT         string   `json:"nbttag"`
	Wiki        []string `json:"info,omitempty"`
}

type NEUConstant struct {
	PetNums  neu.PetNums           `json:"petnums,omitempty"`
	Pets     neu.Pets              `json:"pets,omitempty"`
	Bestiary neu.BestiaryConstants `json:"bestiary,omitempty"`
}
