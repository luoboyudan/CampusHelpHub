//go:build wireinject
// +build wireinject

package setup

import (
	handlerCommon "campushelphub/api/common"
	"campushelphub/api/frontend"
	"campushelphub/api/handlerset"
	serviceCommon "campushelphub/internal/common"
	"campushelphub/internal/config"
	"campushelphub/internal/errors"
	"campushelphub/internal/repository"
	"campushelphub/internal/service"

	"github.com/google/wire"
)

var ErrorSet = wire.NewSet(
	errors.NewError,
)

var ConfigSet = wire.NewSet(
	config.NewConfig,
)

var DBSet = wire.NewSet(
	repository.NewDB,
)

var RepositorySet = wire.NewSet(
	repository.NewMySQLUserRepository,
)

var ServiceSet = wire.NewSet(
	service.NewUserService,
)

var BaseHandlerSet = wire.NewSet(
	handlerCommon.NewHandler,
)

var HandlerSet = wire.NewSet(
	handlerset.NewHandlerSet,
)
var WechatServiceSet = wire.NewSet(
	service.NewWechatService,
)

var FrontendHandlerSet = wire.NewSet(
	frontend.NewUserHandler,
)

var EngineSet = wire.NewSet(
	NewEngine,
)

var AppSet = wire.NewSet(
	NewApp,
)

var IDGenSet = wire.NewSet(
	serviceCommon.NewSnowflakeIDGenerator,
)

func InitializeApp() *App {
	wire.Build(
		ErrorSet,
		ConfigSet,
		DBSet,
		RepositorySet,
		ServiceSet,
		BaseHandlerSet,
		FrontendHandlerSet,
		HandlerSet,
		WechatServiceSet,
		EngineSet,
		AppSet,
		IDGenSet,
	)
	return nil
}
