package auth_rule

import (
	"liu/pkg/database"
)

func Get(idStr string) (_rule AuthRule) {
	database.DB.Where("id", idStr).First(&_rule)
	return
}

func GetByField(field, value string) (_rule AuthRule) {
	database.DB.Where("? = ?", field, value).First(&_rule)
	return
}

func All() (_rule []AuthRule) {
	database.DB.Find(&_rule)
	return
}
