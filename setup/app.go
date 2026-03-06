package setup

import (
	"campushelphub/internal/config"
	"campushelphub/internal/repository"
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
	// 测试用：删除数据库中所有表，上线时请注释掉
	repository.DropAllTables(a.DB)
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
