package notenoughupdates

import (
	"fmt"
	"os"
	neu "skycrypt/src/models/NEU"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func ParseNEURepository() error {
	timeNow := time.Now()

	constantsPath := "NotEnoughUpdates-REPO/constants"
	if _, err := os.Stat(constantsPath); os.IsNotExist(err) {
		return fmt.Errorf("constants directory does not exist: %w", err)
	}

	fmt.Println("[NOT-ENOUGH-UPDATES] Parsing NEU repository...")

	constants, err := os.ReadDir(constantsPath)
	if err != nil {
		return fmt.Errorf("failed to read constants directory: %w", err)
	}

	for _, constant := range constants {
		if constant.IsDir() {
			fmt.Printf("[NOT-ENOUGH-UPDATES] Skipping directory: %s\n", constant.Name())
			continue
		}

		if constant.Name() == "petnums.json" {
			filePath := fmt.Sprintf("%s/%s", constantsPath, constant.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			var petNums neu.PetNums
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			err = json.Unmarshal(data, &petNums)
			if err != nil {
				return fmt.Errorf("failed to unmarshal JSON from %s: %w", filePath, err)
			}

			NEUConstants.PetNums = petNums
		} else if constant.Name() == "pets.json" {
			filePath := fmt.Sprintf("%s/%s", constantsPath, constant.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			var pets neu.Pets
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			err = json.Unmarshal(data, &pets)
			if err != nil {
				return fmt.Errorf("failed to unmarshal JSON from %s: %w", filePath, err)
			}

			NEUConstants.Pets = pets
		}
	}

	fmt.Printf("[NOT-ENOUGH-UPDATES] Parsing completed in %s\n", time.Since(timeNow))

	return nil
}
