package constants

var FARMING_MEDALS = []string{"bronze", "silver", "gold", "platinum", "diamond"}

var CROPS = map[string]string{
	"INK_SACK:3":          "Cocoa Beans",
	"POTATO_ITEM":         "Potato",
	"CARROT_ITEM":         "Carrot",
	"CACTUS":              "Cactus",
	"SUGAR_CANE":          "Sugar Cane",
	"MUSHROOM_COLLECTION": "Mushroom",
	"PUMPKIN":             "Pumpkin",
	"NETHER_STALK":        "Nether Wart",
	"WHEAT":               "Wheat",
	"MELON":               "Melon",
}

var CROP_TO_ID = map[string]string{
	"WHEAT":               "WHEAT",
	"CARROT_ITEM":         "CARROT",
	"POTATO_ITEM":         "POTATO",
	"MELON":               "MELON",
	"PUMPKIN":             "PUMPKIN",
	"SUGAR_CANE":          "SUGAR_CANE",
	"CACTUS":              "CACTUS",
	"MUSHROOM_COLLECTION": "MUSHROOM",
	"NETHER_STALK":        "NETHER_WART",
	"INK_SACK:3":          "COCOA_BEANS",
}
