package model

type File struct {
	Path string `json:"path,omitempty"`
}

type Icon struct {
	Path  string `json:"path,omitempty"`
	Color int    `json:"color,omitempty"`
}
