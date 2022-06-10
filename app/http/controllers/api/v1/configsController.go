package v1

import (
	"liu/app/models/config"
	"liu/app/requests"
	"liu/pkg/response"

	"github.com/gin-gonic/gin"
)

type ConfigsController struct {
	BaseController
}

func (ctrl *ConfigsController) ConfigIndex(ctx *gin.Context) {
	configs := config.All()
	response.Data(ctx, configs)
}

func (ctrl *ConfigsController) ConfigStore(ctx *gin.Context) {
	request := requests.ConfigRequest{}
	if ok := requests.Validate(ctx, &request, requests.ConfigSave); !ok {
		return
	}

	_config := config.Config{
		Title:     request.Title,
		Name:      request.Name,
		Icp:       request.Icp,
		Copyright: request.Copyright,
	}

	_config.Create()

	if _config.ID > 0 {
		response.Created(ctx, _config)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

// 更新
func (ctrl *ConfigsController) ConfigUpdate(ctx *gin.Context) {
	_config := config.Get(ctx.Param("id"))
	if _config.ID == 0 {
		response.Abort404(ctx)
		return
	}

	// 验证数据
	request := requests.ConfigRequest{}
	if ok := requests.Validate(ctx, &request, requests.ConfigSave); !ok {
		return
	}

	_config.Title = request.Title
	_config.Name = request.Name
	_config.Icp = request.Icp
	_config.Copyright = request.Copyright

	rows := _config.Save()

	if rows > 0 {
		response.Data(ctx, _config)
	} else {
		response.Abort500(ctx, "更新失败，请稍后尝试~")
	}
}
