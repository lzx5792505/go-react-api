//Package auth_rule 模型
package auth_rule

import (
	"liu/app/models"
	"liu/pkg/database"
)

type AuthRule struct {
	models.BaseModel

	Pid       uint64 `json:"pid,omitempty"`
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

func (authRule *AuthRule) Create() {
	database.DB.Create(&authRule)
}

func (authRule *AuthRule) Save() (rowsAffected int64) {
	result := database.DB.Save(&authRule)
	return result.RowsAffected
}

func (authRule *AuthRule) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&authRule)
	return result.RowsAffected
}
