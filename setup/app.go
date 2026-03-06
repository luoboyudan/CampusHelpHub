package setup

import (
	"campushelphub/internal/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sign := <-signChan
		a.stop()
		fmt.Println("程序退出,信号:", sign)
		os.Exit(0)
	}()
	a.Engine.Engine.Run(a.Config.Server.Addr)
}

func (a *App) stop() {
	a.Engine.ChromeService.Stop()
}
