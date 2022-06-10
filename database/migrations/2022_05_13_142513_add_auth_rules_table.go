package migrations

import (
	"database/sql"
	"liu/app/models"
	"liu/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type AuthRule struct {
		models.BaseModel

		Name      string `gorm:"column:name;type:char(80);index;not null" json:"name"`
		Title     string `gorm:"column:title;type:char(20);not null" json:"title"`
		Pid       uint64 `gorm:"column:pid;type:mediumint(8);index;not null" json:"pid"`
		Type      uint64 `gorm:"column:type;type:tinyint(1);not null" json:"type"`
		Status    uint64 `gorm:"column:status;type:tinyint(1);not null" json:"status"`
		Condition string `gorm:"column:condition;type:char(100);not null" json:"condition"`
		Menu      uint64 `gorm:"column:menu;type:tinyint(1);not null" json:"menu"`
		Icon      string `gorm:"column:icon;type:varchar(255)" json:"icon"`
		Sort      uint64 `gorm:"column:sort;type:int(11);not null" json:"sort"`

		models.TimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AuthRule{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AuthRule{})
	}

	migrate.Add("2022_05_13_142513_add_auth_rules_table", up, down)
}
