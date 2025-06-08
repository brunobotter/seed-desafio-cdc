package controller

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
)

type CouponController struct {
	svc contract.ServiceManager
}

func NewCouponController(svc contract.ServiceManager) *CouponController {
	return &CouponController{svc: svc}
}

func (s *CouponController) SaveNewCoupon(ctx *gin.Context) {
	var request request.NewCouponRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	response, err := s.svc.InternalService().CouponService().Save(ctx, request)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save coupon")
	}
	util.ResponderApiOk(ctx, response)

}
