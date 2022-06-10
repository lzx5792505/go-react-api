package login_log

import (
	"liu/pkg/app"
	"liu/pkg/database"
	"liu/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Paginate(ctx *gin.Context, perPage int) (loginLogs []LoginLog, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(LoginLog{}),
		&loginLogs,
		app.V1URL(database.TableName(&LoginLog{})),
		perPage,
	)
	return
}
