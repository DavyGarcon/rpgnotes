package service

import (
	"fmt"
	"log"

	"encore.app/common"
	"encore.app/model"
	"github.com/google/uuid"
)

func PersistCampaign(obj model.Campaign) (*string, error) {
	list, err := loadCampaigns()
	common.CheckErr(err)
	obj.Id = uuid.NewString()
	list = append(list, obj)
	if err := Save(DB_FILE_PREFIX+CAMPAIGN_TABLE+DB_FILE_EXTENSION, list); err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	return &obj.Id, nil
}

func LoadCampaigns() ([]model.Campaign, error) {
	list, err := loadCampaigns()
	common.CheckErr(err)
	result := make([]model.Campaign, 0)
	for _, gm := range list {
		result = append(result, gm)
	}
	return result, nil
}

func UpdateCampaign(e model.Campaign) error {
	fmt.Println("Campaign to update: ", e)
	list, err := loadCampaigns()
	common.CheckErr(err)
	id := e.Id
	for index, u := range list {
		if u.Id == id {
			fmt.Println("Campaign found: ", u)
			list[index] = e
			fmt.Println("Campaign after update found: ", list[index])
			if err := Save(DB_FILE_PREFIX+CAMPAIGN_TABLE+DB_FILE_EXTENSION, list); err != nil {
				log.Fatalln(err.Error())
				return err
			}
			break
		}
	}
	return nil
}

func LoadCampaign(id string) (model.Campaign, error) {
	list, err := loadCampaigns()
	common.CheckErr(err)
	for _, u := range list {
		if u.Id == id {
			return u, nil
		}
	}
	return model.Campaign{}, nil
}

func loadCampaigns() ([]model.Campaign, error) {
	list := make([]model.Campaign, 0)
	if err := Load(DB_FILE_PREFIX+CAMPAIGN_TABLE+DB_FILE_EXTENSION, &list); err != nil {
		fmt.Println("Load element Error: ", err)
		return nil, err
	}
	return list, nil
}
