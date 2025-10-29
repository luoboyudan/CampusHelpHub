package frontend

import (
	"campushelphub/api/common"
	"campushelphub/internal/service"
	"campushelphub/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*common.Handler
	UserService   *service.UserService
	WechatService *service.WechatService
}

func NewUserHandler(h *common.Handler, us *service.UserService, ws *service.WechatService) *UserHandler {
	return &UserHandler{
		Handler:       h,
		UserService:   us,
		WechatService: ws,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}
	sessionResp, err := h.WechatService.Login(req.Code)
	if err != nil {
		h.ErrorResponse(ctx, err.GetHTTPStatus(), err.Msg, err.Detail)
		return
	}
	if err = h.UserService.Create(ctx, &req, sessionResp); err != nil {
		h.ErrorResponse(ctx, err.GetHTTPStatus(), err.Msg, err.Detail)
		return
	}
}
