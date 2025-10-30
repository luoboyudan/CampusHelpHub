package setup

import (
	"campushelphub/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Engine *gin.Engine
	Config *config.Config
	DB     *gorm.DB
}

func NewApp(engine *gin.Engine, cfg *config.Config, db *gorm.DB) *App {
	return &App{
		Engine: engine,
		Config: cfg,
		DB:     db,
	}
}

func (a *App) Run() {
	a.Engine.Run(a.Config.Server.Addr)
}
