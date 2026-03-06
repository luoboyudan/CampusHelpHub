package handlerset

import (
	"campushelphub/api/admin"
	"campushelphub/api/frontend"
)

type HandlerSet struct {
	UserHandler                *frontend.UserHandler
	EncryptionHandler          *frontend.EncryptionHandler
	CompetitionHandlerAdmin    *admin.CompetitionHandler
	CompetitionHandlerFrontend *frontend.CompetitionHandler
	CategoryHandler            *admin.CategoryHandler
}

func NewHandlerSet(userHandler *frontend.UserHandler, encryptionHandler *frontend.EncryptionHandler, competitionHandler *admin.CompetitionHandler, competitionHandlerFrontend *frontend.CompetitionHandler, categoryHandler *admin.CategoryHandler) *HandlerSet {
	return &HandlerSet{
		UserHandler:                userHandler,
		EncryptionHandler:          encryptionHandler,
		CompetitionHandlerAdmin:    competitionHandler,
		CompetitionHandlerFrontend: competitionHandlerFrontend,
		CategoryHandler:            categoryHandler,
	}
}
