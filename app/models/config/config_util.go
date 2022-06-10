package config

import (
	"liu/pkg/database"
)

func Get(idStr string) (config Config) {
	database.DB.Where("id", idStr).First(&config)
	return
}

func GetByField(field, value string) (config Config) {
	database.DB.Where("? = ?", field, value).First(&config)
	return
}

func All() (configs []Config) {
	database.DB.Order("created_at desc").Find(&configs)
	return
}
