package notenoughupdates

import (
	"fmt"
	"os"
	"skycrypt/src/models"

	jsoniter "github.com/json-iterator/go"
)

var NEUConstants = models.NEUConstant{}

func GetItem(name string) (models.NEUItem, error) {
	itemsPath := "NotEnoughUpdates-REPO/items"
	if _, err := os.Stat(itemsPath); os.IsNotExist(err) {
		return models.NEUItem{}, fmt.Errorf("items directory does not exist: %w", err)
	}

	itemPath := fmt.Sprintf("%s/%s.json", itemsPath, name)
	data, err := os.ReadFile(itemPath)
	if err != nil {
		return models.NEUItem{}, fmt.Errorf("failed to read item file %s: %w", itemPath, err)
	}

	var item models.NEUItem
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &item)
	if err != nil {
		return models.NEUItem{}, fmt.Errorf("failed to unmarshal JSON from %s: %w", itemPath, err)
	}

	return item, nil
}
