package neu

import (
	"encoding/json"
)

type NEUBestiaryRawIslandData struct {
	Name               string                              `json:"name"`
	Icon               NEUBestiaryRawIslandIcon            `json:"icon"`
	Mobs               []NEUBestiaryRawMob                 `json:"mobs"`
	HasSubcategories   bool                                `json:"hasSubcategories"`
	SubcategoryIslands map[string]NEUBestiaryRawIslandData `json:"-"`
}

func (i *NEUBestiaryRawIslandData) UnmarshalJSON(data []byte) error {
	type Alias NEUBestiaryRawIslandData
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// If HasSubcategories, find all subcategories
	if i.HasSubcategories {
		var raw map[string]json.RawMessage
		if err := json.Unmarshal(data, &raw); err != nil {
			return err
		}
		i.SubcategoryIslands = make(map[string]NEUBestiaryRawIslandData)
		for k, v := range raw {
			// skip known fields
			if k == "name" || k == "icon" || k == "mobs" || k == "hasSubcategories" {
				continue
			}
			var sub NEUBestiaryRawIslandData
			if err := json.Unmarshal(v, &sub); err != nil {
				return err
			}
			i.SubcategoryIslands[k] = sub
		}
	}
	return nil
}

type NEUBestiaryRawIslandIcon struct {
	SkullOwner string `json:"skullOwner"`
	Texture    string `json:"texture"`
	Item       string `json:"item"`
}

type NEUBestiaryRawMob struct {
	Name       string   `json:"name"`
	Item       string   `json:"item,omitempty"`
	SkullOwner string   `json:"skullOwner,omitempty"`
	Texture    string   `json:"texture,omitempty"`
	Cap        int      `json:"cap"`
	Mobs       []string `json:"mobs"`
	Bracket    int      `json:"bracket"`
}

type NEUBestiaryRaw struct {
	Brackets map[string][]int
	Islands  map[string]NEUBestiaryRawIslandData
}

func (n *NEUBestiaryRaw) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	n.Islands = make(map[string]NEUBestiaryRawIslandData)

	for k, v := range raw {
		if k == "brackets" {
			if err := json.Unmarshal(v, &n.Brackets); err != nil {
				return err
			}
		} else {
			var island NEUBestiaryRawIslandData
			if err := json.Unmarshal(v, &island); err != nil {
				return err
			}
			n.Islands[k] = island
		}
	}
	return nil
}

type BestiaryConstants struct {
	Brackets map[string][]int            `json:"brackets"`
	Islands  map[string]BestiaryCategory `json:"islands"`
}

type BestiaryCategory struct {
	Name    string        `json:"name"`
	Texture string        `json:"texture"`
	Mobs    []BestiaryMob `json:"mobs"`
}

type BestiaryMob struct {
	Name    string   `json:"name"`
	Texture string   `json:"texture"`
	Cap     int      `json:"cap"`
	Mobs    []string `json:"mobs"`
	Bracket int      `json:"bracket"`
}
