package setup

import (
	"campushelphub/api/handlerset"
	"campushelphub/internal/common/auth"

	"github.com/gin-gonic/gin"
)

func NewEngine(handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager) *gin.Engine {
	gin := gin.Default()
	SetupRouter(gin, handlerSet, tokenManager)
	return gin
}

func SetupRouter(app *gin.Engine, handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager) {
	app.GET("/docs/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	v1 := app.Group("/v1")
	{
		SetupUserRouter(v1, handlerSet, tokenManager)
	}
}

func SetupUserRouter(routerGroup *gin.RouterGroup, handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager) {
	looseGroup := routerGroup.Group("/user")
	{
		looseGroup.POST("/register", handlerSet.UserHandler.CreateUser)
	}
}
