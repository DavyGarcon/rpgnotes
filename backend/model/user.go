package model

type User struct {
	Id    string `json:"id,omitempty"`
	Label Label  `json:"label,omitempty"`
	Image File   `json:"image,omitempty"`
}
