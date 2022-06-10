package v1

import (
	"liu/app/models/user"
	"liu/app/requests"
	"liu/pkg/config"
	"liu/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseController
}

// 列表
func (us *UsersController) UserList(ctx *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(ctx, &request, requests.Pagination); !ok {
		return
	}

	data, page := user.Paginate(ctx, config.GetInt("paging.per_page"))

	response.JSON(ctx, gin.H{
		"data":  data,
		"pager": page,
	})
}

// 单个用户
func (us *UsersController) UserOne(ctx *gin.Context) {
	_user := user.Get(ctx.Param("id"))
	response.Data(ctx, _user)
}

// 保存
func (us *UsersController) UserStore(ctx *gin.Context) {
	request := requests.UserUpRequest{}
	if ok := requests.Validate(ctx, &request, requests.UserUp); !ok {
		return
	}

	_user := user.User{
		User:        request.User,
		Name:        request.Name,
		Password:    request.Password,
		Status:      request.Status,
		LastLoginAt: time.Now(),
	}

	_user.Create()

	if _user.ID > 0 {
		response.Created(ctx, _user)
	} else {
		response.Abort500(ctx, "创建失败，请稍后尝试~")
	}
}

// 修改数据
func (us *UsersController) UpdateUser(ctx *gin.Context) {
	_user := user.Get(ctx.Param("id"))
	NoticeMsg(ctx, _user.ID)

	request := requests.UserUpRequest{}
	if ok := requests.Validate(ctx, &request, requests.UserUp); !ok {
		return
	}

	_user.User = request.User
	_user.Name = request.Name
	_user.Password = request.Password
	_user.Status = request.Status

	rows := _user.Save()

	RowsMsg(ctx, rows, "更新失败，请稍后尝试~")
}

// 修改状态
func (us *UsersController) UpdateStatus(ctx *gin.Context) {
	_user := user.Get(ctx.Param("id"))
	NoticeMsg(ctx, _user.ID)

	request := requests.UserUpdateStatusRequest{}
	if ok := requests.Validate(ctx, &request, requests.UserUpdateStatus); !ok {
		return
	}

	_user.Status = request.Status
	rows := _user.Save()

	RowsMsg(ctx, rows, "更新失败，请稍后尝试~")
}

// 删除
func (us *UsersController) DelUser(ctx *gin.Context) {
	_user := user.Get(ctx.Param("id"))

	NoticeMsg(ctx, _user.ID)

	rows := _user.Delete()

	RowsMsg(ctx, rows, "删除失败，请稍后尝试~")
}
