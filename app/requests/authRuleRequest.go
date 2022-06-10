package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AuthRuleRequest struct {
	Pid    uint64 `valid:"pid" json:"pid,omitempty"`
	Title  string `valid:"title" json:"title,omitempty"`
	Name   string `valid:"name" json:"name,omitempty"`
	Icon   string `valid:"icon" json:"icon,omitempty"`
	Sort   uint64 `valid:"sort" json:"sort,omitempty"`
	Menu   uint64 `valid:"menu" json:"menu,omitempty"`
	Status uint64 `valid:"status" json:"status,omitempty"`
}

type RuleStatusRequest struct {
	Status uint64 `valid:"status" json:"status,omitempty"`
}

func AuthRuleSave(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"pid":    []string{"required", "not_exists:auth_rules,pid"},
		"title":  []string{"required", "min_cn:2", "max_cn:18", "not_exists:auth_rules,title"},
		"name":   []string{"required", "min_cn:2", "max_cn:18", "not_exists:auth_rules,name"},
		"icon":   []string{"alpha_num", "between:3,32"},
		"sort":   []string{"digits:1"},
		"menu":   []string{"digits:1"},
		"status": []string{"digits:1"},
	}
	messages := govalidator.MapData{
		"pid": []string{
			"required:顶级菜单为必填项",
			"not_exists:顶级菜单名称已存在",
		},
		"title": []string{
			"required:顶级菜单为必填项",
			"not_exists:顶级菜单名称已存在",
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
		"name": []string{
			"required:顶级菜单为必填项",
			"not_exists:顶级菜单名称已存在",
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
		"icon": []string{
			"alpha_num:图标格式错误，只允许数字和英文",
			"between:图标长度需在 3~32 之间",
		},
		"sort": []string{
			"digits:状态长度必须为 1 位的数字",
		},
		"menu": []string{
			"digits:状态长度必须为 1 位的数字",
		},
		"status": []string{
			"digits:状态长度必须为 1 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func RuleStatus(data interface{}, c *gin.Context) map[string][]string {
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
