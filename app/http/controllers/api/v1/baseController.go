package v1

import (
	"liu/pkg/response"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// 500提示
func RowsMsg(ctx *gin.Context, rows int64, msg string) {
	if rows > 0 {
		response.Success(ctx)
		return
	} else {
		response.Abort500(ctx, msg)
	}
}

// 404提示
func NoticeMsg(ctx *gin.Context, data interface{}) {
	if data == 0 {
		response.Abort404(ctx)
		return
	}
}
