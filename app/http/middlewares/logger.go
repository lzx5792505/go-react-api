package middlewares

import (
	"bytes"
	"io/ioutil"
	"liu/pkg/helpers"
	"liu/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBody struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBody) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// 记录请求日志
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 response 内容
		w := &responseBody{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = w
		// 获取请求数据
		var body []byte
		if ctx.Request.Body != nil {
			body, _ = ioutil.ReadAll(ctx.Request.Body)
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		// 设置开始时间
		start := time.Now()
		ctx.Next()
		// 开始记录日志的逻辑
		cost := time.Since(start)
		responStatus := ctx.Writer.Status()
		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}
		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			// 请求的内容
			logFields = append(logFields, zap.String("Request Body", string(body)))
			// 响应的内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}
		if responStatus > 400 && responStatus <= 499 {
			logger.Warn("HTTP Warning "+cast.ToString(responStatus), logFields...)
		} else if responStatus >= 500 && responStatus <= 599 {
			logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
