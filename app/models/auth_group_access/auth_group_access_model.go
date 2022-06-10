//Package auth_group_access 模型
package auth_group_access

import (
	"liu/app/models"
	"liu/pkg/database"
)

type AuthGroupAccess struct {
	models.BaseModel

	Uid     uint64 `json:"uid"`
	GroupId uint64 `json:"group_id"`

	models.TimestampsField
}

func (authGroupAccess *AuthGroupAccess) Create() {
	database.DB.Create(&authGroupAccess)
}

func (authGroupAccess *AuthGroupAccess) Save() (rowsAffected int64) {
	result := database.DB.Save(&authGroupAccess)
	return result.RowsAffected
}

func (authGroupAccess *AuthGroupAccess) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&authGroupAccess)
	return result.RowsAffected
}
