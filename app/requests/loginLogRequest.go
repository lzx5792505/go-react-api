package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginLogRequest struct {
	Search string `valid:"search" form:"search,omitempty" json:"search,omitempty" `
}

func LoginLogSave(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"search": []string{"required", "between:1,18"},
	}
	messages := govalidator.MapData{
		"search": []string{
			"required:搜索字段为必填项，参数名称 search",
			"between:搜索字段长度需在 1~18 之间",
		},
	}
	return validate(data, rules, messages)
}
