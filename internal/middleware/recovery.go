package middleware

import (
	"runtime/debug"
	"synapse/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录堆栈信息
				stack := string(debug.Stack())
				zap.L().Error("panic recovered",
					zap.Any("error", err),
					zap.String("request", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.String("stack", stack),
				)

				// 返回错误响应
				utils.ErrorResponse(c, utils.CodeServerError, "服务器内部错误", "服务暂时不可用")
				c.Abort()
			}
		}()
		c.Next()
	}
}
