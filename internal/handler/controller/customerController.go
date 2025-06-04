package controller

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerController struct {
	svc contract.ServiceManager
}

func NewCustomerController(svc contract.ServiceManager) *CustomerController {
	return &CustomerController{
		svc: svc,
	}
}

func (s *CustomerController) SaveNewCustomer(ctx *gin.Context) {
	var request request.NewCustomerRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	validate := validator.New()
	validate.RegisterValidation("cpf_cnpj", util.CpfCnpjValidator)
	response, err := s.svc.InternalService().CustomerService().Save(ctx, request)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save customer")
		return
	}
	util.ResponderApiOk(ctx, response)

}
