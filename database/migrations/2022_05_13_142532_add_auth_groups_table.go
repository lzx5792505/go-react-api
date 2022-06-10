package migrations

import (
	"database/sql"
	"liu/app/models"
	"liu/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type AuthGroup struct {
		models.BaseModel

		Title  string `gorm:"column:title;type:char(100);index;not null" json:"title"`
		Status uint64 `gorm:"column:status;type:tinyint(1);not null" json:"status"`
		Rules  string `gorm:"column:rules;type:char(80);not null" json:"rules"`

		models.TimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AuthGroup{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AuthGroup{})
	}

	migrate.Add("2022_05_13_142532_add_auth_groups_table", up, down)
}
