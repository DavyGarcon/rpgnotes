package service

import (
	"fmt"
	"log"

	"encore.app/common"
	"encore.app/model"
	"github.com/google/uuid"
)

func PersistTag(obj model.Tag) error {
	list, err := loadTags()
	common.CheckErr(err)
	obj.Id = uuid.NewString()
	list = append(list, obj)
	if err := Save(DB_FILE_PREFIX+TAG_TABLE+DB_FILE_EXTENSION, list); err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func LoadTags() ([]model.Tag, error) {
	list, err := loadTags()
	common.CheckErr(err)
	result := make([]model.Tag, 0)
	for _, gm := range list {
		result = append(result, gm)
	}
	return result, nil
}

func UpdateTag(e model.Tag) error {
	fmt.Println("Tag to update: ", e)
	list, err := loadTags()
	common.CheckErr(err)
	id := e.Id
	for index, u := range list {
		if u.Id == id {
			fmt.Println("Tag found: ", u)
			list[index] = e
			fmt.Println("Tag after update found: ", list[index])
			if err := Save(DB_FILE_PREFIX+TAG_TABLE+DB_FILE_EXTENSION, list); err != nil {
				log.Fatalln(err.Error())
				return err
			}
			break
		}
	}
	return nil
}

func LoadTag(id string) (model.Tag, error) {
	list, err := loadTags()
	common.CheckErr(err)
	for _, u := range list {
		if u.Id == id {
			return u, nil
		}
	}
	return model.Tag{}, nil
}

func loadTags() ([]model.Tag, error) {
	list := make([]model.Tag, 0)
	if err := Load(DB_FILE_PREFIX+TAG_TABLE+DB_FILE_EXTENSION, &list); err != nil {
		fmt.Println("Load element Error: ", err)
		return nil, err
	}
	return list, nil
}
