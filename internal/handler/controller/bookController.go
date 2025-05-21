package controller

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	svc contract.ServiceManager
}

func NewBookController(svc contract.ServiceManager) *BookController {
	return &BookController{
		svc: svc,
	}
}

func (s *BookController) SaveNewBook(ctx *gin.Context) {
	var request request.NewBookRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, 400, err, "Invalid Data")
	}
	categoryId, err := util.GetAndValidateIntParam(ctx, "categoryId", "Category Id invalid", false)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Validation error")
	}
	authorId, err := util.GetAndValidateIntParam(ctx, "authorId", "Author Id invalid", false)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Validation error")
	}
	bookResponse, err := s.svc.InternalService().BookService().Save(ctx, request, categoryId, authorId)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save book")
	}
	util.ResponderApiOk(ctx, bookResponse)

}

func (s *BookController) GetById(ctx *gin.Context) {
	bookId, err := util.GetAndValidateIntParam(ctx, "bookId", "book Id invalid", false)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Validation error")
	}
	bookResponse, err := s.svc.InternalService().BookService().GetById(ctx, bookId)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save book")
	}
	util.ResponderApiOk(ctx, bookResponse)

}
