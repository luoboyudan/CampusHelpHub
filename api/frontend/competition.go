package frontend

import (
	"campushelphub/api/common"
	"campushelphub/internal/common/converter"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"campushelphub/internal/service"
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
	id := c.Param("id")
	if id == "" {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(c, logInfo, h.Error.NewError(errors.ErrBadRequest, http.StatusBadRequest, nil))
		return
	}
	competitionID, err := converter.ToUint64(id)
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(c, logInfo, h.Error.NewError(errors.ErrBadRequest, http.StatusBadRequest, err))
		return
	}
	competition, err := h.CompetitionService.GetCompetition(c, competitionID)
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(c, logInfo, h.Error.NewError(errors.ErrBadRequest, http.StatusBadRequest, err))
		return
	}
	logInfo.Status = common.SuccessStatus
	h.SuccessResponse(c, logInfo, competition)
}
