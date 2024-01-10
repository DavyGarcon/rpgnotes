package service

/*
func PersistIcon(obj model.Icon) error {
	list, err := loadIcons()
	common.CheckErr(err)
	obj.Id = uuid.NewString()
	list = append(list, obj)
	if err := Save(DB_FILE_PREFIX+ICON_TABLE+DB_FILE_EXTENSION, list); err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func LoadIcons() ([]model.Icon, error) {
	list, err := loadIcons()
	common.CheckErr(err)
	result := make([]model.Icon, 0)
	for _, gm := range list {
		result = append(result, gm)
	}
	return result, nil
}

func UpdateIcon(e model.Icon) error {
	fmt.Println("Icon to update: ", e)
	list, err := loadIcons()
	common.CheckErr(err)
	for index, u := range list {
		if u.Id == e.Id {
			fmt.Println("Icon found: ", u)
			list[index] = e
			fmt.Println("Icon after update found: ", list[index])
			if err := Save(DB_FILE_PREFIX+ICON_TABLE+DB_FILE_EXTENSION, list); err != nil {
				log.Fatalln(err.Error())
				return err
			}
			break
		}
	}
	return nil
}

func LoadIcon(id string) (model.Icon, error) {
	list, err := loadIcons()
	common.CheckErr(err)
	for _, u := range list {
		if u.Id == id {
			return u, nil
		}
	}
	return model.Icon{}, nil
}

func loadIcons() ([]model.Icon, error) {
	list := make([]model.Icon, 0)
	if err := Load(DB_FILE_PREFIX+ICON_TABLE+DB_FILE_EXTENSION, &list); err != nil {
		fmt.Println("Load element Error: ", err)
		return nil, err
	}
	return list, nil
}
*/
