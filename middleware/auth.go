package middleware

import (
	"campushelphub/internal/common/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddlerware(tokenManager *auth.TokenManager) func(gin.HandlerFunc) gin.HandlerFunc {
	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			// 从请求头中获取token
			tokenStr := c.GetHeader("Authorization")
			if tokenStr == "" {
				c.JSON(http.StatusUnauthorized, tokenManager.Error.NewError("未授权", "unauthorized", http.StatusUnauthorized, nil))
				c.Abort()
				return
			}
			// 验证token
			claims, err := tokenManager.VerifyToken(tokenStr)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err)
				c.Abort()
				return
			}
			// 将claims添加到请求上下文中
			c.Set("claims", claims)
			next(c)
		}
	}
}
