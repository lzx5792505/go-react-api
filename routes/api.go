package routes

import (
	controllers "liu/app/http/controllers/api/v1"
	"liu/app/http/controllers/api/v1/auth"
	"liu/app/http/middlewares"
	"liu/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	var v1 *gin.RouterGroup

	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	v1.Use(middlewares.LimitIP("20000-H"))
	{
		// 权限相关
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 登录
			login := new(auth.LoginController)
			authGroup.POST("/login", middlewares.GuestJWT(), login.LoginByPassword)
			// 刷新token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), login.RefreshToken)

			// 验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify/code", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
		}

		// 系统信息相关
		setGroup := v1.Group("/seeting")
		{
			set := new(controllers.SettingController)
			// 系统信息
			setGroup.GET("", middlewares.AuthJWT(), set.SetList)
			// ECharts
			setGroup.GET("/chart", middlewares.AuthJWT(), set.ChartList)
		}

		// 用户相关
		usersGroup := v1.Group("/users")
		{
			uc := new(controllers.UsersController)
			// 获取所有用户 & 搜索
			usersGroup.GET("", middlewares.AuthJWT(), uc.UserList)
			// 单个用户
			usersGroup.GET("/one/:id", middlewares.AuthJWT(), uc.UserOne)
			// 保存
			usersGroup.POST("/store", middlewares.AuthJWT(), uc.UserStore)
			// 更新
			usersGroup.PUT("/update/:id", middlewares.AuthJWT(), uc.UpdateUser)
			// 更新状态
			usersGroup.PUT("/status/:id", middlewares.AuthJWT(), uc.UpdateStatus)
			// 删除
			usersGroup.DELETE("/:id", middlewares.AuthJWT(), uc.DelUser)
		}

		// 配置相关
		configGroup := v1.Group("/config")
		{
			cg := new(controllers.ConfigsController)
			// 获取所有配置
			configGroup.GET("", middlewares.AuthJWT(), cg.ConfigIndex)
			// 保存
			configGroup.POST("/store", middlewares.AuthJWT(), cg.ConfigStore)
			// 更新
			configGroup.PUT("/:id", middlewares.AuthJWT(), cg.ConfigUpdate)
		}

		// 日志相关
		logGroup := v1.Group("/log")
		{
			ll := new(controllers.LoginLogsController)
			// 获取所有日志
			logGroup.GET("", middlewares.AuthJWT(), ll.LogIndex)
		}

		// 菜单相关
		menuGroup := v1.Group("/menu")
		{
			ar := new(controllers.AuthRulesController)
			// 获取所有菜单
			menuGroup.GET("", middlewares.AuthJWT(), ar.RuleList)
			// 单个用户
			menuGroup.GET("/one/:id", middlewares.AuthJWT(), ar.RuleOne)
			// 添加
			menuGroup.POST("/store", middlewares.AuthJWT(), ar.RuleStore)
			// 更新
			menuGroup.PUT("/update/:id", middlewares.AuthJWT(), ar.RuleUpdate)
			// 更新状态
			menuGroup.PUT("/status/:id", middlewares.AuthJWT(), ar.RuleStatus)
			// 删除
			menuGroup.DELETE("/:id", middlewares.AuthJWT(), ar.DelRule)
		}

		// 用户组相关
		groupGroup := v1.Group("/group")
		{
			ag := new(controllers.AuthGroupsController)
			// 获取所有菜单
			groupGroup.GET("", middlewares.AuthJWT(), ag.AuthGroupsList)
			// 单个用户
			groupGroup.GET("/one/:id", middlewares.AuthJWT(), ag.AuthGroupsOne)
			// 添加
			groupGroup.POST("/store", middlewares.AuthJWT(), ag.AuthGroupsStore)
			// 添加规则
			groupGroup.POST("/store/rule/:id", middlewares.AuthJWT(), ag.AuthGroupsRules)
			// 更新
			groupGroup.PUT("/update/:id", middlewares.AuthJWT(), ag.AuthGroupsUpdate)
			// 更新状态
			groupGroup.PUT("/status/:id", middlewares.AuthJWT(), ag.AuthGroupsStatus)
			// 删除
			groupGroup.DELETE("/:id", middlewares.AuthJWT(), ag.DelAuthGroups)
		}
	}
}
