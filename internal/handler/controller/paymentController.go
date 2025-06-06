package controller

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	svc contract.ServiceManager
}

func NewPaymentController(svc contract.ServiceManager) *PaymentController {
	return &PaymentController{
		svc: svc,
	}
}

func (s *PaymentController) SaveNewPayment(ctx *gin.Context) {
	var request request.NewPaymentRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	customerId, err := util.GetAndValidateIntParam(ctx, "customerId", "Customer Id invalid", false)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Validation error")
	}
	response, err := s.svc.InternalService().PaymentService().Save(ctx, request, customerId)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save payment")
		return
	}
	util.ResponderApiOk(ctx, response)

}
