package common

import (
	"net/http"

	"campushelphub/internal/errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Error *errors.Error
}

func NewHandler(error *errors.Error) *Handler {
	return &Handler{Error: error}
}

func (h *Handler) ErrorResponse(ctx *gin.Context, err *errors.Error) {
	ctx.JSON(err.HTTPStatus, gin.H{
		"error":  err.Error,
		"detail": err.Detail,
	})
}

func (h *Handler) SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
