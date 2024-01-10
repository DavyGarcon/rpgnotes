package campaign

import (
	"context"
	"fmt"

	"encore.app/apis"
	"encore.app/common"
	"encore.app/model"
	"encore.app/service"
)

// Returns a list of all campaigns of kind of Campaign .
//
//encore:api public method=GET path=/campaigns/:id
func Get(ctx context.Context, id string) (*model.Campaign, error) {
	obj, err := service.LoadCampaign(id)
	common.CheckErr(err)
	fmt.Println("Campaigns: ", obj)
	return &obj, nil
}

// Returns a list of all campaigns of kind of Campaign .
//
//encore:api public method=GET path=/campaigns
func List(ctx context.Context) (apis.ListCampaign, error) {
	list := apis.NewListCampaign()
	list.Content, _ = service.LoadCampaigns()
	fmt.Println("Campaigns: ", list)
	return list, nil
}

// Persists the Campaign passed in parameter.
//
//encore:api public method=POST path=/campaigns
func Create(ctx context.Context, obj model.Campaign) error {
	fmt.Println("Campaign: ", obj)

	campaignId, err := service.PersistCampaign(obj)
	fmt.Println("CampaignId: ", campaignId)

	return err
}

// Updates the Campaign passed in parameter.
//
//encore:api public method=PUT path=/campaigns/:id
func Update(ctx context.Context, id string, obj *model.Campaign) error {
	obj.Id = id
	err := service.UpdateCampaign(*obj)
	common.CheckErr(err)
	return nil
}

// Updates the Campaign passed in parameter.
//
//encore:api public method=PUT path=/campaigns/:id/load_icon
func GetIconContent(ctx context.Context, id string, queryParams *apis.FileParams) (apis.FileContent, error) {
	campaign, err := service.LoadCampaign(id)
	common.CheckErr(err)
	if queryParams != nil && queryParams.Decoded {
		content, err := common.GetDecodedFileContent(campaign.Icon.Path)
		common.CheckErr(err)
		return apis.FileContent{EncodedContent: nil, DecodedContent: &content}, nil
	}
	content, err := common.GetEncodedFileContent(campaign.Icon.Path)
	common.CheckErr(err)
	return apis.FileContent{DecodedContent: nil, EncodedContent: &content}, nil
}

// Persists the Category passed in parameter.
//
//encore:api public method=POST path=/campaigns/:id/categories
func AddCategory(ctx context.Context, id string, obj model.Category) error {
	fmt.Println("Category: ", obj)
	obj.CampaignId = id
	obj.CategoryId = ""
	err := service.PersistCategory(obj)

	return err
}

// Persists the Container passed in parameter.
//
//encore:api public method=POST path=/campaigns/:id/containers
func AddContainer(ctx context.Context, id string, obj model.Container) error {
	fmt.Println("Container: ", obj)
	obj.CategoryId = ""
	obj.CampaignId = id
	err := service.PersistContainer(obj)

	return err
}

// Persists the User passed in parameter.
//
//encore:api public method=POST path=/campaigns/:id/users
func AddPlayer(ctx context.Context, id string, ids apis.ListId) error {
	fmt.Println("User: ", ids)
	campaign, err := service.LoadCampaign(id)
	common.CheckErr(err)
	players := append(*campaign.Players, ids.Content...)
	campaign.Players = &players
	fmt.Println("Players: ", campaign.Players)
	campaignId, err := service.PersistCampaign(campaign)
	common.CheckErr(err)
	fmt.Println("Campaign updated: ", campaignId)
	return err
}

// Returns a list of all containers of kind of Container .
//
//encore:api public method=GET path=/campaigns/:id/containers
func ListContainer(ctx context.Context, id string) (apis.ListContainer, error) {
	list := apis.NewListContainer()
	list.Content, _ = service.FindContainersByCampaignId(id)
	fmt.Println("Containers: ", list)
	return list, nil
}

// Returns a list of all containers of kind of Container .
//
//encore:api public method=GET path=/campaigns/:id/categories
func ListCategories(ctx context.Context, id string) (apis.ListCategory, error) {
	list := apis.NewListCategory()
	list.Content, _ = service.FindCategoriesByCampaignId(id)
	fmt.Println("Categories: ", list)
	return list, nil
}
