package models

import (
	"encoding/json"
	"fmt"
)

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
	Banking   banking           `json:"banking,omitempty"`
}

type Member struct {
	PlayerData          *playerData          `json:"player_data"`
	CoopInvitation      *coopInvitation      `json:"coop_invitation"`
	Profile             *profileData         `json:"profile"`
	JacobsContest       *jacobsContest       `json:"jacobs_contest,omitempty"`
	Pets                *pets                `json:"pets_data,omitempty"`
	Leveling            *leveling            `json:"leveling,omitempty"`
	Currencies          *currencies          `json:"currencies,omitempty"`
	FairySouls          *fairySouls          `json:"fairy_soul,omitempty"`
	Inventory           *inventory           `json:"inventory,omitempty"`
	Rift                *rift                `json:"rift,omitempty"`
	AccessoryBagStorage *accessoryBagStorage `json:"accessory_bag_storage,omitempty"`
	CrimsonIsle         *crimsonIsleData     `json:"nether_island_player_data,omitempty"`
	Mining              *mining              `json:"mining_core,omitempty"`
	Objectives          *objectives          `json:"objectives,omitempty"`
	GlaciteTunnels      *glaciteData         `json:"glacite_player_data,omitempty"`
	Forge               *forge               `json:"forge,omitempty"`
	Quests              *quests              `json:"quests,omitempty"`
	Garden              *gardenProfileData   `json:"garden_player_data,omitempty"`
	PlayerStats         *playerStats         `json:"player_stats,omitempty"`
	TrophyFish          *memberTrophyFish    `json:"trophy_fish,omitempty"`
	Experimentation     *experimentationData `json:"experimentation,omitempty"`
	Dungeons            *Dungeons            `json:"dungeons,omitempty"`
	Slayes              *slayer              `json:"slayer,omitempty"`
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
	Perks          *perks                  `json:"perks,omitempty"`
	UniqueBrackets map[string][]string     `json:"unique_brackets,omitempty"`
	MedalsInv      map[string]int          `json:"medals_inv,omitempty"`
	Contests       map[string]JacobContest `json:"contests,omitempty"`
}

type perks struct {
	FarmingLevelCap int `json:"farming_level_cap,omitempty"`
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
	Balance *float64 `json:"balance,omitempty"`
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
	Access    *riftAccess   `json:"access,omitempty"`
	DeadCats  *deadCats     `json:"dead_cats,omitempty"`
}

type riftInventory struct {
	Inventory  encodedInventory `json:"inv_contents"`
	Armor      encodedInventory `json:"inv_armor"`
	Enderchest encodedInventory `json:"ender_chest_contents"`
	Equipment  encodedInventory `json:"equipment_contents"`
}

type riftAccess struct {
	ConsumedPrism bool `json:"consumed_prism,omitempty"`
}

type accessoryBagStorage struct {
	SelectedPower string `json:"selected_power,omitempty"`
}

type crimsonIsleData struct {
	Abiphone *abiphone `json:"abiphone,omitempty"`
}

type abiphone struct {
	ActiveContacts []string `json:"active_contacts,omitempty"`
}

type deadCats struct {
	FoundCats []string `json:"found_cats,omitempty"`
	Montezuma Pet      `json:"montezuma,omitempty"`
}

type Pet struct {
	Type       string  `json:"type,omitempty"`
	Experience float64 `json:"exp,omitempty"`
	Active     bool    `json:"active,omitempty"`
	Rarity     string  `json:"tier,omitempty"`
	HeldItem   string  `json:"heldItem,omitempty"`
	CandyUsed  int     `json:"candyUsed,omitempty"`
	Skin       string  `json:"skin,omitempty"`
}

type pets struct {
	PetCare *petCare `json:"pet_care,omitempty"`
	Pets    []Pet    `json:"pets,omitempty"`
}

type mining struct {
	Nodes                  map[string]int     `json:"nodes,omitempty"`
	Experience             float64            `json:"experience,omitempty"`
	GreaterMinesLastAccess int64              `json:"greater_mines_last_access,omitempty"`
	LastReset              int64              `json:"last_reset,omitempty"`
	TokensSpent            int                `json:"tokens_spent,omitempty"`
	SelectedPickaxeAbility string             `json:"selected_pickaxe_ability,omitempty"`
	PowderMithril          int                `json:"powder_mithril,omitempty"`
	PowderMithrilTotal     int                `json:"powder_mithril_total,omitempty"`
	PowderSpentMithril     int                `json:"powder_spent_mithril,omitempty"`
	PowderGemstone         int                `json:"powder_gemstone,omitempty"`
	PowderGemstoneTotal    int                `json:"powder_gemstone_total,omitempty"`
	PowderSpentGemstone    int                `json:"powder_spent_gemstone,omitempty"`
	PowderGlacite          int                `json:"powder_glacite,omitempty"`
	PowderGlaciteTotal     int                `json:"powder_glacite_total,omitempty"`
	PowderSpentGlacite     int                `json:"powder_spent_glacite,omitempty"`
	Crystals               map[string]crystal `json:"crystals,omitempty"`
	Biomes                 biomes             `json:"biomes,omitempty"`
}

