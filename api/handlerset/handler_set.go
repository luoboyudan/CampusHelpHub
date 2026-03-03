package handlerset

import (
	"campushelphub/api/frontend"
)

type HandlerSet struct {
	UserHandler       *frontend.UserHandler
	EncryptionHandler *frontend.EncryptionHandler
}

func NewHandlerSet(userHandler *frontend.UserHandler, encryptionHandler *frontend.EncryptionHandler) *HandlerSet {
	return &HandlerSet{
		UserHandler:       userHandler,
		EncryptionHandler: encryptionHandler,
	}
}
