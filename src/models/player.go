package models

type HypixelPlayerResponse struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause,omitempty"`
	Player  Player `json:"player"`
}

type Player struct {
	DisplayName string `json:"displayname"`
	UUID        string `json:"uuid"`
	SocialMedia struct {
		Links struct {
			Twitter string `json:"TWITTER,omitempty"`
			Twitch  string `json:"TWITCH,omitempty"`
			Hypixel string `json:"HYPIXEL,omitempty"`
			Discord string `json:"DISCORD,omitempty"`
		} `json:"links"`
	} `json:"socialMedia"`
	NewPackageRank     string       `json:"newPackageRank,omitempty"`
	MonthlyRankColor   string       `json:"monthlyRankColor,omitempty"`
	MonthlyPackageRank string       `json:"monthlyPackageRank,omitempty"`
	Prefix             string       `json:"prefix"`
	Rank               string       `json:"rank"`
	RankPlusColor      string       `json:"rankPlusColor,omitempty"`
	PackageRank        string       `json:"packageRank,omitempty"`
	Achievements       achievements `json:"achievements,omitempty"`
}

type achievements struct {
	SkillTaming     int `json:"skyblock_domesticator,omitempty"`
	SkillFarming    int `json:"skyblock_harvester,omitempty"`
	SkillMining     int `json:"skyblock_excavator,omitempty"`
	SkillCombat     int `json:"skyblock_combat,omitempty"`
	SkillForaging   int `json:"skyblock_gatherer,omitempty"`
	SkillFishing    int `json:"skyblock_angler,omitempty"`
	SkillEnchanting int `json:"skyblock_augmentation,omitempty"`
	SkillAlchemy    int `json:"skyblock_concoctor,omitempty"`
	HotMCommissions int `json:"skyblock_hard_working_miner,omitempty"`
}
