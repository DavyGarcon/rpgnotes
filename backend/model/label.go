package model

const (
	BLACK int = 0x000000
	WHITE int = 0xffffff
)

type Label struct {
	Value string `json:"value,omitempty"`
	Color int    `json:"color,omitempty"`
}
