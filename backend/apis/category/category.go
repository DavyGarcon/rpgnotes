package category

import (
	"context"
	"fmt"

	"encore.app/apis"
	"encore.app/common"
	"encore.app/model"
	"encore.app/service"
)

// Persists the Category passed in parameter.
//
//encore:api public method=POST path=/categories/:id/categories
func AddCategory(ctx context.Context, id string, obj model.Category) error {
	fmt.Println("Category: ", obj)
	obj.CategoryId = id
	obj.CampaignId = ""
	err := service.PersistCategory(obj)

	return err
}

// Persists the Container passed in parameter.
//
//encore:api public method=POST path=/categories/:id/containers
func AddContainer(ctx context.Context, id string, obj model.Container) error {
	fmt.Println("Container: ", obj)
	obj.CategoryId = id
	obj.CampaignId = ""
	err := service.PersistContainer(obj)

	return err
}

// Returns a list of all categorys of kind of Category .
//
//encore:api public method=GET path=/categorys/:id
func Get(ctx context.Context, id string) (*model.Category, error) {
	obj, err := service.LoadCategory(id)
	common.CheckErr(err)
	fmt.Println("Categorys: ", obj)
	return &obj, nil
}

// Updates the Category passed in parameter.
//
//encore:api public method=PUT path=/categorys/:id
func Update(ctx context.Context, id string, obj *model.Category) error {
	obj.Id = id
	err := service.UpdateCategory(*obj)
	common.CheckErr(err)
	return nil
}

// Returns a list of all containers of kind of Container .
//
//encore:api public method=GET path=/categorys/:id/containers
func ListContainer(ctx context.Context, id string) (apis.ListContainer, error) {
	list := apis.NewListContainer()
	list.Content, _ = service.FindContainersByCategoryId(id)
	fmt.Println("Containers: ", list)
	return list, nil
}

// Returns a list of all containers of kind of Container .
//
//encore:api public method=GET path=/categories/:id/categories
func ListCategories(ctx context.Context, id string) (apis.ListCategory, error) {
	list := apis.NewListCategory()
	list.Content, _ = service.FindCategoriesByCategoryId(id)
	fmt.Println("Categories: ", list)
	return list, nil
}
