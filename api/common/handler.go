package common

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ErrorResponse(ctx *gin.Context, httpStatus int, msg, detail string) {
	ctx.JSON(httpStatus, gin.H{
		"error":  msg,
		"detail": detail,
	})
}
