package controller

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
)

type StateController struct {
	svc contract.ServiceManager
}

func NewStateController(svc contract.ServiceManager) *StateController {
	return &StateController{
		svc: svc,
	}
}

func (s *StateController) SaveNewState(ctx *gin.Context) {
	var request request.NewStateRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	countryId, err := util.GetAndValidateIntParam(ctx, "countryId", "country Id invalid", false)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Validation error")
	}
	response, err := s.svc.InternalService().StateService().Save(ctx, request, countryId)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save state")
		return
	}
	util.ResponderApiOk(ctx, response)

}
