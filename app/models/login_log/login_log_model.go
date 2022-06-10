//Package login_log 模型
package login_log

import (
	"liu/app/models"
)

type LoginLog struct {
	models.BaseModel

	UID         uint64 `json:"uid"`
	User        string `json:"user,omitempty"`
	Name        string `json:"name,omitempty"`
	LastLoginIp string `json:"last_login_ip,omitempty"`

	models.TimestampsField
}
