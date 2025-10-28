package handlerset

import (
	"campushelphub/api/frontend"
)

type HandlerSet struct {
	UserHandler *frontend.UserHandler
}

func NewHandlerSet(userHandler *frontend.UserHandler) *HandlerSet {
	return &HandlerSet{
		UserHandler: userHandler,
	}
}
