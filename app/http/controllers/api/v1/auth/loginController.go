package auth

import (
	v1 "liu/app/http/controllers/api/v1"
	"liu/app/requests"
	"liu/pkg/auth"
	"liu/pkg/jwt"
	"liu/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseController
}

func (login *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}

	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	_user, err := auth.Attempt(request.LoginID, request.Password)

	if err != nil {
		response.Unauthorized(c, "登录失败")
	} else {
		token := jwt.NewJWT().IssueToken(_user.GetUserStringID(), _user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// 刷新 Access Token
func (login *LoginController) RefreshToken(ctx *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(ctx)

	if err != nil {
		response.Error(ctx, err, "令牌刷新失败")
	} else {
		response.JSON(ctx, gin.H{
			"token": token,
		})
	}
}
