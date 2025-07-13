package models

type HypixelCollectionsResponse struct {
	Success     bool                         `json:"success"`
	LastUpdated int64                        `json:"lastUpdated"`
	Version     string                       `json:"version"`
	Collections map[string]HypixelCollection `json:"collections"`
}

type HypixelCollection struct {
	Name  string                           `json:"name"`
	Items map[string]HypixelCollectionItem `json:"items"`
}

type HypixelCollectionItem struct {
	Name     string                  `json:"name"`
	MaxTiers int                     `json:"maxTiers"`
	Tiers    []HypixelCollectionTier `json:"tiers"`
}

type HypixelCollectionTier struct {
	Tier           int `json:"tier"`
	AmountRequired int `json:"amountRequired"`
}

type ProcessedHypixelCollectionItem struct {
	Id      string                  `json:"id"`
	Name    string                  `json:"name"`
	Texture string                  `json:"texture"`
	MaxTier int                     `json:"maxTier"`
	Tiers   []HypixelCollectionTier `json:"tiers"`
}

type ProcessedHypixelCollection map[string]ProcessedHypixelCollectionCategory

type ProcessedHypixelCollectionCategory struct {
	Name        string `json:"name"`
	Texture     string `json:"texture"`
	Collections []ProcessedHypixelCollectionItem
}
