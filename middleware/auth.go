package middleware

import (
	apiCommon "campushelphub/api/common"
	"campushelphub/internal/common/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenManager *auth.TokenManager, baseHandler *apiCommon.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			baseHandler.ErrorResponse(c, baseHandler.Error.NewError("未授权", "unauthorized", http.StatusUnauthorized, nil))
			c.Abort()
			return
		}
		// 验证token
		claims, err := tokenManager.VerifyToken(tokenStr)
		if err != nil {
			baseHandler.ErrorResponse(c, err)
			c.Abort()
			return
		}
		// 将UserID添加到请求上下文中
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
