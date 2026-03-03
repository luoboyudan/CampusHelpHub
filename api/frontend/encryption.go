package frontend

import (
	"campushelphub/api/common"
	RSA "campushelphub/internal/common/RSA"
	"campushelphub/internal/errors"
	"campushelphub/internal/log"
	"campushelphub/model"
	"net/http"

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
	var err error
	res.PublicKey, err = e.RSA.GetPublicKey()
	if err != nil {
		logInfo.Status = common.FailStatus
		e.ErrorResponse(ctx, logInfo, e.Error.NewError(errors.ErrGetPublicKey, http.StatusInternalServerError, err))
		return
	}
	logInfo.Status = common.SuccessStatus
	e.SuccessResponse(ctx, logInfo, res)
}
