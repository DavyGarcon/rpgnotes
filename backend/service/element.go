package service

import (
	"fmt"
	"log"

	"encore.app/common"
	"encore.app/model"
	"github.com/google/uuid"
)

func PersistElement(obj model.Element) error {
	list, err := loadElements()
	common.CheckErr(err)
	obj.Id = uuid.NewString()
	list = append(list, obj)
	if err := Save(DB_FILE_PREFIX+ELEMENT_TABLE+DB_FILE_EXTENSION, list); err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func LoadElements() ([]model.Element, error) {
	list, err := loadElements()
	common.CheckErr(err)
	result := make([]model.Element, 0)
	for _, gm := range list {
		result = append(result, gm)
	}
	return result, nil
}

func UpdateElement(e model.Element) error {
	fmt.Println("Element to update: ", e)
	list, err := loadElements()
	common.CheckErr(err)
	id := e.Id
	for index, u := range list {
		if u.Id == id {
			fmt.Println("Element found: ", u)
			list[index] = e
			fmt.Println("Element after update found: ", list[index])
			if err := Save(DB_FILE_PREFIX+ELEMENT_TABLE+DB_FILE_EXTENSION, list); err != nil {
				log.Fatalln(err.Error())
				return err
			}
			break
		}
	}
	return nil
}

func FindElements(containerId string) ([]model.Element, error) {
	list, err := loadElements()
	result := make([]model.Element, 0)
	common.CheckErr(err)
	if &containerId != nil {
		for _, obj := range list {
			if obj.ContainerId == containerId {
				result = append(result, obj)
				continue
			}
		}
	}
	return result, nil
}

func LoadElement(id string) (model.Element, error) {
	list, err := loadElements()
	common.CheckErr(err)
	for _, u := range list {
		if u.Id == id {
			return u, nil
		}
	}
	return model.Element{}, nil
}

func loadElements() ([]model.Element, error) {
	list := make([]model.Element, 0)
	if err := Load(DB_FILE_PREFIX+ELEMENT_TABLE+DB_FILE_EXTENSION, &list); err != nil {
		fmt.Println("Load element Error: ", err)
		return nil, err
	}
	return list, nil
}
