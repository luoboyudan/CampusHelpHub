//go:build wireinject
// +build wireinject

package setup

import (
	handlerCommon "campushelphub/api/common"
	"campushelphub/api/frontend"
	"campushelphub/api/handlerset"
	RSA "campushelphub/internal/common/RSA"
	"campushelphub/internal/common/auth"
	"campushelphub/internal/common/snowflake"
	"campushelphub/internal/config"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
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

var TokenManagerSet = wire.NewSet(
	auth.NewTokenManager,
)

var RSASet = wire.NewSet(
	RSA.NewRSA,
)

var RepositorySet = wire.NewSet(
	repository.NewMySQLUserRepository,
)

var ServiceSet = wire.NewSet(
	service.NewUserService,
	service.NewChromeService,
	service.NewWechatService,
)

var BaseHandlerSet = wire.NewSet(
	handlerCommon.NewHandler,
)

var LoggerSet = wire.NewSet(
	log.NewLogger,
)

var HandlerSet = wire.NewSet(
	handlerset.NewHandlerSet,
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
	snowflake.NewSnowflakeIDGenerator,
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
		TokenManagerSet,
		RSASet,
		EngineSet,
		AppSet,
		IDGenSet,
		LoggerSet,
	)
	return nil
}
