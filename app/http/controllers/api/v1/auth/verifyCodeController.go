package auth

import (
	v1 "liu/app/http/controllers/api/v1"
	"liu/pkg/captcha"
	"liu/pkg/logger"
	"liu/pkg/response"

	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseController
}

func (v *VerifyCodeController) ShowCaptcha(ctx *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	response.JSON(ctx, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
