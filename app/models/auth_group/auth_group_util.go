package auth_group

import (
	"liu/pkg/database"
)

func Get(idStr string) (authGroup AuthGroup) {
	database.DB.Where("id", idStr).First(&authGroup)
	return
}

func GetByField(field, value string) (authGroup AuthGroup) {
	database.DB.Where("? = ?", field, value).First(&authGroup)
	return
}

func All(search string) (authGroups []AuthGroup) {
	query := database.DB
	if len(search) > 0 {
		query.Where("title LIKE ?", "%"+search+"%").Find(&authGroups)
	} else {
		query.Find(&authGroups)
	}
	return
}
