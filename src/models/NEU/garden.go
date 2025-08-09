package neu

type NEUGardenRaw struct {
	GardenExperience  []int                    `json:"garden_exp"`
	CropMilestones    map[string][]int         `json:"crop_milestones"`
	Visitors          map[string]string        `json:"visitors"`
	Plots             map[string]NEUPlotLayout `json:"plots"`
	CropUpgrades      []int                    `json:"crop_upgrades"`
	ComposterTooltips map[string]string        `json:"composter_tooltips"`
	BarnSkins         map[string]*BarnSkin     `json:"barn"`
}

type NEUPlotLayout struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type NEUGarden struct {
	GardenExperience  map[int]int            `json:"garden_exp"`
	CropMilestones    map[string]map[int]int `json:"crop_milestones"`
	Visitors          map[string]string      `json:"visitors"`
	Plots             map[string]string      `json:"plots"`
	SortedPlots       []string               `json:"sorted_plots"`
	MaxVisitors       map[string]int         `json:"max_visitors"`
	CropUpgrades      map[int]int            `json:"crop_upgrades"`
	ComposterUpgrades []string               `json:"composter_upgrades"`
	BarnSkins         map[string]*BarnSkin   `json:"barn"`
}

type BarnSkin struct {
	Name   string `json:"name"`
	ItemId string `json:"item"`
}
