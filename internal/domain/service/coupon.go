package service

import (
	"context"

	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/response"
)

type couponService struct {
	svc contract.ServiceManager
}

func NewCouponService(svc contract.ServiceManager) contract.CouponService {
	return &couponService{
		svc: svc,
	}
}

func (s couponService) Save(ctx context.Context, request request.NewCouponRequest) (couponResponse response.CouponResponse, err error) {
	coupon := request.ToEntity()
	couponDb, err := s.svc.DB().CouponRepo().Save(ctx, coupon)
	if err != nil {
		return couponResponse, err
	}
	couponResponse = response.FromCouponModel(couponDb)
	return couponResponse, nil

}
