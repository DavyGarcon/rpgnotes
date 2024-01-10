package model

type Container struct {
	Id         string `json:"id,omitempty"`
	Label      Label  `json:"label,omitempty"`
	Comment    Label  `json:"comment,omitempty"`
	Icon       Icon   `json:"icon,omitempty"`
	CampaignId string `json:"campaignId,omitempty"`
	CategoryId string `json:"categoryId,omitempty"`
	Order      int    `json:"index,omitempty"`
}

type Campaign struct {
	Id         string    `json:"id,omitempty"`
	Label      Label     `json:"label,omitempty"`
	Comment    Label     `json:"comment,omitempty"`
	Icon       Icon      `json:"icon,omitempty"`
	Active     bool      `json:"active,omitempty"`
	Notes      *[]Note   `json:"notes,omitempty"`
	GameMaster string    `json:"gameMasterId,omitempty"`
	Players    *[]string `json:"playerIds,omitempty"`
	Order      int       `json:"index,omitempty"`
}

type Category struct {
	Id         string `json:"id,omitempty"`
	Label      Label  `json:"label,omitempty"`
	Comment    Label  `json:"comment,omitempty"`
	Icon       Icon   `json:"icon,omitempty"`
	CampaignId string `json:"campaignId,omitempty"`
	CategoryId string `json:"parentId,omitempty"`
	Order      int    `json:"index,omitempty"`
}
