//go:build wireinject
// +build wireinject

package setup

import (
	"campushelphub/api/admin"
	handlerCommon "campushelphub/api/common"
	"campushelphub/api/frontend"
	"campushelphub/api/handlerset"
	RSA "campushelphub/internal/common/RSA"
	"campushelphub/internal/common/auth"
	"campushelphub/internal/common/cache"
	"campushelphub/internal/common/snowflake"
	"campushelphub/internal/config"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"campushelphub/internal/pkg/redis"
	"campushelphub/internal/pkg/wechat"
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
	repository.NewMySQLCompetitionRepository,
	repository.NewMySQLCategoryRepository,
	repository.NewMySQLReminderRepository,
	repository.NewMySQLMessageRepository,
)

var CacheSet = wire.NewSet(
	cache.NewKeyBuilder,
)

var ServiceSet = wire.NewSet(
	service.NewUserService,
	service.NewChromeService,
	wechat.NewWechatService,
	redis.NewRedisService,
	service.NewCompetitionService,
	service.NewCategoryService,
	service.NewTaskService,
	service.NewReminderService,
	service.NewMessageService,
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
	frontend.NewEncryptionHandler,
	frontend.NewCompetitionHandler,
)

var AdminHandlerSet = wire.NewSet(
	admin.NewCompetitionHandler,
	admin.NewCategoryHandler,
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
		AdminHandlerSet,
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
