package user

import (
	"liu/pkg/app"
	"liu/pkg/database"
	"liu/pkg/paginator"

	"github.com/gin-gonic/gin"
)

// 登录
func GetByMulti(loginID string) (_user User) {
	database.DB.
		Select("id", "user", "name", "password").
		Where("user = ? AND status = ?", loginID, 1).
		First(&_user)

	return
}

// 单条
func Get(idStr string) (_user User) {
	database.DB.Where("id", idStr).First(&_user)

	return
}

// 全部
func All() (_users []User) {
	database.DB.
		Find(&_users)

	return
}

// 分页内容
func Paginate(ctx *gin.Context, perPage int) (_users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(User{}),
		&_users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)

	return
}
