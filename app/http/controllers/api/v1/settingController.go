package v1

import (
	"liu/pkg/response"
	"runtime"

	"github.com/gin-gonic/gin"
)

type SettingController struct {
	BaseController
}

type Setting struct {
	Const_os         string
	Server_software  string
	Language_version string
	Cpu              int
}

type ChartData struct {
	XData interface{}
	YData interface{}
	Title string
}

// 系统信息
func (set *SettingController) SetList(ctx *gin.Context) {
	data := Setting{
		Const_os:         runtime.GOOS,      //操作系统
		Server_software:  runtime.GOARCH,    //架构
		Language_version: runtime.Version(), //go版本
		Cpu:              runtime.NumCPU(),  //cpu数
	}

	response.Data(ctx, data)
}

//  ECharts
func (set *SettingController) ChartList(ctx *gin.Context) {
	title := "测试数据"
	xData := [...]float64{1, 2, 3, 4, 5, 6, 7, 8}
	yData := [...]float64{87, 85, 45, 41, 51, 61, 71, 81}

	data := ChartData{
		Title: title,
		XData: xData,
		YData: yData,
	}

	response.Data(ctx, data)
}
