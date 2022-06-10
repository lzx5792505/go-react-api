package migrations

import (
	"database/sql"
	"liu/app/models"
	"liu/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type LoginLog struct {
		models.BaseModel

		Uid         uint64 `gorm:"column:uid;type:int(11);index;not null" json:"uid"`
		User        string `gorm:"column:user;type:varchar(32);not null" json:"user"`
		Name        string `gorm:"column:name;type:varchar(32);not null" json:"name"`
		LastLoginIp string `gorm:"column:last_login_ip;type:varchar(32);not null" json:"last_login_ip"`

		models.TimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&LoginLog{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&LoginLog{})
	}

	migrate.Add("2022_05_13_142452_add_login_logs_table", up, down)
}
