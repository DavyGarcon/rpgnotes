package model

type Tag struct {
	Id    string `json:"id,omitempty"`
	Label Label  `jsony:"label,omitempty"`
	Color int    `json:"color,omitempty"`
}
