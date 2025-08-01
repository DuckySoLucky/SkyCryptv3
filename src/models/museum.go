package models

type EncodedItem struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

type HypixelMuseumResponse struct {
	Success bool               `json:"success"`
	Cause   string             `json:"cause,omitempty"`
	Members map[string]*Museum `json:"members"`
}

type Museum struct {
	Value     int64                     `json:"value"`
	Appraisal bool                      `json:"appraisal,omitempty"`
	Items     map[string]*RawMuseumItem `json:"items,omitempty"`
	Special   []RawMuseumItem           `json:"special,omitempty"`
}

type RawMuseumItem struct {
	DonatedTime  int64       `json:"donated_time"`
	FeaturedSlot *string     `json:"featured_slot"`
	Borrowing    bool        `json:"borrowing"`
	Items        EncodedItem `json:"items"`
}

type MuseumInventoryItem struct {
	ProcessedItem
	Position      int                   `json:"position"`
	ProgressType  string                `json:"progress_type"`
	InventoryType string                `json:"inventory_type,omitempty"`
	ContainsItems []MuseumInventoryItem `json:"containsItems,omitempty"`
}

type MuseumItem struct {
	ProcessedItem
	Borrowing    bool  `json:"borrowing"`
	DonationTime int64 `json:"donated_time"`
}

type ProcessedMuseumItem struct {
	Items           []ProcessedItem `json:"items"`
	Missing         bool            `json:"missing"`
	DonatedAsAChild bool            `json:"donated_as_a_child"`
	SkyblockID      string          `json:"id"`
}

type DecodedMuseumItems struct {
	Items   map[string]ProcessedMuseumItem `json:"items"`
	Special []ProcessedMuseumItem          `json:"special"`
	Value   int64                          `json:"value"`
}

type MuseumStats struct {
	Amount int `json:"amount"`
	Total  int `json:"total"`
}

type MuseumSpecialStats struct {
	Amount int `json:"amount"`
}

type MuseumMissing struct {
	Main []string `json:"main"`
	Max  []string `json:"max"`
}

type MuseumResult struct {
	Value        int64                          `json:"value"`
	Appraisal    bool                           `json:"appraisal"`
	Total        MuseumStats                    `json:"total"`
	Weapons      MuseumStats                    `json:"weapons"`
	Armor        MuseumStats                    `json:"armor"`
	Rarities     MuseumStats                    `json:"rarities"`
	Special      MuseumSpecialStats             `json:"special"`
	Items        map[string]ProcessedMuseumItem `json:"items"`
	SpecialItems []ProcessedMuseumItem          `json:"specialItems"`
	Missing      MuseumMissing                  `json:"missing"`
}
