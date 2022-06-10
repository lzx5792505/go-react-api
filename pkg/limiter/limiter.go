package limiter

import (
	"liu/pkg/config"
	"liu/pkg/logger"
	"liu/pkg/redis"
	"strings"

	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// 获取 Limitor 的 Key，IP
func GetKeyIP(ctx *gin.Context) string {
	return ctx.ClientIP()
}

// 针对单个路由做限流
func GetKeyRouteWithIP(ctx *gin.Context) string {
	return routeToKeyString(ctx.FullPath()) + ctx.ClientIP()
}

// 检测请求是否超额
func CheckRate(ctx *gin.Context, key string, formatted string) (limiterlib.Context, error) {
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// 初始化存储
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiterlib.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})

	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	// limiter.Rate 对象和存储对象
	limiterObj := limiterlib.New(store, rate)
	if ctx.GetBool("limiter-once") {
		return limiterObj.Peek(ctx, key)
	} else {
		return limiterObj.Get(ctx, key)
	}
}

func routeToKeyString(name string) string {
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "/", "_")
	return name
}
