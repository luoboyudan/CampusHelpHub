package middleware

import (
	"campushelphub/internal/common/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenManager *auth.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "未授权",
			})
			c.Abort()
			return
		}
		// 验证token
		claims, err := tokenManager.VerifyToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Msg,
			})
			c.Abort()
			return
		}
		// 将UserID添加到请求上下文中
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
