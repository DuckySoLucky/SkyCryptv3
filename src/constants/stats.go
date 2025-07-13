package constants

import "sort"

type statData struct {
	Name      string `json:"name"`
	NameLore  string `json:"nameLore"`
	NameShort string `json:"nameShort"`
	NameTiny  string `json:"nameTiny"`
	Symbol    string `json:"symbol"`
	Suffix    string `json:"suffix"`
	Color     string `json:"color"`
	Percent   bool   `json:"percent,omitempty"`
}

type StatsData map[string]statData

var STATS_DATA = StatsData{
	"health": {
		Name:      "Health",
		NameLore:  "Health",
		NameShort: "Health",
		NameTiny:  "HP",
		Symbol:    "❤",
		Suffix:    "",
		Color:     "text-minecraft-c",
	},
	"defense": {
		Name:      "Defense",
		NameLore:  "Defense",
		NameShort: "Defense",
		NameTiny:  "Def",
		Symbol:    "❈",
		Suffix:    "",
		Color:     "text-minecraft-a",
	},
	"strength": {
		Name:      "Strength",
		NameLore:  "Strength",
		NameShort: "Strength",
		NameTiny:  "Str",
		Symbol:    "❁",
		Suffix:    "",
		Color:     "text-minecraft-c",
	},
	"speed": {
		Name:      "Speed",
		NameLore:  "Speed",
		NameShort: "Speed",
		NameTiny:  "Spd",
		Symbol:    "✦",
		Suffix:    "",
		Color:     "text-minecraft-f",
	},
	"critical_chance": {
		Name:      "Crit Chance",
		NameLore:  "Crit Chance",
		NameShort: "Crit Chance",
		NameTiny:  "CC",
		Symbol:    "☣",
		Suffix:    "%",
		Color:     "text-minecraft-9",
		Percent:   true,
	},
	"critical_damage": {
		Name:      "Crit Damage",
		NameLore:  "Crit Damage",
		NameShort: "Crit Damage",
		NameTiny:  "CD",
		Symbol:    "☠",
		Suffix:    "%",
		Color:     "text-minecraft-9",
		Percent:   true,
	},
	"intelligence": {
		Name:      "Intelligence",
		NameLore:  "Intelligence",
		NameShort: "Intelligence",
		NameTiny:  "Int",
		Symbol:    "✎",
		Suffix:    "",
		Color:     "text-minecraft-b",
	},
	"bonus_attack_speed": {
		Name:      "Bonus Attack Speed",
		NameLore:  "Bonus Attack Speed",
		NameShort: "Attack Speed",
		NameTiny:  "Atk",
		Symbol:    "⚔",
		Suffix:    "%",
		Color:     "text-minecraft-e",
		Percent:   true,
	},
	"sea_creature_chance": {
		Name:      "Sea Creature Chance",
		NameLore:  "Sea Creature Chance",
		NameShort: "SC Chance",
		NameTiny:  "SCC",
		Symbol:    "α",
		Suffix:    "%",
		Color:     "text-minecraft-3",
		Percent:   true,
	},
	"magic_find": {
		Name:      "Magic Find",
		NameLore:  "Magic Find",
		NameShort: "Magic Find",
		NameTiny:  "MF",
		Symbol:    "✯",
		Suffix:    "",
		Color:     "text-minecraft-b",
	},
	"pet_luck": {
		Name:      "Pet Luck",
		NameLore:  "Pet Luck",
		NameShort: "Pet Luck",
		NameTiny:  "PL",
		Symbol:    "♣",
		Suffix:    "",
		Color:     "text-minecraft-d",
	},
	"true_defense": {
		Name:      "True Defense",
		NameLore:  "True Defense",
		NameShort: "True Defense",
		NameTiny:  "TD",
		Symbol:    "❂",
		Suffix:    "",
		Color:     "text-minecraft-f",
	},
	"ferocity": {
		Name:      "Ferocity",
		NameLore:  "Ferocity",
		NameShort: "Ferocity",
		NameTiny:  "Frc",
		Symbol:    "⫽",
		Suffix:    "",
		Color:     "text-minecraft-c",
	},
	"ability_damage": {
		Name:      "Ability Damage",
		NameLore:  "Ability Damage",
		NameShort: "Ability Damage",
		NameTiny:  "AD",
		Symbol:    "๑",
		Suffix:    "%",
		Color:     "text-minecraft-c",
		Percent:   true,
	},
	"mining_speed": {
		Name:      "Mining Speed",
		NameLore:  "Mining Speed",
		NameShort: "Mining Speed",
		NameTiny:  "MngSpd",
		Symbol:    "⸕",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"mining_fortune": {
		Name:      "Mining Fortune",
		NameLore:  "Mining Fortune",
		NameShort: "Mining Fortune",
		NameTiny:  "MngFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"ore_fortune": {
		Name:      "Ore Fortune",
		NameLore:  "Ore Fortune",
		NameShort: "Ore Fortune",
		NameTiny:  "OreFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"farming_fortune": {
		Name:      "Farming Fortune",
		NameLore:  "Farming Fortune",
		NameShort: "Farming Fortune",
		NameTiny:  "FrmFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"foraging_fortune": {
		Name:      "Foraging Fortune",
		NameLore:  "Foraging Fortune",
		NameShort: "Foraging Fortune",
		NameTiny:  "FrgFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"pristine": {
		Name:      "Pristine",
		NameLore:  "Pristine",
		NameShort: "Pristine",
		NameTiny:  "Prs",
		Symbol:    "✧",
		Suffix:    "",
		Color:     "text-minecraft-5",
	},
	"fishing_speed": {
		Name:      "Fishing Speed",
		NameLore:  "Fishing Speed",
		NameShort: "Fishing Speed",
		NameTiny:  "FS",
		Symbol:    "☂",
		Suffix:    "",
		Color:     "text-minecraft-b",
	},
	"breaking_power": {
		Name:      "Breaking Power",
		NameLore:  "Breaking Power",
		NameShort: "Breaking Power",
		NameTiny:  "BP",
		Symbol:    "Ⓟ",
		Suffix:    "",
		Color:     "text-minecraft-2",
	},
	"health_regen": {
		Name:      "Health Regen",
		NameLore:  "Health Regen",
		NameShort: "Health Regen",
		NameTiny:  "HPR",
		Symbol:    "❣",
		Suffix:    "",
		Color:     "text-minecraft-c",
	},
	"vitality": {
		Name:      "Vitality",
		NameLore:  "Vitality",
		NameShort: "Vitality",
		NameTiny:  "Vit",
		Symbol:    "♨",
		Suffix:    "",
		Color:     "text-minecraft-5",
	},
	"mending": {
		Name:      "Mending",
		NameLore:  "Mending",
		NameShort: "Mending",
		NameTiny:  "Mend",
		Symbol:    "☄",
		Suffix:    "",
		Color:     "text-minecraft-a",
	},
	"mana_regen": {
		Name:      "Mana Regen",
		NameLore:  "Mana Regen",
		NameShort: "Mana Regen",
		NameTiny:  "MPR",
		Symbol:    "🗲",
		Suffix:    "",
		Color:     "text-minecraft-b",
	},
	"rift_time": {
		Name:      "Rift Time",
		NameLore:  "Rift Time",
		NameShort: "Rift Time",
		NameTiny:  "RT",
		Symbol:    "ф",
		Suffix:    "",
		Color:     "text-minecraft-a",
	},
	"alchemy_wisdom": {
		Name:      "Alchemy Wisdom",
		NameLore:  "Alchemy Wisdom",
		NameShort: "Alchemy Wisdom",
		NameTiny:  "AW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"carpentry_wisdom": {
		Name:      "Carpentry Wisdom",
		NameLore:  "Carpentry Wisdom",
		NameShort: "Carpentry Wisdom",
		NameTiny:  "CW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"combat_wisdom": {
		Name:      "Combat Wisdom",
		NameLore:  "Combat Wisdom",
		NameShort: "Combat Wisdom",
		NameTiny:  "CW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"enchanting_wisdom": {
		Name:      "Enchanting Wisdom",
		NameLore:  "Enchanting Wisdom",
		NameShort: "Enchanting Wisdom",
		NameTiny:  "EW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"farming_wisdom": {
		Name:      "Farming Wisdom",
		NameLore:  "Farming Wisdom",
		NameShort: "Farming Wisdom",
		NameTiny:  "FW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"fishing_wisdom": {
		Name:      "Fishing Wisdom",
		NameLore:  "Fishing Wisdom",
		NameShort: "Fishing Wisdom",
		NameTiny:  "FW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"foraging_wisdom": {
		Name:      "Foraging Wisdom",
		NameLore:  "Foraging Wisdom",
		NameShort: "Foraging Wisdom",
		NameTiny:  "FW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"mining_wisdom": {
		Name:      "Mining Wisdom",
		NameLore:  "Mining Wisdom",
		NameShort: "Mining Wisdom",
		NameTiny:  "MW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"runecrafting_wisdom": {
		Name:      "Runecrafting Wisdom",
		NameLore:  "Runecrafting Wisdom",
		NameShort: "Runecrafting Wisdom",
		NameTiny:  "RW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"social_wisdom": {
		Name:      "Social Wisdom",
		NameLore:  "Social Wisdom",
		NameShort: "Social Wisdom",
		NameTiny:  "SW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"taming_wisdom": {
		Name:      "Taming Wisdom",
		NameLore:  "Taming Wisdom",
		NameShort: "Taming Wisdom",
		NameTiny:  "TW",
		Symbol:    "☯",
		Suffix:    "",
		Color:     "text-minecraft-3",
	},
	"bonus_pest_chance": {
		Name:      "Bonus Pest Chance",
		NameLore:  "Bonus Pest Chance",
		NameShort: "Bonus Pest Chance",
		NameTiny:  "BPC",
		Symbol:    "ൠ",
		Suffix:    "",
		Color:     "text-minecraft-2",
	},
	"swing_range": {
		Name:      "Swing Range",
		NameLore:  "Swing Range",
		NameShort: "Swing Range",
		NameTiny:  "SR",
		Symbol:    "Ⓢ",
		Suffix:    "",
		Color:     "text-minecraft-e",
	},
	"cold_resistance": {
		Name:      "Cold Resistance",
		NameLore:  "Cold Resistance",
		NameShort: "Cold Resistance",
		NameTiny:  "CRes",
		Symbol:    "❄",
		Suffix:    "",
		Color:     "text-minecraft-b",
	},
	"mining_spread": {
		Name:      "Mining Spread",
		NameLore:  "Mining Spread",
		NameShort: "",
		NameTiny:  "MS",
		Symbol:    "▚",
		Suffix:    "",
		Color:     "text-minecraft-e",
	},
	"gemstone_spread": {
		Name:      "Gemstone Spread",
		NameLore:  "Gemstone Spread",
		NameShort: "",
		NameTiny:  "GS",
		Symbol:    "▚",
		Suffix:    "",
		Color:     "text-minecraft-e",
	},
	"block_fortune": {
		Name:      "Block Fortune",
		NameLore:  "Block Fortune",
		NameShort: "",
		NameTiny:  "BFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"dwarven_metal_fortune": {
		Name:      "Dwarven Metal Fortune",
		NameLore:  "Dwarven Metal Fortune",
		NameShort: "",
		NameTiny:  "DMFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"gemstone_fortune": {
		Name:      "Gemstone Fortune",
		NameLore:  "Gemstone Fortune",
		NameShort: "",
		NameTiny:  "GFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"wheat_fortune": {
		Name:      "Wheat Fortune",
		NameLore:  "Wheat Fortune",
		NameShort: "",
		NameTiny:  "WFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"carrot_fortune": {
		Name:      "Carrot Fortune",
		NameLore:  "Carrot Fortune",
		NameShort: "",
		NameTiny:  "CFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"potato_fortune": {
		Name:      "Potato Fortune",
		NameLore:  "Potato Fortune",
		NameShort: "",
		NameTiny:  "PFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"pumpkin_fortune": {
		Name:      "Pumpkin Fortune",
		NameLore:  "Pumpkin Fortune",
		NameShort: "",
		NameTiny:  "PkFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"melon_fortune": {
		Name:      "Melon Fortune",
		NameLore:  "Melon Fortune",
		NameShort: "",
		NameTiny:  "MFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"mushroom_fortune": {
		Name:      "Mushroom Fortune",
		NameLore:  "Mushroom Fortune",
		NameShort: "",
		NameTiny:  "MsFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"cactus_fortune": {
		Name:      "Cactus Fortune",
		NameLore:  "Cactus Fortune",
		NameShort: "",
		NameTiny:  "CFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"sugar_cane_fortune": {
		Name:      "Sugar Cane Fortune",
		NameLore:  "Sugar Cane Fortune",
		NameShort: "",
		NameTiny:  "SCFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"nether_wart_fortune": {
		Name:      "Nether Wart Fortune",
		NameLore:  "Nether Wart Fortune",
		NameShort: "",
		NameTiny:  "NWFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"cocoa_beans_fortune": {
		Name:      "Cocoa Beans Fortune",
		NameLore:  "Cocoa Beans Fortune",
		NameShort: "",
		NameTiny:  "CBFrt",
		Symbol:    "☘",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"double_hook_chance": {
		Name:      "Double Hook Chance",
		NameLore:  "Double Hook Chance",
		NameShort: "",
		NameTiny:  "DHC",
		Symbol:    "⚓",
		Suffix:    "",
		Color:     "text-minecraft-9",
	},
	"trophy_fish_chance": {
		Name:      "Trophy Fish Chance",
		NameLore:  "Trophy Fish Chance",
		NameShort: "",
		NameTiny:  "TFC",
		Symbol:    "♔",
		Suffix:    "",
		Color:     "text-minecraft-6",
	},
	"heat_resistance": {
		Name:      "Heat Resistance",
		NameLore:  "Heat Resistance",
		NameShort: "",
		NameTiny:  "HRes",
		Symbol:    "♨",
		Suffix:    "",
		Color:     "text-minecraft-c",
	},
	"fear": {
		Name:      "Fear",
		NameLore:  "Fear",
		NameShort: "",
		NameTiny:  "Fr",
		Symbol:    "☠",
		Suffix:    "",
		Color:     "text-minecraft-5",
	},
	"damage": {
		Name:      "Damage",
		NameLore:  "Damage",
		NameShort: "",
		NameTiny:  "DMG",
		Symbol:    "❁",
		Suffix:    "",
		Color:     "text-minecraft-c",
	},
}

var STATS_BONUS = map[string]map[int]map[string]float64{
	// Skills
	"skill_farming": {
		1:  {"health": 2, "farming_fortune": 4},
		15: {"health": 3, "farming_fortune": 4},
		20: {"health": 4, "farming_fortune": 4},
		26: {"health": 5, "farming_fortune": 4},
	},
	"skill_mining": {
		1:  {"defense": 1, "mining_fortune": 4},
		15: {"defense": 2, "mining_fortune": 4},
	},
	"skill_combat": {
		1: {"critical_chance": 0.5},
	},
	"skill_foraging": {
		1:  {"strength": 1, "foraging_fortune": 4},
		15: {"strength": 2, "foraging_fortune": 4},
	},
	"skill_fishing": {
		1:  {"health": 2},
		15: {"health": 3},
		20: {"health": 4},
		26: {"health": 5},
	},
	"skill_enchanting": {
		1:  {"intelligence": 1, "ability_damage": 0.5},
		15: {"intelligence": 2, "ability_damage": 0.5},
	},
	"skill_alchemy": {
		1:  {"intelligence": 1},
		15: {"intelligence": 2},
	},
	"skill_taming": {
		1: {"pet_luck": 1},
	},
	"skill_dungeoneering": {
		1:  {"health": 2},
		51: {"health": 0},
	},
	"skill_social":       {},
	"skill_carpentry":    {1: {"health": 1}},
	"skill_runecrafting": {},
	// Slayers
	"slayer_zombie": {
		1: {"health": 2},
		3: {"health": 3},
		5: {"health": 4},
		7: {"health": 5},
		8: {"health": 5, "health_regen": 50},
		9: {"health": 6},
	},
	"slayer_spider": {
		1: {"critical_damage": 1},
		5: {"critical_damage": 2},
		7: {"critical_damage": 2},
		8: {"critical_damage": 3},
	},
	"slayer_wolf": {
		1: {"speed": 1},
		2: {"health": 2},
		3: {"speed": 1},
		4: {"health": 2},
		5: {"critical_damage": 1},
		6: {"health": 3},
		7: {"critical_damage": 2},
		8: {"speed": 1},
		9: {"health": 5},
	},
	"slayer_enderman": {
		1: {"health": 1},
		2: {"intelligence": 2},
		3: {"health": 2},
		4: {"true_defense": 1},
		5: {"health": 3},
		6: {"intelligence": 5},
		7: {"health": 4},
		8: {"intelligence": 4},
		9: {"health": 5},
	},
	"slayer_blaze": {
		1: {"health": 3},
		2: {"strength": 1},
		3: {"health": 4},
		4: {"true_defense": 1},
		5: {"health": 5},
		6: {"strength": 2},
		7: {"health": 6},
		8: {"true_defense": 2},
		9: {"health": 7},
	},
}

func GetBonusStats(level int, statsBonus map[int]map[string]float64) map[string]float64 {
	bonus := make(map[string]float64)
	if statsBonus == nil {
		return bonus
	}

	steps := make([]int, 0, len(statsBonus))
	for k := range statsBonus {
		steps = append(steps, k)
	}
	sort.Ints(steps)

	if len(steps) == 0 {
		return bonus
	}

	for x := steps[0]; x <= len(statsBonus); x++ {
		if level < x {
			break
		}

		step := steps[0]
		for _, s := range steps {
			if s <= x {
				step = s
			}
		}
		stepBonuses := statsBonus[step]
		for statName, value := range stepBonuses {
			bonus[statName] += value
		}
	}
	return bonus
}

func GetBonusStat(level int, key string, max int) map[string]float64 {
	bonus := make(map[string]float64)
	objOfLevelBonuses, ok := STATS_BONUS[key]
	if !ok {
		return bonus
	}

	steps := make([]int, 0, len(objOfLevelBonuses))
	for k := range objOfLevelBonuses {
		steps = append(steps, k)
	}
	sort.Ints(steps)

	if len(steps) == 0 {
		return bonus
	}

	for x := steps[0]; x <= max; x++ {
		if level < x {
			break
		}

		step := steps[0]
		for i := len(steps) - 1; i >= 0; i-- {
			if steps[i] <= x {
				step = steps[i]
				break
			}
		}
		stepBonuses := objOfLevelBonuses[step]
		for statName, value := range stepBonuses {
			bonus[statName] += value
		}
	}
	return bonus
}
