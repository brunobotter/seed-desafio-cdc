package response

import (
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
)

type CouponResponse struct {
	ID              int64     `json:"id"`
	Code            string    `json:"code"`
	DiscountPercent float64   `json:"discount_percent"`
	ValidUntil      time.Time `json:"valid_until"`
}

func FromCouponModel(m model.CouponModel) CouponResponse {
	return CouponResponse{
		ID:              m.ID,
		Code:            m.Code,
		DiscountPercent: m.DiscountPercent,
		ValidUntil:      m.ValidUntil,
	}
}
