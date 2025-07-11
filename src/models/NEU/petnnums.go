package neu

type PetNums map[string]PetRarityData

type PetRarityData map[string]PetRarityInfo

type PetRarityInfo struct {
	StatsLevellingCurve string    `json:"stats_levelling_curve"`
	Level1              *PetLevel `json:"1"`
	Level100            *PetLevel `json:"100"`
}

// PetLevel represents stats and other numbers for a specific pet level
type PetLevel struct {
	OtherNums []float64          `json:"otherNums"`
	StatNums  map[string]float64 `json:"statNums"`
}
