package controller

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
)

type CountryController struct {
	svc contract.ServiceManager
}

func NewCountryController(svc contract.ServiceManager) *CountryController {
	return &CountryController{
		svc: svc,
	}
}

func (s *CountryController) SaveNewCountry(ctx *gin.Context) {
	var request request.NewCountryRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	response, err := s.svc.InternalService().CountryService().Save(ctx, request)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save countru")
		return
	}
	util.ResponderApiOk(ctx, response)

}
