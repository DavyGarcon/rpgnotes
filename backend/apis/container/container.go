package container

import (
	"context"
	"fmt"

	"encore.app/apis"
	"encore.app/common"
	"encore.app/model"
	"encore.app/service"
)

// Returns a list of all containers of kind of Container .
//
//encore:api public method=GET path=/containers/:id
func Get(ctx context.Context, id string) (*model.Container, error) {
	obj, err := service.LoadContainer(id)
	common.CheckErr(err)
	fmt.Println("Containers: ", obj)
	return &obj, nil
}

// Updates the Container passed in parameter.
//
//encore:api public method=PUT path=/containers/:id
func Update(ctx context.Context, id string, obj *model.Container) error {
	obj.Id = id
	err := service.UpdateContainer(*obj)
	common.CheckErr(err)
	return nil
}

// Persists the Element passed in parameter.
//
//encore:api public method=POST path=/containers/:id/elements
func AddElement(ctx context.Context, id string, obj model.Element) error {
	fmt.Println("Element: ", obj)
	obj.ContainerId = id
	err := service.PersistElement(obj)

	return err
}

// Returns a list of all elements of kind of Element .
//
//encore:api public method=GET path=/containers/:containerId/elements
func ListElement(ctx context.Context, containerId string) (apis.ListElement, error) {
	list := apis.NewListElement()
	list.Content, _ = service.FindElements(containerId)
	fmt.Println("Elements: ", list)
	return list, nil
}

// Updates the Element passed in parameter.
//
//encore:api public method=PUT path=/containers/:containerId/elements/:elementId
func UpdateElement(ctx context.Context, containerId string, elementId string, obj *model.Element) error {
	obj.Id = elementId
	obj.ContainerId = containerId
	err := service.UpdateElement(*obj)
	common.CheckErr(err)
	return nil
}

// Updates the Campaign passed in parameter.
//
//encore:api public method=PUT path=/containers/:containerId/elements/:elementId/load_icon
func GetIconContent(ctx context.Context, containerId string, elementId string, queryParams *apis.FileParams) (apis.FileContent, error) {
	element, err := service.LoadElement(elementId)
	common.CheckErr(err)
	if queryParams != nil && queryParams.Decoded {
		content, err := common.GetDecodedFileContent(element.Icon.Path)
		common.CheckErr(err)
		return apis.FileContent{EncodedContent: nil, DecodedContent: &content}, nil
	}
	content, err := common.GetEncodedFileContent(element.Icon.Path)
	common.CheckErr(err)
	return apis.FileContent{DecodedContent: nil, EncodedContent: &content}, nil
}
