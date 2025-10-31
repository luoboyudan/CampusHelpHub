package frontend

import (
	"campushelphub/api/common"
	"campushelphub/internal/common/auth"
	"campushelphub/internal/service"
	"campushelphub/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*common.Handler
	UserService   *service.UserService
	WechatService *service.WechatService
	TokenManager  *auth.TokenManager
}

func NewUserHandler(h *common.Handler, us *service.UserService, ws *service.WechatService, tm *auth.TokenManager) *UserHandler {
	return &UserHandler{
		Handler:       h,
		UserService:   us,
		WechatService: ws,
		TokenManager:  tm,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}
	sessionResp, err := h.WechatService.Login(req.Code)
	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}
	user, err := h.UserService.Create(ctx, &req, sessionResp)
	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}
	token, err := h.TokenManager.GenerateToken(int64(user.ID))
	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}
	h.SuccessResponse(ctx, model.CreateUserResponse{
		Token: token,
	})
}
