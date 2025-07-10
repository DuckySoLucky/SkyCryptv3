package neu

type PetNums map[string]petRarities

type petRarities map[string]struct {
	petLevels
	StatsLevellingCurve *string `json:"stats_levelling_curve,omitempty"`
}

type petLevels map[string]petStats

type petStats struct {
	OtherNums []float64          `json:"otherNums"`
	StatNums  map[string]float64 `json:"statNums"`
}
