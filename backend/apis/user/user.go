package user

import (
	"context"
	"fmt"

	"encore.app/apis"
	"encore.app/common"
	"encore.app/model"
	"encore.app/service"
)

// Returns a list of all users of kind of User .
//
//encore:api public method=GET path=/users/:id
func Get(ctx context.Context, id string) (*model.User, error) {
	obj, err := service.LoadUser(id)
	common.CheckErr(err)
	fmt.Println("Users: ", obj)
	return &obj, nil
}

// Returns a list of all users of kind of User .
//
//encore:api public method=GET path=/users
func List(ctx context.Context) (apis.ListUser, error) {
	list := apis.NewListUser()
	list.Content, _ = service.LoadUsers()
	fmt.Println("Users: ", list)
	return list, nil
}

// Persists the User passed in parameter.
//
//encore:api public method=POST path=/users
func Create(ctx context.Context, obj model.User) error {
	fmt.Println("User: ", obj)

	_, err := service.PersistUser(obj)

	return err
}

// Updates the User passed in parameter.
//
//encore:api public method=PUT path=/users/:id
func Update(ctx context.Context, id string, obj *model.User) error {
	obj.Id = id
	err := service.UpdateUser(*obj)
	common.CheckErr(err)
	return nil
}
