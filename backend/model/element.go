package model

type Element struct {
	Id          string   `json:"id,omitempty"`
	Label       Label    `json:"label,omitempty"`
	Comment     Label    `json:"comment,omitempty"`
	Description string   `json:"description,omitempty"`
	Icon        Icon     `json:"icon,omitempty"`
	Notes       []Note   `json:"notes,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Connections []string `json:"connections,omitempty"`
	Images      []File   `json:"images,omitempty"`
	ContainerId string   `json:"containerId,omitempty"`
	Order       int      `json:"index,omitempty"`
}