type crystal struct {
	State       string `json:"state,omitempty"`
	TotalFound  int    `json:"total_found,omitempty"`
	TotalPlaced int    `json:"total_placed,omitempty"`
}

type biomes struct {
	Precursor precursor `json:"precursor,omitempty"`
}

type precursor struct {
	PartsDelivered []string `json:"parts_delivered,omitempty"`
}

type objectives struct {
	Tutorial []string `json:"tutorial,omitempty"`
}

type glaciteData struct {
	FossilsDonated    []string       `json:"fossils_donated,omitempty"`
	FossilDust        float64        `json:"fossil_dust,omitempty"`
	CorpsesLooted     map[string]int `json:"corpses_looted,omitempty"`
	MineshaftsEntered int            `json:"mineshafts_entered,omitempty"`
}

type forge struct {
	ForgeProcesses forgeProcesses `json:"forge_processes"`
}

type forgeProcesses struct {
	Forge map[string]forgeProcess `json:"forge_1"`
}

type forgeProcess struct {
	Id        string `json:"id"`
	StartTime int64  `json:"startTime"`
	Slot      int    `json:"slot"`
}

type quests struct {
	TrapperQuest *trapperQuest `json:"trapper_quest,omitempty"`
}

type trapperQuest struct {
	PeltCount int `json:"pelt_count,omitempty"`
}

type gardenProfileData struct {
	Copper        int `json:"copper,omitempty"`
	LarvaConsumed int `json:"larva_consumed,omitempty"`
}

type JacobContest struct {
	Collected           int    `json:"collected"`
	ClaimedPosition     *int   `json:"claimed_position,omitempty"`
	ClaimedParticipants *int   `json:"claimed_participants,omitempty"`
	ClaimedMedal        string `json:"claimed_medal"`
}

type playerStats struct {
	Kills       map[string]float64 `json:"kills,omitempty"`
	Deaths      map[string]float64 `json:"deaths,omitempty"`
	ItemsFished struct {
		Total         float64 `json:"total,omitempty"`
		Normal        float64 `json:"normal,omitempty"`
		Treasure      float64 `json:"treasure,omitempty"`
		LargeTreasure float64 `json:"large_treasure,omitempty"`
		TrophyFish    float64 `json:"trophy_fish,omitempty"`
	} `json:"items_fished"`
	ShredderRod struct {
		Fished float64 `json:"fished,omitempty"`
		Bait   float64 `json:"bait,omitempty"`
	} `json:"shredder_rod"`
	Pets struct {
		Milestone struct {
			SeaCreaturesKilled float64 `json:"sea_creatures_killed,omitempty"`
		} `json:"milestone,omitempty"`
	} `json:"pets,omitempty"`
}

type memberTrophyFish struct {
	Rewards     []int          `json:"rewards"`
	TotalCaught int            `json:"total_caught"`
	Extra       map[string]int `json:"-"`
}
type ExperimentationGame struct {
	LastAttempt int64            `json:"last_attempt,omitempty"`
	LastClaimed int64            `json:"last_claimed,omitempty"`
	BonusClicks int              `json:"bonus_clicks,omitempty"`
	Claimed     bool             `json:"claimed,omitempty"`
	Attempts    map[int]int      `json:"-"`
	Claims      map[int]int      `json:"-"`
	BestScores  map[int]int      `json:"-"`
	Raw         map[string]int64 `json:"-"`
}

