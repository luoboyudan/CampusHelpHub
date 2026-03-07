package frontend

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
	CompetitionService *service.CompetitionService
	Error              *errors.Error
	Logger             *log.Logger
}

func NewCompetitionHandler(handler *common.Handler, competitionService *service.CompetitionService, errorsError *errors.Error, logger *log.Logger) *CompetitionHandler {
	return &CompetitionHandler{
		Handler:            handler,
		CompetitionService: competitionService,
		Error:              errorsError,
		Logger:             logger,
	}
}

func (h *CompetitionHandler) GetCompetitionList(c *gin.Context) {
	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeUserCheck,
		ClientIP:     c.ClientIP(),
	}

	competitions, err := h.CompetitionService.GetCompetitionsList(c)
	if err != nil {
		logInfo.Status = common.FailStatus
		return
	}
	logInfo.Status = common.SuccessStatus
	h.SuccessResponse(c, logInfo, competitions)
}

func (h *CompetitionHandler) GetCompetition(c *gin.Context) {
	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeUserCheck,
		ClientIP:     c.ClientIP(),
	}
	var req model.GetCompetitionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ErrorResponse(c, logInfo, h.Error.NewError(errors.ErrBadRequest, http.StatusBadRequest, err))
		return
	}
	competition, err := h.CompetitionService.GetCompetition(c, uint64(req.ID))
	if err != nil {
		logInfo.Status = common.FailStatus
		return
	}
	logInfo.Status = common.SuccessStatus
	h.SuccessResponse(c, logInfo, competition)
}
