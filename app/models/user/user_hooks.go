package user

import (
	"liu/pkg/hash"

	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (_user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(_user.Password) {
		_user.Password = hash.BcryptHash(_user.Password)
	}
	return
}
