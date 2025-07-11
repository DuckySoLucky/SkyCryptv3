package constants

var PET_REWARDS = map[int]map[string]float64{
	0:   {"magic_find": 0},
	10:  {"magic_find": 1},
	25:  {"magic_find": 2},
	50:  {"magic_find": 3},
	75:  {"magic_find": 4},
	100: {"magic_find": 5},
	130: {"magic_find": 6},
	175: {"magic_find": 7},
	225: {"magic_find": 8},
	275: {"magic_find": 9},
	325: {"magic_find": 10},
	375: {"magic_find": 11},
	450: {"magic_find": 12},
	500: {"magic_find": 13},
}

var PET_PARENTS = map[string]string{
	"DROPLET_WISP": "WISP",
	"FROST_WISP":   "WISP",
	"GLACIAL_WISP": "WISP",
}
