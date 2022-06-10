package v1

import (
	"liu/app/models/auth_rule"
	"liu/app/requests"
	"liu/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthRulesController struct {
	BaseController
}

// 列表
func (ctrl *AuthRulesController) RuleList(ctx *gin.Context) {
	authRules := auth_rule.All()
	response.Data(ctx, authRules)
}

// 保存
func (ctrl *AuthRulesController) RuleStore(ctx *gin.Context) {
	request := requests.AuthRuleRequest{}
	if ok := requests.Validate(ctx, &request, requests.AuthRuleSave); !ok {
		return
	}

	_rule := auth_rule.AuthRule{
		Pid:    request.Pid,
		Title:  request.Title,
		Name:   request.Name,
		Icon:   request.Icon,
		Sort:   request.Sort,
		Menu:   request.Menu,
		Status: request.Status,
	}

	_rule.Create()

	if _rule.ID > 0 {
		response.Created(ctx, _rule)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

// 更新
func (ctrl *AuthRulesController) RuleUpdate(ctx *gin.Context) {
	_rule := auth_rule.Get(ctx.Param("id"))
	NoticeMsg(ctx, _rule.ID)

	request := requests.AuthRuleRequest{}
	if ok := requests.Validate(ctx, &request, requests.AuthRuleSave); !ok {
		return
	}

	_rule.Pid = request.Pid
	_rule.Title = request.Title
	_rule.Name = request.Name
	_rule.Icon = request.Icon
	_rule.Sort = request.Sort
	_rule.Menu = request.Menu
	_rule.Status = request.Status

	rows := _rule.Save()

	RowsMsg(ctx, rows, "更新失败，请稍后尝试~")
}

// 禁用
func (ctrl *AuthRulesController) RuleStatus(ctx *gin.Context) {
	_rule := auth_rule.Get(ctx.Param("id"))
	NoticeMsg(ctx, _rule.ID)

	request := requests.RuleStatusRequest{}
	if ok := requests.Validate(ctx, &request, requests.RuleStatus); !ok {
		return
	}

	_rule.Status = request.Status

	rows := _rule.Save()

	RowsMsg(ctx, rows, "更新失败，请稍后尝试~")
}

// 删除
func (ctrl *AuthRulesController) DelRule(ctx *gin.Context) {
	_rule := auth_rule.Get(ctx.Param("id"))
	NoticeMsg(ctx, _rule.ID)

	rows := _rule.Delete()

	RowsMsg(ctx, rows, "删除失败，请稍后尝试~")
}
