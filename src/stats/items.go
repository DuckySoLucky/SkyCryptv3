package stats

import (
	"fmt"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"strings"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

func GetItems(useProfile *models.Member, profileId string) (map[string][]models.Item, error) {
	encodedInventories := map[string]*string{
		"inventory":      &useProfile.Inventory.Inventory.Data,
		"enderchest":     &useProfile.Inventory.Enderchest.Data,
		"armor":          &useProfile.Inventory.Armor.Data,
		"equipment":      &useProfile.Inventory.Equipment.Data,
		"personal_vault": &useProfile.Inventory.PersonalVault.Data,
		"wardrobe":       &useProfile.Inventory.Wardrobe.Data,

		// rift
		"rift_inventory":  &useProfile.Rift.Inventory.Inventory.Data,
		"rift_enderchest": &useProfile.Rift.Inventory.Enderchest.Data,
		"rift_armor":      &useProfile.Rift.Inventory.Armor.Data,
		"rift_equipment":  &useProfile.Rift.Inventory.Equipment.Data,

		// bags
		"potion_bag":   &useProfile.Inventory.BagContents.PotionBag.Data,
		"talisman_bag": &useProfile.Inventory.BagContents.TalismanBag.Data,
		"fishing_bag":  &useProfile.Inventory.BagContents.FishingBag.Data,
		// "sacks_bag": &useProfile.Inventory.BagContents.SacksBag.Data,
		"quiver": &useProfile.Inventory.BagContents.Quiver.Data,
	}

	for backpackId, backpackData := range useProfile.Inventory.Backpack {
		encodedInventories[fmt.Sprintf("backpack_%s", backpackId)] = &backpackData.Data
	}

	for backpackIconId, backpackIconData := range useProfile.Inventory.BackpackIcons {
		encodedInventories[fmt.Sprintf("backpack_icon_%s", backpackIconId)] = &backpackIconData.Data
	}

	type result struct {
		inventoryId string
		items       []models.Item
		err         error
	}

	resultChan := make(chan result, len(encodedInventories))
	var wg sync.WaitGroup

	for inventoryId, inventoryData := range encodedInventories {
		wg.Add(1)
		go func(id string, data *string) {
			defer wg.Done()

			decodedInventory, err := utility.DecodeInventory(data)
			if err != nil {
				resultChan <- result{id, nil, err}
				return
			}

			resultChan <- result{id, decodedInventory.Items, nil}
		}(inventoryId, inventoryData)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	decodedInventory := make(map[string][]models.Item, len(encodedInventories))
	for res := range resultChan {
		if res.err != nil {
			fmt.Printf("Error decoding inventory %s: %v\n", res.inventoryId, res.err)
			continue
		}

		decodedInventory[res.inventoryId] = res.items
	}

	output := make(map[string][]models.Item)
	for inventoryId, items := range decodedInventory {
		if !strings.Contains(inventoryId, "backpack") {
			output[inventoryId] = items
		}

		if strings.HasPrefix(inventoryId, "backpack_") && !strings.Contains(inventoryId, "icon") {
			if output["backpack"] == nil {
				output["backpack"] = []models.Item{}
			}

			backpackIndex := strings.Split(inventoryId, "_")[1]
			backpackIcon, iconExists := decodedInventory[fmt.Sprintf("backpack_icon_%s", backpackIndex)]
			if iconExists && len(backpackIcon) > 0 {
				backpackIcon[0].ContainsItems = items

				output["backpack"] = append(output["backpack"], backpackIcon[0])
			} else {
				fmt.Printf("No icon found for backpack %s\n", backpackIndex)
			}
		}
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonData, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("Error marshaling items for caching: %v\n", err)
	} else {
		redis.Set(fmt.Sprintf("items:%s", profileId), string(jsonData), 5*60)
	}

	return output, nil
}
