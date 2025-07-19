package tools

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	assetsRoot := "assets"
	packDirs, err := os.ReadDir(assetsRoot)
	if err != nil {
		fmt.Printf("Failed to read assets directory: %v\n", err)
		return
	}

	for _, packDir := range packDirs {
		if !packDir.IsDir() {
			continue
		}

		if packDir.Name() == "Vanilla" {
			continue
		}

		packAssetsPath := filepath.Join(assetsRoot, packDir.Name(), "assets")
		if _, err := os.Stat(packAssetsPath); os.IsNotExist(err) {
			continue
		}

		filepath.WalkDir(packAssetsPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			if !strings.Contains(path, "/models/item/") {
				return nil
			}

			if strings.HasSuffix(path, ".json") {
				return nil
			}

			/*
				var model models.VanillaTexture
				if err := json.Unmarshal(data, &model); err != nil {
					fmt.Printf("Failed to parse %s: %v\n", path, err)
					return nil
				}

				textureId := fmt.Sprintf("%s:%d", model.VanillaId, model.Damage)
				fileName := strings.ReplaceAll(filepath.Base(path), ".json", ".png")
				VANILLA_ITEM_MAP[textureId] = models.ItemTexture{
					Parent:    "item/generated",
					Textures:  map[string]string{"layer0": GetTexturePath(packDir.Name(), fileName)},
					Overrides: []models.Override{},
				}
			*/
			return nil

		})
	}
}
