package setup

import (
	apiCommon "campushelphub/api/common"
	"campushelphub/api/handlerset"
	"campushelphub/internal/common/auth"
	"campushelphub/middleware"

	"github.com/gin-gonic/gin"
)

func NewEngine(handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, baseHandler *apiCommon.Handler) *gin.Engine {
	gin := gin.Default()
	SetupRouter(gin, handlerSet, tokenManager, baseHandler)
	return gin
}

func SetupRouter(app *gin.Engine, handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, baseHandler *apiCommon.Handler) {
	app.GET("/docs/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	v1 := app.Group("/v1")
	v1.Use(middleware.AuthMiddleware(tokenManager, baseHandler))
	{
		v1.POST("/user", handlerSet.UserHandler.CreateUser)
	}
}
