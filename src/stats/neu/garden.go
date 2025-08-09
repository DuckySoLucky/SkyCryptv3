package neustats

import (
	neu "skycrypt/src/models/NEU"
	"sort"
)

func formatXPTable(xpTable []int) map[int]int {
	output := make(map[int]int)
	for i, xp := range xpTable {
		if xp == 0 {
			xp = 1
		}

		output[i+1] = xp
	}

	return output
}

func getMaxVisitors(gardenConstants neu.NEUGardenRaw) map[string]int {
	output := map[string]int{}
	for _, visitorRarity := range gardenConstants.Visitors {
		output[visitorRarity]++
	}

	return output
}

func getCropMilestones(cropMilestones map[string][]int) map[string]map[int]int {
	output := make(map[string]map[int]int)
	for crop, milestones := range cropMilestones {
		output[crop] = formatXPTable(milestones)
	}

	return output
}

func getComposterUpgrades(composterTooltips map[string]string) []string {
	output := make([]string, 0, len(composterTooltips))
	for upgrade := range composterTooltips {
		output = append(output, upgrade)
	}

	return output
}

func getSortedPlots(plots map[string]neu.NEUPlotLayout) []string {
	plotIds := make([]string, 0, len(plots))
	for plotId := range plots {
		plotIds = append(plotIds, plotId)
	}

	sort.Slice(plotIds, func(i, j int) bool {
		a := plots[plotIds[i]]
		b := plots[plotIds[j]]

		if a.Y == b.Y {
			return a.X < b.X
		}
		return a.Y < b.Y
	})

	return plotIds
}

func getPlotLayout(plots map[string]neu.NEUPlotLayout) map[string]string {
	output := make(map[string]string)
	for plotId, plot := range plots {
		output[plotId] = plot.Name
	}

	return output
}

func FormatGardenConstants(gardenConstants neu.NEUGardenRaw) neu.NEUGarden {
	return neu.NEUGarden{
		GardenExperience:  formatXPTable(gardenConstants.GardenExperience),
		CropMilestones:    getCropMilestones(gardenConstants.CropMilestones),
		Visitors:          gardenConstants.Visitors,
		Plots:             getPlotLayout(gardenConstants.Plots),
		SortedPlots:       getSortedPlots(gardenConstants.Plots),
		MaxVisitors:       getMaxVisitors(gardenConstants),
		CropUpgrades:      formatXPTable(gardenConstants.CropUpgrades),
		ComposterUpgrades: getComposterUpgrades(gardenConstants.ComposterTooltips),
		BarnSkins:         gardenConstants.BarnSkins,
	}
}
