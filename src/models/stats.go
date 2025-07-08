package models

type ProfilesStats struct {
	ProfileId string `json:"profile_id"`
	CuteName  string `json:"profile_cute_name"`
	GameMode  string `json:"game_mode"`
	Selected  bool   `json:"selected"`
}

type MemberStats struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Removed bool   `json:"removed"`
}
