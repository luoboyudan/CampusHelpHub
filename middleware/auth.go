package middleware

import (
	"campushelphub/internal/common/auth"
	"campushelphub/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenManager *auth.TokenManager, logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": AuthNoTokenMsg,
			})
			logInfo := &log.BusinessLogInfo{
				BusinessType: AuthBusinessType,
				ClientIP:     c.ClientIP(),
				Status:       AuthFailStatus,
				Extra:        map[string]interface{}{"error": AuthNoTokenMsg},
			}
			logger.Info(logInfo)
			c.Abort()
			return
		}
		// 验证token
		claims, err := tokenManager.VerifyToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Msg,
			})
			logInfo := &log.BusinessLogInfo{
				BusinessType: AuthBusinessType,
				ClientIP:     c.ClientIP(),
				Status:       AuthFailStatus,
				Extra:        map[string]interface{}{"error": err.Msg},
			}
			logger.Info(logInfo)
			c.Abort()
			return
		}
		// 将UserID添加到请求上下文中
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
