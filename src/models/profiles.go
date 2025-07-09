package models

type HypixelProfilesResponse struct {
	Success  bool      `json:"success"`
	Cause    string    `json:"cause,omitempty"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	ProfileID string            `json:"profile_id"`
	CuteName  string            `json:"cute_name"`
	Selected  bool              `json:"selected"`
	Members   map[string]Member `json:"members"`
	GameMode  string            `json:"game_mode,omitempty"`
	Banking   *banking          `json:"banking,omitempty"`
}

type Member struct {
	PlayerData     *playerData     `json:"player_data"`
	CoopInvitation *coopInvitation `json:"coop_invitation"`
	Profile        *profileData    `json:"profile"`
	JacobsContest  *jacobsContest  `json:"jacobs_contest,omitempty"`
	PetsData       *petsData       `json:"pets_data,omitempty"`
	Leveling       *leveling       `json:"leveling,omitempty"`
	Currencies     *currencies     `json:"currencies,omitempty"`
	FairySouls     *fairySouls     `json:"fairy_soul,omitempty"`
	Inventory      *inventory      `json:"inventory,omitempty"`
	Rift           *rift           `json:"rift,omitempty"`
}

type coopInvitation struct {
	Confirmed bool `json:"confirmed,omitempty"`
}

type playerData struct {
	Experience *experience `json:"experience"`
}

type experience struct {
	SkillFishing    float64 `json:"SKILL_FISHING"`
	SkillAlchemy    float64 `json:"SKILL_ALCHEMY"`
	SkillMining     float64 `json:"SKILL_MINING"`
	SkillFarming    float64 `json:"SKILL_FARMING"`
	SkillEnchanting float64 `json:"SKILL_ENCHANTING"`
	SkillTaming     float64 `json:"SKILL_TAMING"`
	SkillForaging   float64 `json:"SKILL_FORAGING"`
	SkillSocial     float64 `json:"SKILL_SOCIAL"`
	SkillCarpentry  float64 `json:"SKILL_CARPENTRY"`
	SkillCombat     float64 `json:"SKILL_COMBAT"`
}

type profileData struct {
	DeletionNotice *deletionNotice `json:"deletion_notice"`
	FirstJoin      int64           `json:"first_join,omitempty"`
	BankAccount    float64         `json:"bank_account,omitempty"`
}

type deletionNotice struct {
	Timestamp int64 `json:"timestamp,omitempty"`
}

type jacobsContest struct {
	Perks *perks `json:"perks,omitempty"`
}

type perks struct {
	FarmingLevelCap int `json:"farming_level_cap,omitempty"`
}

type petsData struct {
	PetCare *petCare `json:"pet_care,omitempty"`
}

type petCare struct {
	PetTypesSacrificed []string `json:"pet_types_sacrificed,omitempty"`
}

type leveling struct {
	Experience int `json:"experience,omitempty"`
}

type currencies struct {
	CoinPurse float64 `json:"coin_purse,omitempty"`
}

type banking struct {
	Balance float64 `json:"balance,omitempty"`
}

type fairySouls struct {
	TotalCollected int `json:"total_collected,omitempty"`
}

type inventory struct {
	Inventory     encodedInventory            `json:"inv_contents"`
	Enderchest    encodedInventory            `json:"ender_chest_contents"`
	BackpackIcons map[string]encodedInventory `json:"backpack_icons"`
	Armor         encodedInventory            `json:"inv_armor"`
	Equipment     encodedInventory            `json:"equipment_contents"`
	PersonalVault encodedInventory            `json:"personal_vault_contents"`
	Backpack      map[string]encodedInventory `json:"backpack_contents"`
	Wardrobe      encodedInventory            `json:"wardrobe_contents"`
	BagContents   bagContents                 `json:"bag_contents"`
}

type encodedInventory struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

type bagContents struct {
	PotionBag   encodedInventory `json:"potion_bag,omitempty"`
	TalismanBag encodedInventory `json:"talisman_bag,omitempty"`
	FishingBag  encodedInventory `json:"fishing_bag,omitempty"`
	SacksBag    encodedInventory `json:"sacks_bag,omitempty"`
	Quiver      encodedInventory `json:"quiver,omitempty"`
}

type rift struct {
	Inventory riftInventory `json:"inventory,omitempty"`
}

type riftInventory struct {
	Inventory  encodedInventory `json:"inv_contents"`
	Armor      encodedInventory `json:"inv_armor"`
	Enderchest encodedInventory `json:"ender_chest_contents"`
	Equipment  encodedInventory `json:"equipment_contents"`
}
