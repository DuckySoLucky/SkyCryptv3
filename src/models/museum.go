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
	Value     int64                 `json:"value"`
	Appraisal bool                  `json:"appraisal,omitempty"`
	Items     map[string]museumItem `json:"items,omitempty"`
	Special   []specialItems        `json:"special,omitempty"`
}

type museumItem struct {
	DonatedTime  int64       `json:"donated_time"`
	FeaturedSlot *string     `json:"featured_slot"`
	Borrowing    bool        `json:"borrowing"`
	Items        EncodedItem `json:"items"`
}

type specialItems struct {
	DonatedTime int64       `json:"donated_time"`
	Items       EncodedItem `json:"items"`
}
