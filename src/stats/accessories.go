package stats

import (
	notenoughupdates "skycrypt/src/NotEnoughUpdates"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	stats "skycrypt/src/stats/items"
	"skycrypt/src/utility"
	"slices"
	"sort"
	"strings"
)

func GetAccessories(useProfile *models.Member, items map[string][]models.Item) models.GetMissingAccessoresOutput {
	if items == nil {
		return models.GetMissingAccessoresOutput{}
	}

	talismanBag := items["talisman_bag"]
	accessoryIds := make([]models.AccessoryIds, 0)
	accessories := make([]models.InsertAccessory, 0)
	for _, item := range stats.ProcessItems(&talismanBag, "talisman_bag") {
		id := stats.GetId(item)
		if len(id) == 0 {
			continue
		}

		newAccessory := models.InsertAccessory{
			ProcessedItem: item,
			Id:            id,
			Rarity:        item.Rarity,
		}

		newAccessoryId := models.AccessoryIds{
			Id:     id,
			Rarity: item.Rarity,
		}

		accessories = append(accessories, newAccessory)
		accessoryIds = append(accessoryIds, newAccessoryId)
	}

	var processedItems = make(map[string][]models.ProcessedItem)
	inventoryKeys := []string{"inventory", "enderchest", "backpack"}
	for _, inventoryId := range inventoryKeys {
		inventoryData := items[inventoryId]
		if len(inventoryData) == 0 {
			continue
		}

		processedItems[inventoryId] = stats.ProcessItems(&inventoryData, inventoryId)
	}

	for _, inventoryId := range inventoryKeys {
		if processedItems[inventoryId] == nil {
			continue
		}

		for _, item := range processedItems[inventoryId] {
			if utility.Contains(item.Categories, "accessory") {
				id := stats.GetId(item)
				if len(id) == 0 {
					continue
				}

				item.Lore = append(item.Lore, "", "§7Inactive: §cNot in accessory bag")
				newAccessory := models.InsertAccessory{
					ProcessedItem: item,
					Id:            id,
					Rarity:        item.Rarity,
					IsInactive:    true,
				}

				accessories = append(accessories, newAccessory)
			}
		}
	}

	var activeAccessories []models.InsertAccessory
	for _, a := range accessories {
		if !a.IsInactive {
			activeAccessories = append(activeAccessories, a)
		}
	}

	// Process duplicates
	for _, accessory := range activeAccessories {
		id := accessory.Id
		rarity := accessory.Rarity

		var duplicates []models.InsertAccessory
		for _, a := range accessories {
			if constants.GetBaseIdFromAlias(a.Id) == constants.GetBaseIdFromAlias(id) {
				duplicates = append(duplicates, a)
			}
		}

		if len(duplicates) > 1 {
			for i, duplicate := range duplicates {
				duplicateRarity := duplicate.Rarity

				if utility.RarityNameToInt(duplicateRarity) < utility.RarityNameToInt(rarity) {
					if len(duplicate.Tag.Display.Lore) == 0 || !strings.Contains(duplicate.Tag.Display.Lore[len(duplicate.Tag.Display.Lore)-1], "§7Inactive: ") {
						duplicates[i].Tag.Display.Lore = append(duplicates[i].Tag.Display.Lore, "", "§7Inactive: §cFound a higher rarity accessory")
						// fmt.Printf("Marking accessory %s as inactive due to higher rarity %s\n", duplicate.Id, rarity)
					}
					duplicates[i].IsInactive = true
				} else if duplicate.Rarity == rarity {
					if len(duplicate.Tag.Display.Lore) == 0 || !strings.Contains(duplicate.Tag.Display.Lore[len(duplicate.Tag.Display.Lore)-1], "§7Inactive:") {
						duplicates[i].Tag.Display.Lore = append(duplicates[i].Tag.Display.Lore, "", "§7Inactive: §cFound a duplicate accessory")
						// fmt.Printf("Marking accessory %s as inactive due to duplicate\n", duplicate.Id)
					}
					duplicates[i].IsInactive = true
				}
			}

			// Update accessories slice with modified duplicates
			for j, acc := range accessories {
				for _, dup := range duplicates {
					if acc.Id == dup.Id && acc.Rarity == dup.Rarity {
						accessories[j] = dup
						break
					}
				}
			}

			// Check if all duplicates are inactive
			allInactive := true
			for _, dup := range duplicates {
				if !dup.IsInactive {
					allInactive = false
					break
				}
			}

			if allInactive {
				// Find and reactivate the current accessory
				for i, acc := range accessories {
					if acc.Id == id && acc.Rarity == rarity {
						accessories[i].IsInactive = false
						break
					}
				}
			}
		}
	}

	// Process upgrade accessories
	for _, accessory := range accessories {
		id := accessory.Id
		upgradeList := constants.GetUpgradeList(id)
		for _, upgrade := range upgradeList {
			if slices.Index(upgradeList, upgrade) < slices.Index(upgradeList, id) {
				for j, acc := range accessories {
					if acc.Id == upgrade {
						accessories[j].IsInactive = true
					}
				}
			}
		}
	}

	for i, accessory := range accessories {
		if accessory.IsInactive {
			if accessory.Source != "" && len(accessory.Tag.Display.Lore) > 0 {
				accessories[i].Tag.Display.Lore = append(accessories[i].Tag.Display.Lore, "", "§7Location: §c"+utility.TitleCase(accessory.Source))
			}
		}
	}

	if useProfile.Rift.Access.ConsumedPrism {
		riftPrismItem, _ := notenoughupdates.GetItem("RIFT_PRISM")

		itemId := 397
		processedItem := stats.ProcessItem(&models.Item{
			Tag:    &riftPrismItem.NBT,
			ID:     &itemId,
			Damage: &riftPrismItem.Damage,
		}, "Rift")

		accessoryIds = append(accessoryIds, models.AccessoryIds{Id: "RIFT_PRISM", Rarity: "rare"})
		accessories = append(accessories, models.InsertAccessory{
			ProcessedItem: processedItem,
			Id:            "RIFT_PRISM",
			Rarity:        "rare",
			IsInactive:    false,
		})
	}

	sort.Sort(itemSorter(accessories))
	output := models.AccessoriesOutput{
		Accessories:  accessories,
		AccessoryIds: accessoryIds,
	}

	return GetMissingAccessories(output, useProfile)
}

type itemSorter []models.InsertAccessory

func (s itemSorter) Len() int {
	return len(s)
}

func (s itemSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s itemSorter) Less(i, j int) bool {
	a, b := s[i], s[j]

	if a.Rarity != b.Rarity {
		aIndex := utility.IndexOf(constants.RARITIES, a.Rarity)
		bIndex := utility.IndexOf(constants.RARITIES, b.Rarity)
		return bIndex < aIndex
	}

	if a.Source == "inventory" && b.Source != "inventory" {
		return true
	}

	if a.Source != "inventory" && b.Source == "inventory" {
		return false
	}

	return strings.Compare(a.DisplayName, b.DisplayName) < 0
}
