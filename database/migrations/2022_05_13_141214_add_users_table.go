package migrations

import (
	"database/sql"
	"liu/app/models"
	"liu/pkg/migrate"
	"time"

	"gorm.io/gorm"
)

func init() {
	type User struct {
		models.BaseModel

		User           string    `gorm:"column:user;type:varchar(18);not null;index" json:"user,omitempty"`
		Name           string    `gorm:"column:name;type:varchar(32);not null" json:"name,omitempty"`
		Head           string    `gorm:"column:head;type:varchar(255)" json:"head,omitempty"`
		Password       string    `gorm:"column:password;type:varchar(108);not null" json:"password,omitempty"`
		LoginCount     int64     `gorm:"column:login_count;type:int(11)" json:"login_count,omitempty"`
		LastLoginIp    string    `gorm:"column:last_login_ip;type:varchar(32);index" json:"last_login_ip,omitempty"`
		LastLoginAt    time.Time `gorm:"column:last_login_at;index" json:"last_login_at,omitempty"`
		Status         uint64    `gorm:"column:status;type:tinyint(1)" json:"status,omitempty"`
		Updatapassword string    `gorm:"column:updatapassword;type:tinyint(1)" json:"updatapassword,omitempty"`

		models.TimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_05_13_141214_add_users_table", up, down)
}
