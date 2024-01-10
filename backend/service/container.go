package service

import (
	"fmt"
	"log"

	"encore.app/common"
	"encore.app/model"
	"github.com/google/uuid"
)

func PersistContainer(obj model.Container) error {
	list, err := loadContainers()
	common.CheckErr(err)
	obj.Id = uuid.NewString()
	list = append(list, obj)
	if err := Save(DB_FILE_PREFIX+CONTAINER_TABLE+DB_FILE_EXTENSION, list); err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func LoadContainers() ([]model.Container, error) {
	list, err := loadContainers()
	common.CheckErr(err)
	result := make([]model.Container, 0)
	for _, gm := range list {
		result = append(result, gm)
	}
	return result, nil
}

func UpdateContainer(e model.Container) error {
	fmt.Println("Container to update: ", e)
	list, err := loadContainers()
	common.CheckErr(err)
	id := e.Id
	for index, u := range list {
		if u.Id == id {
			fmt.Println("Container found: ", u)
			list[index] = e
			fmt.Println("Container after update found: ", list[index])
			if err := Save(DB_FILE_PREFIX+CONTAINER_TABLE+DB_FILE_EXTENSION, list); err != nil {
				log.Fatalln(err.Error())
				return err
			}
			break
		}
	}
	return nil
}

func FindContainersByCampaignId(id string) ([]model.Container, error) {
	list, err := loadContainers()
	result := make([]model.Container, 0)
	common.CheckErr(err)
	fmt.Println("Id parameter: ", id)
	for _, obj := range list {
		fmt.Println("obj.CampaignId: ", obj.CampaignId)
		fmt.Println("obj.CategoryId: ", obj.CategoryId)
		if id == obj.CampaignId {
			result = append(result, obj)
			continue
		}
	}
	return result, nil
}

func FindContainersByCategoryId(id string) ([]model.Container, error) {
	list, err := loadContainers()
	result := make([]model.Container, 0)
	common.CheckErr(err)
	fmt.Println("Searched CategoryId: ", id)
	for _, obj := range list {
		if id == obj.CategoryId {
			result = append(result, obj)
			continue
		}
	}
	return result, nil
}
func LoadContainer(id string) (model.Container, error) {
	list, err := loadContainers()
	common.CheckErr(err)
	for _, u := range list {
		if u.Id == id {
			return u, nil
		}
	}
	return model.Container{}, nil
}

func loadContainers() ([]model.Container, error) {
	list := make([]model.Container, 0)
	if err := Load(DB_FILE_PREFIX+CONTAINER_TABLE+DB_FILE_EXTENSION, &list); err != nil {
		fmt.Println("Load element Error: ", err)
		return nil, err
	}
	return list, nil
}
