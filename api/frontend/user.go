package frontend

import (
	"campushelphub/api/common"
	RSA "campushelphub/internal/common/RSA"
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
	ChromeService *service.ChromeService
	RSA           *RSA.RSA
}

func NewUserHandler(h *common.Handler, us *service.UserService, ws *service.WechatService, tm *auth.TokenManager, cs *service.ChromeService, rs *RSA.RSA) *UserHandler {
	return &UserHandler{
		Handler:       h,
		UserService:   us,
		WechatService: ws,
		TokenManager:  tm,
		ChromeService: cs,
		RSA:           rs,
	}
}

func (h *UserHandler) CheckUser(ctx *gin.Context) {
	var req model.CheckUserRequest

	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeUserCheck,
		ClientIP:     ctx.ClientIP(),
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrUserCheckRequest, http.StatusBadRequest, err))
		return
	}
	sessionResp, err := h.WechatService.Login(req.Code)
	// 检查用户是否存在
	exist, err := h.UserService.CheckUser(ctx, sessionResp.OpenID)
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, err)
		return
	}
	h.SuccessResponse(ctx, logInfo, model.CheckUserResponse{
		Exist: exist,
	})
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

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	var req model.LoginUserRequest
	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeUserLogin,
		ClientIP:     ctx.ClientIP(),
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrUserLoginRequest, http.StatusBadRequest, err))
		return
	}
	// 登录用户
	sessionResp, err := h.WechatService.Login(req.Code)
	user, err := h.UserService.Login(ctx, sessionResp)
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
	h.SuccessResponse(ctx, logInfo, model.LoginUserResponse{
		Token:    token,
		Username: user.Username,
	})
}

func (h *UserHandler) VerifyUser(ctx *gin.Context) {
	var req model.VerifyUserRequest
	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeUserVerify,
		ClientIP:     ctx.ClientIP(),
	}
	userID, ok := ctx.Get("user_id")

	if !ok {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrUserVerifyRequest, http.StatusBadRequest, nil))
		return
	}
	req.UserID = uint64(userID.(float64))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrUserVerifyRequest, http.StatusBadRequest, err))
		return
	}
	//解密
	password, err := h.RSA.Decrypt(req.RSAPassword)
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrUserVerifyRequest, http.StatusBadRequest, err))
		return
	}
	// 验证用户
	if err := h.ChromeService.VerifyStudent(&model.ChromeStudentVerify{
		StudentID: req.StudentID,
		Password:  string(password),
	}); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, err)
		return
	}
	// 验证用户
	if err := h.UserService.Verify(ctx, &req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, err)
		return
	}
	logInfo.Status = common.SuccessStatus
	h.SuccessResponse(ctx, logInfo, model.VerifyUserResponse{
		Result: true,
	})
}
