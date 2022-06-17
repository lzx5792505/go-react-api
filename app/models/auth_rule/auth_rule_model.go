//Package auth_rule 模型
package auth_rule

import (
	"liu/app/models"
	"liu/pkg/database"
)

type AuthRule struct {
	models.BaseModel

	Pid       string `json:"pid,omitempty"`
	Title     string `json:"title,omitempty"`
	Name      string `json:"name,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Sort      uint64 `json:"sort,omitempty"`
	Menu      uint64 `json:"menu,omitempty"`
	Status    uint64 `json:"status,omitempty"`
	Type      uint64 `json:"type,omitempty"`
	Condition string `json:"condition,omitempty"`

	models.TimestampsField
}

func (_rule *AuthRule) Create() {
	database.DB.Create(&_rule)
}

func (_rule *AuthRule) Save() (rowsAffected int64) {
	result := database.DB.Save(&_rule)
	return result.RowsAffected
}

func (_rule *AuthRule) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&_rule)
	return result.RowsAffected
}
