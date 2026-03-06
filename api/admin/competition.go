package admin

import (
	"campushelphub/api/common"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"campushelphub/internal/service"
	"campushelphub/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompetitionHandler struct {
	*common.Handler
	Error   *errors.Error
	Logger  *log.Logger
	Service *service.CompetitionService
}

func NewCompetitionHandler(h *common.Handler, e *errors.Error, l *log.Logger, s *service.CompetitionService) *CompetitionHandler {
	return &CompetitionHandler{
		Handler: h,
		Error:   e,
		Logger:  l,
		Service: s,
	}
}

func (h *CompetitionHandler) CreateCompetition(ctx *gin.Context) {
	var req model.CreateCompetitionRequest
	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeCreateCompetition,
		ClientIP:     ctx.ClientIP(),
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrCreateCompetitionRequest, http.StatusBadRequest, err))
		return
	}
	if err := h.Service.CreateCompetition(ctx, &req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, logInfo, h.Error.NewError(errors.ErrCreateCompetitionDB, http.StatusInternalServerError, err))
		return
	}
	logInfo.Status = common.SuccessStatus
	h.SuccessResponse(ctx, logInfo, model.CreateCompetitionResponse{
		Result: true,
	})
}
