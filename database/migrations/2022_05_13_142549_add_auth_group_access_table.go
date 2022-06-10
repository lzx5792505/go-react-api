package migrations

import (
	"database/sql"
	"liu/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type AuthGroupAccess struct {
		Uid     uint64 `gorm:"column:uid;type:mediumint(8);index;not null" json:"uid"`
		GroupId uint64 `gorm:"column:group_id;type:mediumint(8);index;not null" json:"group_id"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AuthGroupAccess{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AuthGroupAccess{})
	}

	migrate.Add("2022_05_13_142549_add_auth_group_access_table", up, down)
}
