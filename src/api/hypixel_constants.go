package api

import (
	"fmt"
	"io"
	"net/http"
	"skycrypt/src/constants"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func getSkyBlockItems() ([]models.HypixelItem, error) {
	cachedData, err := redis.Get("skyblock_items")
	if err == nil && cachedData != "" {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		var data models.HypixelItemsResponse
		err = json.Unmarshal([]byte(cachedData), &data)
		if err == nil {
			return data.Items, nil
		}
	}

	resp, err := http.Get("https://api.hypixel.net/v2/resources/skyblock/items")
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var data models.HypixelItemsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set("skyblock_items", string(body), 12*60*60) // Cache for 12 hours

	return data.Items, nil
}

func processItems(items *[]models.HypixelItem) map[string]models.ProcessedHypixelItem {
	processed := make(map[string]models.ProcessedHypixelItem)
	for _, item := range *items {
		if item.Rarity == "" {
			item.Rarity = "common"
		}

		processed[item.SkyBlockID] = models.ProcessedHypixelItem{
			SkyblockID:        item.SkyBlockID,
			Name:              item.Name,
			ItemId:            constants.BUKKIT_TO_ID[item.Material],
			Rarity:            strings.ToLower(item.Rarity),
			Damage:            item.Damage,
			Texture:           utility.GetSkinHash(item.Skin.Value),
			Category:          strings.ToLower(item.Category),
			Origin:            item.Origin,
			RiftTransferrable: item.RiftTransferrable,
		}
	}

	return processed

}

func LoadSkyBlockItems() error {
	timeNow := time.Now()
	items, err := getSkyBlockItems()
	if err != nil {
		return fmt.Errorf("failed to get SkyBlock items: %v", err)
	}

	constants.ITEMS = processItems(&items)

	fmt.Printf("[ITEMS] Loaded %d items in %s\n", len(constants.ITEMS), time.Since(timeNow))

	return nil
}
