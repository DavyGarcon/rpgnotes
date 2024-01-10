package model

type Note struct {
	Text  string `json:"text,omitempty"`
	Order int    `json:"index,omitempty"`
}
