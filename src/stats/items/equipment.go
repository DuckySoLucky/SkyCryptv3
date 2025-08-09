package stats

import (
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"slices"
)

func GetEquipment(equipment []models.ProcessedItem) models.EquipmentResult {
	if utility.Every(equipment, isInvalidItem) {
		return models.EquipmentResult{
			Equipment: []models.StrippedItem{},
			Stats:     map[string]float64{},
		}
	}

	reversedEquipment := make([]models.ProcessedItem, len(equipment))
	copy(reversedEquipment, equipment)
	slices.Reverse(reversedEquipment)

	return models.EquipmentResult{
		Equipment: StripItems(&reversedEquipment),
		Stats:     GetStatsFromItems(equipment),
	}

}
