package v1

import (
	"fmt"
	"liu/app/models/auth_group"
	"liu/app/requests"
	"liu/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthGroupsController struct {
	BaseController
}

func (ctrl *AuthGroupsController) AuthGroupsList(ctx *gin.Context) {
	search := ctx.Query("search")
	authGroups := auth_group.All(search)

	response.Data(ctx, authGroups)
}

func (ctrl *AuthGroupsController) AuthGroupsOne(ctx *gin.Context) {
	_group := auth_group.Get(ctx.Param("id"))
	response.Data(ctx, _group)
}

// 保存
func (ctrl *AuthGroupsController) AuthGroupsStore(ctx *gin.Context) {
	request := requests.AuthGroupRequest{}
	if ok := requests.Validate(ctx, &request, requests.AuthGroupSave); !ok {
		return
	}

	_group := auth_group.AuthGroup{
		Title:  request.Title,
		Status: 1,
	}

	_group.Create()

	if _group.ID > 0 {
		response.Created(ctx, _group)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

// 添加规则
func (ctrl *AuthGroupsController) AuthGroupsRules(ctx *gin.Context) {
	_group := auth_group.Get(ctx.Param("id"))
	
	NoticeMsg(ctx, _group.ID)

	request := requests.AuthGroupRulesRequest{}
	fmt.Println(request)
	if ok := requests.Validate(ctx, &request, requests.AuthGroupUpdateRules); !ok {
		return
	}

	_group.Rules = request.Rules

	rows := _group.Save()

	RowsMsg(ctx, rows, "更新失败，请稍后尝试~")
}

// 更新
func (ctrl *AuthGroupsController) AuthGroupsUpdate(ctx *gin.Context) {
	_group := auth_group.Get(ctx.Param("id"))
	NoticeMsg(ctx, _group.ID)

	request := requests.AuthGroupRequest{}
	if ok := requests.Validate(ctx, &request, requests.AuthGroupSave); !ok {
		return
	}

	_group.Title = request.Title

	rows := _group.Save()

	RowsMsg(ctx, rows, "更新失败，请稍后尝试~")
}

// 禁用
func (ctrl *AuthGroupsController) AuthGroupsStatus(ctx *gin.Context) {
	_group := auth_group.Get(ctx.Param("id"))
	NoticeMsg(ctx, _group.ID)

	request := requests.AuthGroupStatusRequest{}
	if ok := requests.Validate(ctx, &request, requests.AuthGroupUpdateStatus); !ok {
		return
	}

	_group.Status = request.Status

	rows := _group.Save()

	RowsMsg(ctx, rows, "更新失败，请稍后尝试~")
}

// 删除
func (ctrl *AuthGroupsController) DelAuthGroups(ctx *gin.Context) {
	_group := auth_group.Get(ctx.Param("id"))
	NoticeMsg(ctx, _group.ID)

	rows := _group.Delete()

	RowsMsg(ctx, rows, "删除失败，请稍后尝试~")
}
