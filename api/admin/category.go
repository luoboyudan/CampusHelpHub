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

type CategoryHandler struct {
	*common.Handler
	Error   *errors.Error
	Logger  *log.Logger
	Service *service.CategoryService
}

func NewCategoryHandler(h *common.Handler, s *service.CategoryService, e *errors.Error, l *log.Logger) *CategoryHandler {
	return &CategoryHandler{
		Handler: h,
		Service: s,
		Error:   e,
		Logger:  l,
	}
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var req model.CreateCategoryRequest
	logInfo := log.BusinessLogInfo{
		BusinessType: common.BusinessTypeCreateCategory,
		ClientIP:     ctx.ClientIP(),
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, &logInfo, h.Error.NewError(errors.ErrCreateCategoryRequest, http.StatusBadRequest, err))
		return
	}
	if err := h.Service.CreateCategory(ctx, &req); err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, &logInfo, h.Error.NewError(errors.ErrCreateCategoryDB, http.StatusInternalServerError, err))
		return
	}
	h.SuccessResponse(ctx, &logInfo, model.CreateCategoryResponse{
		Result: true,
	})
}

func (h *CategoryHandler) GetAllCategory(ctx *gin.Context) {
	logInfo := log.BusinessLogInfo{
		BusinessType: common.BusinessTypeGetAllCategory,
		ClientIP:     ctx.ClientIP(),
	}
	categories, err := h.Service.GetAllCategory(ctx)
	if err != nil {
		logInfo.Status = common.FailStatus
		h.ErrorResponse(ctx, &logInfo, h.Error.NewError(errors.ErrGetAllCategoryDB, http.StatusInternalServerError, err))
		return
	}
	logInfo.Status = common.SuccessStatus
	h.SuccessResponse(ctx, &logInfo, model.GetAllCategoryResponse{
		Categories: categories,
	})
}
