package setup

import (
	"campushelphub/internal/config"

	"gorm.io/gorm"
)

type App struct {
	Engine *Engine
	Config *config.Config
	DB     *gorm.DB
}

func NewApp(engine *Engine, cfg *config.Config, db *gorm.DB) *App {
	return &App{
		Engine: engine,
		Config: cfg,
		DB:     db,
	}
}

func (a *App) Run() {
	a.Engine.Engine.Run(a.Config.Server.Addr)
}

func (a *App) Stop() {
	a.Engine.ChromeService.Stop()
}
