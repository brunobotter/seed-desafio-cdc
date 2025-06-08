package request

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type NewCouponRequest struct {
	Code            string    `json:"code" validate:"required,min=3,max=50"`
	DiscountPercent float64   `json:"discount_percent" validate:"required,gt=0,lt=100"`
	ValidUntil      time.Time `json:"valid_until" validate:"required,gt"`
}

func (r NewCouponRequest) ToEntity() entity.Coupon {
	return entity.Coupon{
		Code:            r.Code,
		DiscountPercent: r.DiscountPercent,
		ValidUntil:      r.ValidUntil,
	}
}