func (e *ExperimentationGame) UnmarshalJSON(data []byte) error {
	type Alias ExperimentationGame
	aux := &struct {
		*Alias
	}{Alias: (*Alias)(e)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	e.Attempts = make(map[int]int)
	e.Claims = make(map[int]int)
	e.BestScores = make(map[int]int)
	e.Raw = make(map[string]int64)

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for k, v := range raw {
		switch {
		case len(k) > 9 && k[:9] == "attempts_":
			var idx int
			fmt.Sscanf(k, "attempts_%d", &idx)
			e.Attempts[idx] = int(v.(float64))
		case len(k) > 7 && k[:7] == "claims_":
			var idx int
			fmt.Sscanf(k, "claims_%d", &idx)
			e.Claims[idx] = int(v.(float64))
		case len(k) > 11 && k[:11] == "best_score_":
			var idx int
			fmt.Sscanf(k, "best_score_%d", &idx)
			e.BestScores[idx] = int(v.(float64))
		}
	}
	return nil
}

type experimentationData struct {
	Simon                 *ExperimentationGame `json:"simon,omitempty"`
	Pairings              *ExperimentationGame `json:"pairings,omitempty"`
	Numbers               *ExperimentationGame `json:"numbers,omitempty"`
	ClaimsResets          int64                `json:"claims_resets,omitempty"`
	ClaimsResetsTimestamp int64                `json:"claims_resets_timestamp,omitempty"`
	SerumsDrank           int                  `json:"serums_drank,omitempty"`
	ClaimedRetroactiveRng bool                 `json:"claimed_retroactive_rng,omitempty"`
	ChargeTrackTimestamp  int64                `json:"charge_track_timestamp,omitempty"`
}

type Dungeons struct {
	DungeonTypes         map[string]DungeonData `json:"dungeon_types,omitempty"`
	Classes              map[string]playerClass `json:"player_classes,omitempty"`
	SelectedDungeonClass string                 `json:"selected_dungeon_class,omitempty"`
	Secrets              float64                `json:"secrets,omitempty"`
}

type playerClass struct {
	Experience float64 `json:"experience,omitempty"`
}

type DungeonData struct {
	Experience float64 `json:"experience,omitempty"`

	HighestTierCompleted int                   `json:"highest_tier_completed,omitempty"`
	TimesPlayed          map[string]float64    `json:"times_played,omitempty"`
	TierCompletions      map[string]float64    `json:"tier_completions,omitempty"`
	MilestoneCompletions map[string]float64    `json:"milestone_completions,omitempty"`
	MobsKilled           map[string]float64    `json:"mobs_killed,omitempty"`
	MostMobsKilled       map[string]float64    `json:"most_mobs_killed,omitempty"`
	WatcherKills         map[string]float64    `json:"watcher_kills,omitempty"`
	MostDamageBerserk    map[string]float64    `json:"most_damage_berserk,omitempty"`
	MostDamageMage       map[string]float64    `json:"most_damage_mage,omitempty"`
	MostDamageHealer     map[string]float64    `json:"most_damage_healer,omitempty"`
	MostDamageArcher     map[string]float64    `json:"most_damage_archer,omitempty"`
	MostDamageTank       map[string]float64    `json:"most_damage_tank,omitempty"`
	MostHealing          map[string]float64    `json:"most_healing,omitempty"`
	FastestTime          map[string]float64    `json:"fastest_time,omitempty"`
	FastestTimeS         map[string]float64    `json:"fastest_time_s,omitempty"`
	FastestTimeSPlus     map[string]float64    `json:"fastest_time_s_plus,omitempty"`
	BestScore            map[string]float64    `json:"best_score,omitempty"`
	BestRuns             map[string]*[]BestRun `json:"best_runs,omitempty"`
}

type BestRun struct {
	Timestamp        int64   `json:"timestamp"`
	ScoreExploration int     `json:"score_exploration"`
	ScoreSpeed       int     `json:"score_speed"`
	ScoreSkill       int     `json:"score_skill"`
	ScoreBonus       int     `json:"score_bonus"`
	DungeonClass     string  `json:"dungeon_class"`
	ElapsedTime      int64   `json:"elapsed_time"`
	DamageDealt      float64 `json:"damage_dealt"`
	Deaths           int     `json:"deaths"`
	MobsKilled       int     `json:"mobs_killed"`
	SecretsFound     int     `json:"secrets_found"`
	DamageMitigated  float64 `json:"damage_mitigated"`
}

type slayer struct {
	SlayerBosses map[string]SlayerBoss `json:"slayer_bosses,omitempty"`
}

type SlayerBoss struct {
	BossKillsTier0    int     `json:"boss_kills_tier_0,omitempty"`
	BossKillsTier1    int     `json:"boss_kills_tier_1,omitempty"`
	BossKillsTier2    int     `json:"boss_kills_tier_2,omitempty"`
	BossKillsTier3    int     `json:"boss_kills_tier_3,omitempty"`
	BossKillsTier4    int     `json:"boss_kills_tier_4,omitempty"`
	BossAttemptsTier0 int     `json:"boss_attempts_tier_0,omitempty"`
	BossAttemptsTier1 int     `json:"boss_attempts_tier_1,omitempty"`
	BossAttemptsTier2 int     `json:"boss_attempts_tier_2,omitempty"`
	BossAttemptsTier3 int     `json:"boss_attempts_tier_3,omitempty"`
	BossAttemptsTier4 int     `json:"boss_attempts_tier_4,omitempty"`
	Experience        float64 `json:"xp,omitempty"`
}
