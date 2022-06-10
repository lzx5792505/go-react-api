package auth

import (
	"errors"
	"liu/app/models/user"
	"liu/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 尝试登录
func Attempt(name string, password string) (user.User, error) {
	_user := user.GetByMulti(name)

	if _user.ID == 0 {
		return user.User{}, errors.New("账号不存在 或者 被管理员禁用")
	}
	if !_user.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return _user, nil
}

// 从 gin.context 中获取当前登录用户
func CurrentUser(ctx *gin.Context) user.User {
	_user, ok := ctx.MustGet("current_user").(user.User)

	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}

	return _user
}

// 从 gin.context 中获取当前登录用户 ID
func CurrentUID(ctx *gin.Context) string {
	return ctx.GetString("current_user_id")
}
