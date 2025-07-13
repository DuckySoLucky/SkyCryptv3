package models

type CollectionsOutput struct {
	Categories       map[string]CollectionCategory `json:"categories"`
	TotalCollections int                           `json:"totalCollections"`
	MaxedCollections int                           `json:"maxedCollections"`
}

type CollectionCategory struct {
	Name       string                   `json:"name"`
	Texture    string                   `json:"texture"`
	Items      []CollectionCategoryItem `json:"items"`
	TotalTiers int                      `json:"totalTiers"`
	MaxedTiers int                      `json:"maxedTiers"`
}

type CollectionCategoryItem struct {
	Name        string                         `json:"name"`
	Id          string                         `json:"id"`
	Texture     string                         `json:"texture"`
	Amount      int                            `json:"amount"`
	TotalAmount int                            `json:"totalAmount"`
	Tier        int                            `json:"tier"`
	MaxTier     int                            `json:"maxTier"`
	Amounts     []CollectionCategoryItemAmount `json:"amounts"`
}

type CollectionCategoryItemAmount struct {
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}

type BossCollectionsFloorData struct {
	FloorId int
	Item    CollectionCategoryItem
}
