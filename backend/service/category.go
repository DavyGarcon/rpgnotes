package service

import (
	"fmt"
	"log"

	"encore.app/common"
	"encore.app/model"
	"github.com/google/uuid"
)

func PersistCategory(obj model.Category) error {
	list, err := loadCategories()
	common.CheckErr(err)
	obj.Id = uuid.NewString()
	list = append(list, obj)
	if err := Save(DB_FILE_PREFIX+CATEGORY_TABLE+DB_FILE_EXTENSION, list); err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func LoadCategories() ([]model.Category, error) {
	list, err := loadCategories()
	common.CheckErr(err)
	result := make([]model.Category, 0)
	for _, gm := range list {
		result = append(result, gm)
	}
	return result, nil
}

func UpdateCategory(e model.Category) error {
	fmt.Println("Category to update: ", e)
	list, err := loadCategories()
	common.CheckErr(err)
	id := e.Id
	for index, u := range list {
		if u.Id == id {
			fmt.Println("Category found: ", u)
			list[index] = e
			fmt.Println("Category after update found: ", list[index])
			if err := Save(DB_FILE_PREFIX+CATEGORY_TABLE+DB_FILE_EXTENSION, list); err != nil {
				log.Fatalln(err.Error())
				return err
			}
			break
		}
	}
	return nil
}

func FindCategoriesByCampaignId(id string) ([]model.Category, error) {
	campaign, err := LoadCampaign(id)
	common.CheckErr(err)
	list, err := loadCategories()
	result := make([]model.Category, 0)
	common.CheckErr(err)
	fmt.Println("Id parameter: ", id)
	for _, obj := range list {
		fmt.Println("obj.CampaignId: ", obj.CampaignId)
		fmt.Println("obj.CategoryId: ", obj.CategoryId)
		if campaign.Id == obj.CampaignId {
			result = append(result, obj)
			continue
		}
	}
	return result, nil
}

func FindCategoriesByCategoryId(id string) ([]model.Category, error) {
	list, err := loadCategories()
	result := make([]model.Category, 0)
	common.CheckErr(err)
	for _, obj := range list {
		if id == obj.CategoryId {
			result = append(result, obj)
			continue
		}
	}
	return result, nil
}
func LoadCategory(id string) (model.Category, error) {
	list, err := loadCategories()
	common.CheckErr(err)
	for _, u := range list {
		if u.Id == id {
			return u, nil
		}
	}
	return model.Category{}, nil
}

func loadCategories() ([]model.Category, error) {
	list := make([]model.Category, 0)
	if err := Load(DB_FILE_PREFIX+CATEGORY_TABLE+DB_FILE_EXTENSION, &list); err != nil {
		fmt.Println("Load element Error: ", err)
		return nil, err
	}
	return list, nil
}
