package constants

var DEFAULT_SKILL_CAPS = map[string]int{
	"farming":       50,
	"mining":        60,
	"combat":        60,
	"foraging":      50,
	"fishing":       50,
	"enchanting":    60,
	"alchemy":       50,
	"taming":        50,
	"carpentry":     50,
	"runecrafting":  25,
	"social":        25,
	"dungeoneering": 50,
}

const NON_RUNECRAFTING_LEVEL_CAP = 3

var MaxedSkillCaps = map[string]int{
	"farming": 60,
	"taming":  60,
}

var RUNECRAFTING_XP = map[int]int{
	1: 50, 2: 100, 3: 125, 4: 160, 5: 200,
	6: 250, 7: 315, 8: 400, 9: 500, 10: 625,
	11: 785, 12: 1000, 13: 1250, 14: 1600, 15: 2000,
	16: 2465, 17: 3125, 18: 4000, 19: 5000, 20: 6200,
	21: 7800, 22: 9800, 23: 12200, 24: 15300, 25: 19050,
}

var DUNGEONEERING_XP = map[int]int{
	1: 50, 2: 75, 3: 110, 4: 160, 5: 230,
	6: 330, 7: 470, 8: 670, 9: 950, 10: 1340,
	11: 1890, 12: 2665, 13: 3760, 14: 5260, 15: 7380,
	16: 10300, 17: 14400, 18: 20000, 19: 27600, 20: 38000,
	21: 52500, 22: 71500, 23: 97000, 24: 132000, 25: 180000,
	26: 243000, 27: 328000, 28: 445000, 29: 600000, 30: 800000,
	31: 1065000, 32: 1410000, 33: 1900000, 34: 2500000, 35: 3300000,
	36: 4300000, 37: 5600000, 38: 7200000, 39: 9200000, 40: 12000000,
	41: 15000000, 42: 19000000, 43: 24000000, 44: 30000000, 45: 38000000,
	46: 48000000, 47: 60000000, 48: 75000000, 49: 93000000, 50: 116250000,
	51: 200000000,
}

var SOCIAL_XP = map[int]int{
	1: 50, 2: 100, 3: 150, 4: 250, 5: 500,
	6: 750, 7: 1000, 8: 1250, 9: 1500, 10: 2000,
	11: 2500, 12: 3000, 13: 3750, 14: 4500, 15: 6000,
	16: 8000, 17: 10000, 18: 12500, 19: 15000, 20: 20000,
	21: 25000, 22: 30000, 23: 35000, 24: 40000, 25: 50000,
}

var DEFAULT_LEVELLING_XP = map[int]int{
	1: 50, 2: 125, 3: 200, 4: 300, 5: 500,
	6: 750, 7: 1000, 8: 1500, 9: 2000, 10: 3500,
	11: 5000, 12: 7500, 13: 10000, 14: 15000, 15: 20000,
	16: 30000, 17: 50000, 18: 75000, 19: 100000, 20: 200000,
	21: 300000, 22: 400000, 23: 500000, 24: 600000, 25: 700000,
	26: 800000, 27: 900000, 28: 1000000, 29: 1100000, 30: 1200000,
	31: 1300000, 32: 1400000, 33: 1500000, 34: 1600000, 35: 1700000,
	36: 1800000, 37: 1900000, 38: 2000000, 39: 2100000, 40: 2200000,
	41: 2300000, 42: 2400000, 43: 2500000, 44: 2600000, 45: 2750000,
	46: 2900000, 47: 3100000, 48: 3400000, 49: 3700000, 50: 4000000,
	51: 4300000, 52: 4600000, 53: 4900000, 54: 5200000, 55: 5500000,
	56: 5800000, 57: 6100000, 58: 6400000, 59: 6700000, 60: 7000000,
}

var HOTM_XP = map[int]int{
	1: 0, 2: 3000, 3: 9000, 4: 25000, 5: 60000,
	6: 100000, 7: 150000, 8: 210000, 9: 290000, 10: 400000,
}

var SKYBLOCK_XP = map[int]int{
	1: 100,
}

var COSMETIC_SKILLS = []string{"runecrafting", "social"}

var INFINITE = []string{"dungeoneering", "skyblock_level"}

var SKILL_ICONS = map[string]string{
	"skyblock_level": "http://localhost:8080/api/head/2e2cc42015e6678f8fd49ccc01fbf787f1ba2c32bcf559a015332fc5db50",
	"farming":        "http://localhost:8080/api/item/GOLDEN_HOE",
	"combat":         "http://localhost:8080/api/item/STONE_SWORD",
	"fishing":        "http://localhost:8080/api/item/FISHING_ROD",
	"alchemy":        "http://localhost:8080/api/item/BREWING_STAND",
	"runecrafting":   "http://localhost:8080/api/item/MAGMA_CREAM",
	"taming":         "http://localhost:8080/api/item/SPAWN_EGG",
	"mining":         "http://localhost:8080/api/item/STONE_PICKAXE",
	"foraging":       "http://localhost:8080/api/item/SAPLING:3",
	"enchanting":     "http://localhost:8080/api/item/ENCHANTING_TABLE",
	"carpentry":      "http://localhost:8080/api/item/CRAFTING_TABLE",
	"social":         "http://localhost:8080/api/item/EMERALD",
	"dungeoneering":  "http://localhost:8080/api/head/964e1c3e315c8d8fffc37985b6681c5bd16a6f97ffd07199e8a05efbef103793",
	"healer":         "http://localhost:8080/api/potion/normal/f52423",
	"mage":           "http://localhost:8080/api/item/BLAZE_ROD",
	"archer":         "http://localhost:8080/api/item/BOW",
	"berserk":        "http://localhost:8080/api/item/IRON_SWORD",
	"tank":           "http://localhost:8080/api/leather/chestplate/955e3b",
	"garden":         "http://localhost:8080/api/item/DOUBLE_PLANT",
}
