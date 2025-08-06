package models

type CrimsonIsleFactions struct {
	SelectedFaction     string `json:"selectedFaction"`
	BarbarianReputation int    `json:"barbariansReputation"`
	MagesReputation     int    `json:"magesReputation"`
}

type CrimsonIsleKuudra struct {
	TotalKills int                     `json:"totalKills"`
	Tiers      []CrimsonIsleKuudraTier `json:"tiers"`
}

type CrimsonIsleKuudraTier struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Texture string `json:"texture"`
	Kills   int    `json:"kills"`
}

type CrimsonIsleDojo struct {
	TotalPoints int                        `json:"totalPoints"`
	Challenges  []CrimsonIsleDojoChallenge `json:"challenges"`
}

type CrimsonIsleDojoChallenge struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Texture string `json:"texture"`
	Points  int    `json:"points"`
	Time    int    `json:"time"`
	Rank    string `json:"rank"`
}

type CrimsonIsleOutput struct {
	Factions CrimsonIsleFactions `json:"factions"`
	Kuudra   CrimsonIsleKuudra   `json:"kuudra"`
	Dojo     CrimsonIsleDojo     `json:"dojo"`
}
