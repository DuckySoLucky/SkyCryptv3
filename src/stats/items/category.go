package stats

import (
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"slices"
	"sort"
	"strings"
)

type itemSorter []models.ProcessedItem

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

func GetCategory(allItems []models.ProcessedItem, category string) []models.ProcessedItem {
	var output []models.ProcessedItem

	for _, item := range allItems {
		if item.Categories != nil {
			if slices.Contains(item.Categories, category) {
				output = append(output, item)
			}
		}
	}

	for _, item := range allItems {
		if item.ContainsItems != nil {
			containedItems := GetCategory(item.ContainsItems, category)
			output = append(output, containedItems...)
		}
	}

	sort.Sort(itemSorter(output))
	return output
}

type weaponsResult struct {
	Weapons               []models.ProcessedItem `json:"weapons"`
	HighestPriorityWeapon *models.ProcessedItem  `json:"highest_priority_weapon"`
}

func GetWeapons(allItems []models.ProcessedItem) weaponsResult {
	weapons := GetCategory(allItems, "weapon")

	countsOfID := make(map[string]int)
	for _, weapon := range weapons {
		id := GetId(weapon)
		countsOfID[id]++
	}

	var filteredWeapons []models.ProcessedItem
	itemCounts := make(map[string]int)

	for _, weapon := range weapons {
		id := GetId(weapon)
		itemCounts[id]++

		if itemCounts[id] <= 2 {
			filteredWeapons = append(filteredWeapons, weapon)
		}
	}

	weapons = filteredWeapons

	swords := GetCategory(allItems, "sword")
	var highestPriorityWeapon *models.ProcessedItem
	if len(swords) > 0 {
		highestPriorityWeapon = &swords[0]
	}

	return weaponsResult{
		Weapons:               weapons,
		HighestPriorityWeapon: highestPriorityWeapon,
	}
}

type skillToolsResult struct {
	Tools               []models.ProcessedItem `json:"tools"`
	HighestPriorityTool *models.ProcessedItem  `json:"highest_priority_tool"`
}

func GetSkillTools(skill string, allItems []models.ProcessedItem) skillToolsResult {
	toolCategory := skill + "_tool"
	tools := GetCategory(allItems, toolCategory)

	var highestPriorityTool *models.ProcessedItem
	if len(tools) > 0 {
		highestPriorityTool = &tools[0]
	}

	return skillToolsResult{
		Tools:               tools,
		HighestPriorityTool: highestPriorityTool,
	}
}
