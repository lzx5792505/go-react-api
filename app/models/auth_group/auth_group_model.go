//Package auth_group 模型
package auth_group

import (
	"liu/app/models"
	"liu/pkg/database"
)

type AuthGroup struct {
	models.BaseModel

	Title  string `json:"title,omitempty"`
	Status uint64 `json:"status,omitempty"`
	Rules  string `json:"rules,omitempty"`

	models.TimestampsField
}

func (authGroup *AuthGroup) Create() {
	database.DB.Create(&authGroup)
}

func (authGroup *AuthGroup) Save() (rowsAffected int64) {
	result := database.DB.Save(&authGroup)
	return result.RowsAffected
}

func (authGroup *AuthGroup) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&authGroup)
	return result.RowsAffected
}
