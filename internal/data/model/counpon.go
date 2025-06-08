package model

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type CouponModel struct {
	ID              int64
	Code            string
	DiscountPercent float64
	ValidUntil      time.Time
	CreatedAt       time.Time
}

func ToCouponModel(e entity.Coupon) CouponModel {
	return CouponModel{
		Code:            e.Code,
		DiscountPercent: e.DiscountPercent,
		ValidUntil:      e.ValidUntil,
		CreatedAt:       time.Now(),
	}
}
