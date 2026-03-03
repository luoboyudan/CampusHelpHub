package setup

import (
	"campushelphub/api/handlerset"
	"campushelphub/internal/common/auth"
	"campushelphub/internal/log"
	"campushelphub/middleware"

	"github.com/gin-gonic/gin"
)

func NewEngine(handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, logger *log.Logger) *gin.Engine {
	gin := gin.Default()
	SetupRouter(gin, handlerSet, tokenManager, logger)
	return gin
}

func SetupRouter(app *gin.Engine, handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, logger *log.Logger) {
	app.GET("/docs/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	v1 := app.Group("/v1")
	{
		SetupUserRouter(v1, handlerSet, tokenManager, logger)
		SetupEncryptionRouter(v1, handlerSet)
	}
}

func SetupUserRouter(routerGroup *gin.RouterGroup, handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, logger *log.Logger) {
	openGroup := routerGroup.Group("/user")
	{
		openGroup.POST("/register", handlerSet.UserHandler.CreateUser)
		openGroup.POST("/login", handlerSet.UserHandler.LoginUser)
	}
	authGroup := routerGroup.Group("/user-auth", middleware.AuthMiddleware(tokenManager, logger))
	{
		authGroup.POST("/verify", handlerSet.UserHandler.VerifyUser)
	}
}

func SetupEncryptionRouter(routerGroup *gin.RouterGroup, handlerSet *handlerset.HandlerSet) {
	openGroup := routerGroup.Group("/encryption")
	{
		openGroup.GET("/public-key", handlerSet.EncryptionHandler.GetPublicKey)
	}
}
