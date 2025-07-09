package stats

import (
	"fmt"
	"skycrypt/src/models"
)

func GetWardrobe(wardrobeInventory []models.ProcessedItem) [][]*models.ProcessedItem {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("panic occurred: %v", r)
			fmt.Print("Recovered from panic in GetWardrobe: ", err)
		}
	}()

	wardrobeColumns := len(wardrobeInventory) / 4

	var wardrobe [][]*models.ProcessedItem
	for i := range wardrobeColumns {
		page := i / 9

		var wardrobeSlot []*models.ProcessedItem
		for j := range 4 {
			index := 36*page + (i % 9) + j*9

			if len(GetId(wardrobeInventory[index])) > 0 {
				wardrobeSlot = append(wardrobeSlot, &wardrobeInventory[index])
			} else {
				wardrobeSlot = append(wardrobeSlot, nil)
			}

		}

		hasItems := false
		for _, item := range wardrobeSlot {
			if item != nil {
				hasItems = true
				break
			}
		}

		if hasItems {
			wardrobe = append(wardrobe, wardrobeSlot)
		}
	}

	return wardrobe
}
