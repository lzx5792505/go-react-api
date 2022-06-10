package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AuthGroupRequest struct {
	Title string `valid:"title"  json:"title,omitempty"`
}

type AuthGroupStatusRequest struct {
	Status uint64 `valid:"status" json:"status,omitempty"`
}

type AuthGroupRulesRequest struct {
	Rules string `valid:"rules" json:"rules,omitempty"`
}

func AuthGroupSave(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min_cn:2", "max_cn:18", "not_exists:auth_groups,name"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:组名为必填项",
			"min_cn:组名长度需至少 2 个字",
			"max_cn:组名长度不能超过 18 个字",
			"not_exists:组名已存在",
		},
	}
	return validate(data, rules, messages)
}

func AuthGroupUpdateStatus(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"status": []string{"required", "digits:1"},
	}
	messages := govalidator.MapData{
		"status": []string{
			"required:状态为必填项，参数名称 status",
			"digits:状态长度必须为 1 位的数字",
		},
	}
	errs := validate(data, rules, messages)
	return errs
}

func AuthGroupUpdateRules(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"rules": []string{"required", "between:1,108"},
	}
	messages := govalidator.MapData{
		"rules": []string{
			"required:规则ID为必填项，参数名称 rules",
			"between:规则ID长度需在 1~108 之间",
		},
	}
	errs := validate(data, rules, messages)
	return errs
}
