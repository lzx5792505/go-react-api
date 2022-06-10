package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserUpRequest struct {
	User     string `valid:"user" json:"user,omitempty"`
	Name     string `valid:"name" json:"name,omitempty"`
	Password string `valid:"password" json:"password,omitempty"`
	Status   uint64 `valid:"status" json:"status,omitempty"`
	GroupId  uint64 `valid:"group_id" json:"group_id,omitempty"`
}

type UserUpdateStatusRequest struct {
	Status uint64 `valid:"status" json:"status,omitempty"`
}

func UserUp(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"user": []string{
			"required",
			"alpha_num",
			"between:3,20",
			"not_exists:users,user",
		},
		"name": []string{
			"required",
			"alpha_num",
			"between:3,18",
			"not_exists:users,name",
		},
		"password": []string{"min:6"},
		"group_id": []string{"digits:1"},
		"status":   []string{"required", "digits:1"},
	}
	messages := govalidator.MapData{
		"user": []string{
			"required:账号为必填项，参数名称 user",
			"alpha_num:账号格式错误，只允许数字和英文",
			"between:账号长度需在 3~18 之间",
			"not_exists:账号 已被占用",
		},
		"name": []string{
			"required:用户名为必填项，参数名称 name",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~18 之间",
			"not_exists:用户名 已被占用",
		},
		"password": []string{
			"min:密码长度需大于 6",
		},
		"status": []string{
			"digits:状态长度必须为 1 位的数字",
		},
		"group_id": []string{
			"digits:状态长度必须为 1 位的数字",
		},
	}
	errs := validate(data, rules, messages)
	return errs
}

func UserUpdateStatus(data interface{}, c *gin.Context) map[string][]string {
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
