package frontend

import (
	"campushelphub/api/common"
	"campushelphub/internal/service"
	"campushelphub/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*common.Handler
	UserService *service.UserService
}

func NewUserHandler(h *common.Handler, us *service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     h,
		UserService: us,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}
	if err := h.UserService.Create(ctx, &req); err != nil {
		return
	}
}
