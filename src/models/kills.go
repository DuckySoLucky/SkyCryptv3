package models

type Kill struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Texture string `json:"texture"`
	Amount  int    `json:"amount"`
}
