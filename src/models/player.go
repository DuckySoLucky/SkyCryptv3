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
		Links SocialMediaLinks `json:"links"`
	} `json:"socialMedia"`
	NewPackageRank            string       `json:"newPackageRank,omitempty"`
	MonthlyRankColor          string       `json:"monthlyRankColor,omitempty"`
	MonthlyPackageRank        string       `json:"monthlyPackageRank,omitempty"`
	Prefix                    string       `json:"prefix"`
	Rank                      string       `json:"rank"`
	RankPlusColor             string       `json:"rankPlusColor,omitempty"`
	PackageRank               string       `json:"packageRank,omitempty"`
	Achievements              achievements `json:"achievements,omitempty"`
	ClaimedPotatoTalisman     int64        `json:"claimed_potato_talisman,omitempty"`
	ClaimedPotatoBasket       int64        `json:"claimed_potato_basket,omitempty"`
	ClaimPotatoWarSilverMedal int64        `json:"claim_potato_war_silver_medal,omitempty"`
	ClaimPotatoWarCrown       int64        `json:"claim_potato_war_crown,omitempty"`
	SkyblockFreeCookie        int64        `json:"skyblock_free_cookie,omitempty"`
	ClaimedCenturyCake        int64        `json:"claimed_century_cake,omitempty"`
	ClaimedCenturyCake200     int64        `json:"claimed_century_cake200,omitempty"`
}

type SocialMediaLinks struct {
	Twitter string `json:"TWITTER,omitempty"`
	Twitch  string `json:"TWITCH,omitempty"`
	Hypixel string `json:"HYPIXEL,omitempty"`
	Discord string `json:"DISCORD,omitempty"`
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
