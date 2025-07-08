package constants

type Rank struct {
	Color     string `json:"color"`
	Tag       string `json:"tag"`
	Plus      string `json:"plus,omitempty"`
	PlusColor string `json:"plusColor,omitempty"`
}

var RANKS = map[string]Rank{
	"STAFF": {
		Color: "#c43c3c",
		Tag:   "ዞ",
	},
	"OWNER": {
		Color: "#c43c3c",
		Tag:   "OWNER",
	},
	"ADMIN": {
		Color: "#c43c3c",
		Tag:   "ADMIN",
	},
	"GAME_MASTER": {
		Color: "#00aa00",
		Tag:   "GM",
	},
	"YOUTUBER": {
		Color: "#c43c3c",
		Tag:   "YOUTUBE",
	},
	"SUPERSTAR": {
		Color:     "#d88f07",
		Tag:       "MVP",
		Plus:      "++",
		PlusColor: "#c43c3c",
	},
	"MVP_PLUS": {
		Color:     "#33aec3",
		Tag:       "MVP",
		Plus:      "+",
		PlusColor: "#c43c3c",
	},
	"MVP": {
		Color: "#33aec3",
		Tag:   "MVP",
	},
	"VIP_PLUS": {
		Color:     "#40bb40",
		Tag:       "VIP",
		Plus:      "+",
		PlusColor: "#d88f07",
	},
	"VIP": {
		Color: "#40bb40",
		Tag:   "VIP",
	},
	"PIG+++": {
		Color:     "#e668c6",
		Tag:       "PIG",
		Plus:      "+++",
		PlusColor: "#33aec3",
	},
	"MAYOR": {
		Color: "#e668c6",
		Tag:   "MAYOR",
	},
	"MINISTER": {
		Color: "#c43c3c",
		Tag:   "MINISTER",
	},
}

var RANK_PLUS_COLORS = map[string]string{
	"BLACK":        "#000000",
	"DARK_BLUE":    "#0b277a",
	"DARK_GREEN":   "#00aa00",
	"DARK_AQUA":    "#038d8d",
	"DARK_RED":     "#920909",
	"DARK_PURPLE":  "#a305a3",
	"GOLD":         "#d88f07",
	"GRAY":         "#636363",
	"DARK_GRAY":    "#2f2f2f",
	"BLUE":         "#4444f3",
	"GREEN":        "#40bb40",
	"AQUA":         "#33aec3",
	"RED":          "#c43c3c",
	"LIGHT_PURPLE": "#e668c6",
	"YELLOW":       "#efc721",
	"WHITE":        "#929292",
}
