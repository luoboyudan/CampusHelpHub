package frontend

import (
	"campushelphub/api/common"
	"campushelphub/internal/common/auth"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"campushelphub/internal/service"
	"campushelphub/model"
	"net/http"

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

	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeUserRegister,
		ClientIP:     ctx.ClientIP(),
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrUserRegisterRequest, http.StatusBadRequest, err))
		return
	}

	sessionResp, err := h.WechatService.Login(req.Code)
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, err)
		return
	}

	user, err := h.UserService.Create(ctx, &req, sessionResp)
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, err)
		return
	}
	logInfo.ID = user.ID

	token, err := h.TokenManager.GenerateToken(int64(user.ID))
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, err)
		return
	}

	logInfo.Status = common.SuccessStatus
	h.SuccessResponse(ctx, logInfo, model.CreateUserResponse{
		Token: token,
	})
}
