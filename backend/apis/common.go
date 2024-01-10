package apis

import "encore.app/model"

type ListId struct {
	Content []string
}

type ListCampaign struct {
	Content []model.Campaign
}

func NewListCampaign() ListCampaign {
	return ListCampaign{Content: make([]model.Campaign, 0)}

}

func (lu *ListCampaign) add(campaign model.Campaign) {
	lu.Content = append(lu.Content, campaign)
}

type CategoryParams struct {
	CampaignId string `query:"campaignId,omitempty"`
	CategoryId string `query:"categoryId,omitempty"`
}
type ContainerParams struct {
	CampaignId string `query:"campaignId,omitempty"`
	CategoryId string `query:"categoryId,omitempty"`
}

type ListCategory struct {
	Content []model.Category
}

func NewListCategory() ListCategory {
	return ListCategory{Content: make([]model.Category, 0)}

}

func (lu *ListCategory) add(category model.Category) {
	lu.Content = append(lu.Content, category)
}

type CampaignParams struct {
	CampaignId string `query:"campaignId,omitempty"`
	CategoryId string `query:"categoryId,omitempty"`
}

type ListContainer struct {
	Content []model.Container
}

func NewListContainer() ListContainer {
	return ListContainer{Content: make([]model.Container, 0)}

}

func (lu *ListContainer) add(container model.Container) {
	lu.Content = append(lu.Content, container)
}

type FileParams struct {
	Decoded bool `query:"decoded,omitempty"`
}

type FileContent struct {
	DecodedContent *string `json:"decoded,omitempty"`
	EncodedContent *[]byte `json:"encoded,omitempty"`
}

type ListFile struct {
	Content []model.File `json:"content,omitempty"`
}

type ElementParams struct {
	ContainerId string `query:"containerId,omitempty"`
}

type ListElement struct {
	Content []model.Element
}

func NewListElement() ListElement {
	return ListElement{Content: make([]model.Element, 0)}

}

func (lu *ListElement) add(element model.Element) {
	lu.Content = append(lu.Content, element)
}

type ListTag struct {
	Content []model.Tag
}

func NewListTag() ListTag {
	return ListTag{Content: make([]model.Tag, 0)}

}

func (lu *ListTag) add(tag model.Tag) {
	lu.Content = append(lu.Content, tag)
}

type ListUser struct {
	Content []model.User
}

func NewListUser() ListUser {
	return ListUser{Content: make([]model.User, 0)}

}

func (lu *ListUser) add(user model.User) {
	lu.Content = append(lu.Content, user)
}
