package auth_rule

import (
	"liu/pkg/database"
)

func Get(idStr string) (authRule AuthRule) {
	database.DB.Where("id", idStr).First(&authRule)
	return
}

func GetByField(field, value string) (authRule AuthRule) {
	database.DB.Where("? = ?", field, value).First(&authRule)
	return
}

func All() (authRules []AuthRule) {
	database.DB.Find(&authRules)
	return
}
