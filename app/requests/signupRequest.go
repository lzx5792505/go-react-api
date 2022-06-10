package requests

import (
	"liu/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupUserRequest struct {
	User            string `valid:"user" json:"user"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm"`
}

func SignupUsingUser(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"user":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,user"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"user": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUserRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)

	return errs
}
