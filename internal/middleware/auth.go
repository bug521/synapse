package middleware

import (
	"net/http"
	"strings"
	"synapse/internal/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从Header中获取token
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "缺少认证令牌")
			return
		}

		// 检查Bearer token格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "令牌格式错误")
			return
		}

		// 解析令牌
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "未授权", "无效令牌")
			return
		}

		// 将用户ID存储到上下文中
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
