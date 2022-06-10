package v1

import (
	"liu/app/models/login_log"
	"liu/app/requests"
	"liu/pkg/config"
	"liu/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginLogsController struct {
	BaseController
}

// 列表
func (ctrl *LoginLogsController) LogIndex(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, page := login_log.Paginate(ctx, config.GetInt("paging.per_page"))

	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": page,
	})
}
