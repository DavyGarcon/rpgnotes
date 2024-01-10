package element

import (
	"context"
	"fmt"

	"encore.app/common"
	"encore.app/model"
	"encore.app/service"
)

// Persists the Element passed in parameter.
//
//encore:api public method=POST path=/elements
func Create(ctx context.Context, obj model.Element) error {
	fmt.Println("Element: ", obj)

	err := service.PersistElement(obj)

	return err
}

// Returns a list of all elements of kind of Element .
//
//encore:api public method=GET path=/elements/:id
func Get(ctx context.Context, id string) (*model.Element, error) {
	obj, err := service.LoadElement(id)
	common.CheckErr(err)
	fmt.Println("Elements: ", obj)
	return &obj, nil
}

// Updates the Element passed in parameter.
//
//encore:api public method=PUT path=/elements/:id
func Update(ctx context.Context, id string, obj *model.Element) error {
	obj.Id = id
	err := service.UpdateElement(*obj)
	common.CheckErr(err)
	return nil
}
