package stats

import (
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"slices"
)

type equipmentResult struct {
	Equipment []models.ProcessedItem `json:"equipment"`
	Stats     map[string]float64     `json:"stats"`
}

func GetEquipment(equipment []models.ProcessedItem) equipmentResult {
	if utility.Every(equipment, isInvalidItem) {
		return equipmentResult{
			Equipment: []models.ProcessedItem{},
			Stats:     map[string]float64{},
		}
	}

	reversedEquipment := make([]models.ProcessedItem, len(equipment))
	copy(reversedEquipment, equipment)
	slices.Reverse(reversedEquipment)

	return equipmentResult{
		Equipment: reversedEquipment,
		Stats:     GetStatsFromItems(equipment),
	}
}
