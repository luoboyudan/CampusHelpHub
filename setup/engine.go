package setup

import (
	"campushelphub/api/handlerset"
	"campushelphub/internal/common/auth"
	"campushelphub/internal/log"
	"campushelphub/internal/service"
	"campushelphub/middleware"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	Engine        *gin.Engine
	ChromeService *service.ChromeService
}

func NewEngine(handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, logger *log.Logger, chromeService *service.ChromeService) *Engine {
	gin := gin.Default()
	setupRouter(&Engine{Engine: gin}, handlerSet, tokenManager, logger)
	return &Engine{Engine: gin, ChromeService: chromeService}
}

func setupRouter(app *Engine, handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, logger *log.Logger) {
	app.Engine.GET("/docs/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	v1 := app.Engine.Group("/v1")
	{
		setupUserRouter(v1, handlerSet, tokenManager, logger)
		setupEncryptionRouter(v1, handlerSet)
		setupCompetitionRouter(v1, handlerSet)
		setupCategoryRouter(v1, handlerSet)
	}
}

func setupUserRouter(routerGroup *gin.RouterGroup, handlerSet *handlerset.HandlerSet, tokenManager *auth.TokenManager, logger *log.Logger) {
	openGroup := routerGroup.Group("/user")
	{
		openGroup.POST("/register", handlerSet.UserHandler.CreateUser)
		openGroup.POST("/login", handlerSet.UserHandler.LoginUser)
		openGroup.POST("/check", handlerSet.UserHandler.CheckUser)
	}
	authGroup := routerGroup.Group("/user-auth", middleware.AuthMiddleware(tokenManager, logger))
	{
		authGroup.POST("/verify", handlerSet.UserHandler.VerifyUser)
	}
}

func setupEncryptionRouter(routerGroup *gin.RouterGroup, handlerSet *handlerset.HandlerSet) {
	openGroup := routerGroup.Group("/encryption")
	{
		openGroup.GET("/public-key", handlerSet.EncryptionHandler.GetPublicKey)
	}
}

func setupCompetitionRouter(routerGroup *gin.RouterGroup, handlerSet *handlerset.HandlerSet) {
	adminGroup := routerGroup.Group("/competition")
	{
		adminGroup.POST("/create", handlerSet.CompetitionHandlerAdmin.CreateCompetition)
	}
	openGroup := routerGroup.Group("/competition")
	{
		openGroup.GET("/", handlerSet.CompetitionHandlerFrontend.GetCompetition)
	}
}

func setupCategoryRouter(routerGroup *gin.RouterGroup, handlerSet *handlerset.HandlerSet) {
	adminGroup := routerGroup.Group("/category")
	{
		adminGroup.POST("/create", handlerSet.CategoryHandler.CreateCategory)
		adminGroup.GET("/", handlerSet.CategoryHandler.GetAllCategory)
	}
}
