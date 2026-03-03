package frontend

import (
	"campushelphub/api/common"
	RSA "campushelphub/internal/common/RSA"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"campushelphub/model"

	"github.com/gin-gonic/gin"
)

type EncryptionHandler struct {
	*common.Handler
	Error  *errors.Error
	Logger *log.Logger
	RSA    *RSA.RSA
}

func NewEncryptionHandler(h *common.Handler, e *errors.Error, l *log.Logger, rs *RSA.RSA) *EncryptionHandler {
	return &EncryptionHandler{
		Handler: h,
		Error:   e,
		Logger:  l,
		RSA:     rs,
	}
}

func (e *EncryptionHandler) GetPublicKey(ctx *gin.Context) {
	var res model.GetPublicKeyResponse
	logInfo := &log.BusinessLogInfo{
		BusinessType: common.BusinessTypeGetPublicKey,
		ClientIP:     ctx.ClientIP(),
	}
	res.PublicKey = e.RSA.GetPublicKey()
	e.SuccessResponse(ctx, logInfo, res)
}
