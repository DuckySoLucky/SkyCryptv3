package neu

type customPetLeveling struct {
	Type         int             `json:"type,omitempty"`
	PetLevels    []int           `json:"pet_levels,omitempty"`
	MaxLevel     int             `json:"max_level,omitempty"`
	RarityOffset map[string]*int `json:"rarity_offset,omitempty"`
	XPMultiplier int             `json:"xp_multiplier,omitempty"`
}

type Pets struct {
	PetRarityOffset        map[string]*int              `json:"pet_rarity_offset"`
	PetLevels              []int                        `json:"pet_levels"`
	CustomPetLeveling      map[string]customPetLeveling `json:"custom_pet_leveling"`
	PetTypes               map[string]string            `json:"pet_types"`
	IDToDisplayName        map[string]string            `json:"id_to_display_name"`
	PetItemDisplayNameToID map[string]string            `json:"pet_item_display_name_to_id"`
}
