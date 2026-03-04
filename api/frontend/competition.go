package frontend

import (
	"campushelphub/api/common"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"campushelphub/internal/service"

	"github.com/gin-gonic/gin"
)

type CompetitionHandler struct {
	handler            *common.Handler
	competitionService *service.CompetitionService
	errorsError        *errors.Error
	logger             *log.Logger
}

func NewCompetitionHandler(handler *common.Handler, competitionService *service.CompetitionService, errorsError *errors.Error, logger *log.Logger) *CompetitionHandler {
	return &CompetitionHandler{
		handler:            handler,
		competitionService: competitionService,
		errorsError:        errorsError,
		logger:             logger,
	}
}

func (h *CompetitionHandler) GetCompetition(c *gin.Context) {
	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeUserCheck,
		ClientIP:     c.ClientIP(),
	}

	competitions, err := h.competitionService.GetCompetitions(c)
	if err != nil {
		logInfo.Status = common.FailStatus
		return
	}
	logInfo.Status = common.SuccessStatus
	h.handler.SuccessResponse(c, logInfo, competitions)
}
