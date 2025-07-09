package stats

import (
	"regexp"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"sort"
	"strconv"
	"strings"
)

type ItemStats map[string]float64

var regex = regexp.MustCompile(`^([A-Za-z ]+): ([+-]([0-9]+(?:,[0-9]{3})*(?:\.[0-9]{0,2})?))`)

func GetStatsFromItem(item models.ProcessedItem) ItemStats {
	stats := make(ItemStats)

	if item.Tag.Display.Lore == nil {
		return stats
	}

	lore := make([]string, len(item.Tag.Display.Lore))
	for i, line := range item.Tag.Display.Lore {
		lore[i] = utility.GetRawLore(line)
	}

	for _, line := range lore {
		matches := regex.FindStringSubmatch(line)

		if matches == nil {
			continue
		}

		var statName string
		for key, statInfo := range constants.STATS_DATA {
			if statInfo.NameLore == matches[1] {
				statName = key
				break
			}
		}

		if statName != "" {
			statValueStr := strings.ReplaceAll(matches[2], ",", "")
			statValue, err := strconv.ParseFloat(statValueStr, 64)
			if err != nil {
				continue
			}

			if _, exists := stats[statName]; !exists {
				stats[statName] = 0
			}
			stats[statName] += statValue
		}
	}

	return stats
}

func GetStatsFromItems(items []models.ProcessedItem) ItemStats {
	stats := make(ItemStats)

	for _, item := range items {
		itemStats := GetStatsFromItem(item)
		for stat, value := range itemStats {
			if _, exists := stats[stat]; !exists {
				stats[stat] = 0
			}
			stats[stat] += value
		}
	}

	type statEntry struct {
		key   string
		value float64
	}

	entries := make([]statEntry, 0, len(stats))
	for k, v := range stats {
		entries = append(entries, statEntry{k, v})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].value > entries[j].value
	})

	sortedStats := make(ItemStats)
	for _, entry := range entries {
		sortedStats[entry.key] = entry.value
	}

	return sortedStats
}
