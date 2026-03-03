package common

import (
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Error *errors.Error
	Log   *log.Logger
}

func NewHandler(error *errors.Error, log *log.Logger) *Handler {
	return &Handler{Error: error, Log: log}
}

func (h *Handler) ErrorResponse(ctx *gin.Context, logInfo *log.BusinessLogInfo, err *errors.Error) {
	ctx.JSON(err.HTTPStatus, gin.H{
		"error":  err.Error(),
		"detail": err.Detail,
	})
	logInfo.Extra = map[string]interface{}{
		"error": err.Error(),
	}
	h.Log.Info(logInfo)
}

func (h *Handler) SuccessResponse(ctx *gin.Context, logInfo *log.BusinessLogInfo, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
	h.Log.Info(logInfo)
}
