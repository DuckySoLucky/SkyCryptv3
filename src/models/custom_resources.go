package models

type ItemTexture struct {
	Parent    string            `json:"parent"`
	Textures  map[string]string `json:"textures"`
	Overrides []Override        `json:"overrides"`
}

type Override struct {
	Predicate map[string]interface{} `json:"predicate"`
	Texture   string                 `json:"model"`
}

type TextureItem struct {
	Count  *int                       `nbt:"Count" json:"Count,omitempty"`
	Damage *int                       `nbt:"Damage" json:"Damage,omitempty"`
	ID     *int                       `nbt:"id" json:"id,omitempty"`
	Tag    TextureItemExtraAttributes `nbt:"tag" json:"tag,omitempty"`
}

type TextureItemExtraAttributes struct {
	ExtraAttributes map[string]interface{} `nbt:"ExtraAttributes" json:"ExtraAttributes,omitempty"`
	Display         Display                `nbt:"display" json:"display"`
	SkullOwner      *SkullOwner            `nbt:"SkullOwner" json:"SkullOwner,omitempty"`
}

type VanillaTexture struct {
	VanillaId string `json:"vanillaId"`
	Damage    int    `json:"damage"`
}

type McMeta struct {
	Animation McMetaAnimation `json:"animation"`
}

type McMetaAnimation struct {
	Frametime int `json:"frametime"`
}
