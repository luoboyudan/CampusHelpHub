package setup

import (
	"campushelphub/api/handlerset"

	"github.com/gin-gonic/gin"
)

func NewEngine(handlerSet *handlerset.HandlerSet) *gin.Engine {
	gin := gin.Default()
	SetupRouter(gin, handlerSet)
	return gin
}

func SetupRouter(app *gin.Engine, handlerSet *handlerset.HandlerSet) {
	app.GET("/docs/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	v1 := app.Group("/v1")
	{
		v1.POST("/user", handlerSet.UserHandler.CreateUser)
	}
}
