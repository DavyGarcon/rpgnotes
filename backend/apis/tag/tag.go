package tag

import (
	"context"
	"fmt"

	"encore.app/apis"
	"encore.app/common"
	"encore.app/model"
	"encore.app/service"
)

// Returns a list of all tags of kind of Tag .
//
//encore:api public method=GET path=/tags/:id
func Get(ctx context.Context, id string) (*model.Tag, error) {
	obj, err := service.LoadTag(id)
	common.CheckErr(err)
	fmt.Println("Tags: ", obj)
	return &obj, nil
}

// Returns a list of all tags of kind of Tag .
//
//encore:api public method=GET path=/tags
func List(ctx context.Context) (apis.ListTag, error) {
	list := apis.NewListTag()
	list.Content, _ = service.LoadTags()
	fmt.Println("Tags: ", list)
	return list, nil
}

// Persists the Tag passed in parameter.
//
//encore:api public method=POST path=/tags
func Create(ctx context.Context, obj model.Tag) error {
	fmt.Println("Tag: ", obj)

	err := service.PersistTag(obj)

	return err
}

// Updates the Tag passed in parameter.
//
//encore:api public method=PUT path=/tags/:id
func Update(ctx context.Context, id string, obj *model.Tag) error {
	obj.Id = id
	err := service.UpdateTag(*obj)
	common.CheckErr(err)
	return nil
}
