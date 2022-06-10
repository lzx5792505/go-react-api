package migrations

import (
	"database/sql"
	"liu/app/models"
	"liu/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type Config struct {
		models.BaseModel

		Name      string `gorm:"column:name;type:varchar(32);index;not null" json:"name"`
		Title     string `gorm:"column:title;type:varchar(32);index;not null" json:"title"`
		Icp       string `gorm:"column:icp;type:varchar(108);not null" json:"icp"`
		Copyright string `gorm:"column:copyright;type:varchar(108);not null" json:"copyright"`

		models.TimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Config{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Config{})
	}

	migrate.Add("2022_05_13_141314_add_configs_table", up, down)
}
