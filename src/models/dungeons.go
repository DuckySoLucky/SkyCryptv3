package models

type DungeonsOutput struct {
	Level           Skill                   `json:"level"`
	Classes         ClassData               `json:"classes,omitempty"`
	Catacombs       []FormattedDungeonFloor `json:"catacombs,omitempty"`
	MasterCatacombs []FormattedDungeonFloor `json:"master_catacombs,omitempty"`
	Stats           DungeonStatsOutput      `json:"stats,omitempty"`
}

type ClassData struct {
	SelectedClass            string           `json:"selected_class,omitempty"`
	Classes                  map[string]Skill `json:"classes,omitempty"`
	ClassAverage             float64          `json:"class_average,omitempty"`
	ClassAverageWithProgress float64          `json:"class_average_with_progress,omitempty"`
	TotalExperience          float64          `json:"totalClassExp,omitempty"`
}

type FormattedDungeonFloor struct {
	Name    string            `json:"name"`
	Texture string            `json:"texture"`
	Stats   DungeonFloorStats `json:"stats"`
	BestRun *BestRunOutput    `json:"best_run,omitempty"`
}

type BestRunOutput struct {
	Grade            string  `json:"grade"`
	Timestamp        int64   `json:"timestamp"`
	ScoreExploration int     `json:"score_exploration"`
	ScoreSpeed       int     `json:"score_speed"`
	ScoreSkill       int     `json:"score_skill"`
	ScoreBonus       int     `json:"score_bonus"`
	DungeonClass     string  `json:"dungeon_class"`
	ElapsedTime      int64   `json:"elapsed_time"`
	DamageDealt      float64 `json:"damage_dealt"`
	Deaths           int     `json:"deaths"`
	MobsKilled       int     `json:"mobs_killed"`
	SecretsFound     int     `json:"secrets_found"`
	DamageMitigated  float64 `json:"damage_mitigated"`
}

type DungeonFloorStats struct {
	TimesPlayed          float64          `json:"times_played,omitempty"`
	TierCompletions      float64          `json:"tier_completions,omitempty"`
	MilestoneCompletions float64          `json:"milestone_completions,omitempty"`
	MobsKilled           float64          `json:"mobs_killed,omitempty"`
	BestScore            float64          `json:"best_score,omitempty"`
	WatcherKills         float64          `json:"watcher_kills,omitempty"`
	MostMobsKilled       float64          `json:"most_mobs_killed,omitempty"`
	FastestTime          float64          `json:"fastest_time,omitempty"`
	FastestTimeS         float64          `json:"fastest_time_s,omitempty"`
	FastestTimeSPlus     float64          `json:"fastest_time_s_plus,omitempty"`
	MostHealing          float64          `json:"most_healing,omitempty"`
	MostDamage           MostDamageOutput `json:"most_damage,omitempty"`
}

type MostDamageOutput struct {
	Damage float64 `json:"damage"`
	Type   string  `json:"type"`
}

type DungeonStatsOutput struct {
	Secrets                  SecretsOutput `json:"secrets"`
	HighestFloorBeatenNormal int           `json:"highest_floor_beaten_normal"`
	HighestFloorBeatenMaster int           `json:"highest_floor_beaten_master"`
	BloodMobKills            int           `json:"blood_mob_kills"`
}

type SecretsOutput struct {
	Found         int     `json:"found"`
	SecretsPerRun float64 `json:"secrets_per_run"`
}
