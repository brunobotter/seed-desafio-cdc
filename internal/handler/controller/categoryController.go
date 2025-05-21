package controller

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	svc contract.ServiceManager
}

func NewCategoryController(svc contract.ServiceManager) *CategoryController {
	return &CategoryController{
		svc: svc,
	}
}

func (s *CategoryController) SaveNewCategory(ctx *gin.Context) {
	var request request.NewCategoryRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	response, err := s.svc.InternalService().CategoryService().Save(ctx, request)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save category")
	}
	util.ResponderApiOk(ctx, response)

}
