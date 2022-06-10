package user

import (
	"liu/app/models"
	"liu/pkg/database"
	"liu/pkg/hash"
	"time"

	"github.com/spf13/cast"
)

type User struct {
	models.BaseModel

	User           string    `json:"user,omitempty"`
	Name           string    `json:"name,omitempty"`
	Head           string    `json:"head,omitempty"`
	Password       string    `json:"-"`
	LoginCount     uint64    `json:"login_count,omitempty"`
	LastLoginIp    string    `json:"last_login_ip,omitempty"`
	LastLoginAt    time.Time `json:"last_login_at,omitempty"`
	Status         uint64    `json:"status,omitempty"`
	Updatapassword uint64    `json:"updatapassword,omitempty"`

	models.TimestampsField
}

func (_user *User) Create() {
	database.DB.Create(&_user)
}

// 密码是否正确
func (_user *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, _user.Password)
}

// 获取ID
func (_user *User) GetUserStringID() string {
	return cast.ToString(_user.ID)
}

// 保存
func (user *User) Save() (rowsAffected int64) {
	res := database.DB.Save(&user)
	return res.RowsAffected
}

// 删除
func (user *User) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&user)
	return result.RowsAffected
}
