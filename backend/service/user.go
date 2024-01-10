package service

import (
	"fmt"
	"log"

	"encore.app/common"
	"encore.app/model"
	"github.com/google/uuid"
)

func PersistUser(obj model.User) (*string, error) {
	list, err := loadUsers()
	common.CheckErr(err)
	obj.Id = uuid.NewString()
	list = append(list, obj)
	if err := Save(DB_FILE_PREFIX+USER_TABLE+DB_FILE_EXTENSION, list); err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	return &obj.Id, nil
}

func LoadUsers() ([]model.User, error) {
	list, err := loadUsers()
	common.CheckErr(err)
	result := make([]model.User, 0)
	for _, gm := range list {
		result = append(result, gm)
	}
	return result, nil
}

func UpdateUser(e model.User) error {
	fmt.Println("User to update: ", e)
	list, err := loadUsers()
	common.CheckErr(err)
	id := e.Id
	for index, u := range list {
		if u.Id == id {
			fmt.Println("User found: ", u)
			list[index] = e
			fmt.Println("User after update found: ", list[index])
			if err := Save(DB_FILE_PREFIX+USER_TABLE+DB_FILE_EXTENSION, list); err != nil {
				log.Fatalln(err.Error())
				return err
			}
			break
		}
	}
	return nil
}

func LoadUser(id string) (model.User, error) {
	list, err := loadUsers()
	common.CheckErr(err)
	for _, u := range list {
		if u.Id == id {
			return u, nil
		}
	}
	return model.User{}, nil
}

func loadUsers() ([]model.User, error) {
	list := make([]model.User, 0)
	if err := Load(DB_FILE_PREFIX+USER_TABLE+DB_FILE_EXTENSION, &list); err != nil {
		fmt.Println("Load element Error: ", err)
		return nil, err
	}
	return list, nil
}
