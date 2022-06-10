package auth_group_access

import (
	"liu/pkg/database"
)

func Get(idStr string) (authGroupAccess AuthGroupAccess) {
	database.DB.Where("id", idStr).First(&authGroupAccess)
	return
}

func GetByField(field, value string) (authGroupAccess AuthGroupAccess) {
	database.DB.Where("? = ?", field, value).First(&authGroupAccess)
	return
}

func All() (authGroupAccesses []AuthGroupAccess) {
	database.DB.Find(&authGroupAccesses)
	return
}
