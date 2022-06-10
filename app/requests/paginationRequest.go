package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Sort    string `valid:"sort" form:"sort" json:"sort,omitempty"`
	Order   string `valid:"order" form:"order" json:"order,omitempty"`
	PerPage string `valid:"per_page" form:"per_page" json:"per_page,omitempty"`
	Search  string `valid:"search" form:"search" json:"search,omitempty"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"sort":     []string{"in:id,sort,created_at,updated_at,last_login_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
		"search":   []string{"between:1,18"},
	}
	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 id,sort,created_at,updated_at,last_login_at",
		},
		"order": []string{
			"in:排序规则仅支持 asc（正序）,desc（倒序）",
		},
		"per_page": []string{
			"numeric_between:每页条数的值介于 2~100 之间",
		},
		"search": []string{
			"between:搜索字段长度需在 1~18 之间",
		},
	}

	return validate(data, rules, messages)
}
