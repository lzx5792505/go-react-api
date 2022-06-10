package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ConfigRequest struct {
	Title     string `valid:"title" json:"title,omitempty"`
	Name      string `valid:"name" json:"name,omitempty"`
	Icp       string `valid:"icp" json:"icp,omitempty"`
	Copyright string `valid:"copyright" json:"copyright,omitempty"`
}

func ConfigSave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title":     []string{"required", "between:3,32"},
		"name":      []string{"required", "between:3,32"},
		"icp":       []string{"required", "between:3,108"},
		"copyright": []string{"required", "between:3,108"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:网站标题为必填项",
			"between:网站标题长度需在 3~32 之间",
		},
		"name": []string{
			"required:网站名称为必填项",
			"between:网站名称长度需在 3~32 之间",
		},
		"icp": []string{
			"required:ICP备案为必填项",
			"between:ICP备长度需在 3~108 之间",
		},
		"copyright": []string{
			"required:网站描述为必填项",
			"between:网站描述长度需在 3~108 之间",
		},
	}
	return validate(data, rules, messages)
}
