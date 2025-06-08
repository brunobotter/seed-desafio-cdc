package entity

import "time"

type Coupon struct {
	Code            string
	DiscountPercent float64
	ValidUntil      time.Time
}
