package stats

import (
	"fmt"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"skycrypt/src/utility"
)

func ProcessItems(items *[]models.Item, source string) []models.ProcessedItem {
	var processedItems []models.ProcessedItem
	for _, item := range *items {
		processedItem := ProcessItem(&item, source)
		processedItems = append(processedItems, processedItem)
	}

	return processedItems
}

func ProcessItem(item *models.Item, source string) models.ProcessedItem {
	if item.Tag == nil {
		return models.ProcessedItem{}
	}

	processedItem := models.ProcessedItem{
		Item:        *item,
		DisplayName: item.Tag.Display.Name,
		Lore:        item.Tag.Display.Lore,
		Source:      source,
	}

	// POTIONS
	if *item.ID == 373 {
		color := constants.POTION_COLORS[*item.Damage]
		var potionType string
		if *item.Damage&16384 != 0 {
			potionType = "splash"
		} else {
			potionType = "normal"
		}

		// TODO: Once pack support is implemented, this should be removed
		processedItem.Texture = "sky.shiiyu.moe/api/potion/" + potionType + "/" + color
		// processedItem.Texture = "/api/potion/" + potionType + "/" + color
	}

	rawLore := make([]string, len(processedItem.Lore))
	for i, lore := range processedItem.Lore {
		rawLore[i] = utility.GetRawLore(lore)
	}

	itemType := ParseItemTypeFromLore(rawLore, *item)
	processedItem.Rarity = itemType.Rarity
	processedItem.Categories = itemType.Categories
	processedItem.Recombobulated = item.Tag.ExtraAttributes.Recombobulated == 1

	if processedItem.Recombobulated {
		processedItem.Lore = append(processedItem.Lore, "§8(Recombobulated)")
	}

	if item.Tag.ExtraAttributes.Timestamp != nil {
		if timestamp, ok := item.Tag.ExtraAttributes.Timestamp.(float64); ok {
			processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Obtained: §c{TIMESTAMP:%.0f}", timestamp))
		} else if timestamp, ok := item.Tag.ExtraAttributes.Timestamp.(string); ok {
			parsedTimestamp := utility.ParseTimestamp(timestamp)
			processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Obtained: §c{TIMESTAMP:%d}", parsedTimestamp))
		} else if timestamp, ok := item.Tag.ExtraAttributes.Timestamp.(int64); ok {
			processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Obtained: §c{TIMESTAMP:%d}", timestamp))
		} else {
			fmt.Printf("Unexpected type for timestamp: %T, %s\n", item.Tag.ExtraAttributes.Timestamp, item.Tag.ExtraAttributes.Timestamp)
		}
	}

	if item.Tag.Display.Color != 0 {
		hex := fmt.Sprintf("#%06X", item.Tag.Display.Color)

		processedItem.Lore = append(processedItem.Lore, "", fmt.Sprintf("§7Color: %s", hex))
	}

	if item.Tag.ExtraAttributes.Gems != nil {
		gems := ParseItemGems(item.Tag.ExtraAttributes.Gems, itemType.Rarity)
		if len(gems) > 0 {
			processedItem.Lore = append(processedItem.Lore, "", "§7Applied Gemstones:")
			for _, gem := range gems {
				processedItem.Lore = append(processedItem.Lore, fmt.Sprintf("§7 - %s", gem.Lore))
			}
		}
	}

	if item.Tag.ExtraAttributes.HecatombSRuns != nil {
		AddLevelableEnchantmentsToLore(*item.Tag.ExtraAttributes.HecatombSRuns, constants.ENCHANTMENT_LADDERS["hecatomb_s_runs"], &processedItem.Lore)
	}

	if item.Tag.ExtraAttributes.ChampionCombatXP != nil {
		AddLevelableEnchantmentsToLore(int(*item.Tag.ExtraAttributes.ChampionCombatXP), constants.ENCHANTMENT_LADDERS["champion_combat_xp"], &processedItem.Lore)
	}

	if item.Tag.ExtraAttributes.FarmedCultivating != nil {
		AddLevelableEnchantmentsToLore(*item.Tag.ExtraAttributes.FarmedCultivating, constants.ENCHANTMENT_LADDERS["farmed_cultivating"], &processedItem.Lore)
	}

	if item.Tag.ExtraAttributes.ExpertiseKills != nil {
		AddLevelableEnchantmentsToLore(*item.Tag.ExtraAttributes.ExpertiseKills, constants.ENCHANTMENT_LADDERS["expertise_kills"], &processedItem.Lore)
	}

	if item.Tag.ExtraAttributes.CompactBlocks != nil {
		AddLevelableEnchantmentsToLore(*item.Tag.ExtraAttributes.CompactBlocks, constants.ENCHANTMENT_LADDERS["compact_blocks"], &processedItem.Lore)
	}

	if processedItem.Texture == "" {
		// TODO: Once pack support is implemented, this should be removed
		processedItem.Texture = "http://sky.shiiyu.moe/api/item/" + item.Tag.ExtraAttributes.ID
	}

	if item.ContainsItems != nil {
		processedItem.ContainsItems = ProcessItems(&item.ContainsItems, source)
	}

	return processedItem
}
