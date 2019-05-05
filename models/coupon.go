package models

import "time"

/*
	DiscountType must be either 'PERCENT', 'FIXED_VALUE'
	Amount must be the discount amount [0,1] for PERCENt
	CartPrice must be the minimum cart price which the coupon aplies
*/

// Coupon model
type Coupon struct {
	DiscountType   string  `json:"type"`
	Amount         float64 `json:"amount"`
	CartPrice      float64 `json:"cartPrice"`
	TimesUsed      int64   `json:"timesUsed"`
	ExpirationTime time.Time
}

// CouponRequest for storage
type CouponRequest struct {
	Name           string  `json:"name"`
	DiscountType   string  `json:"type"`
	Amount         float64 `json:"amount"`
	CartPrice      float64 `json:"cartPrice"`
	ExpirationTime string  `json:"expirationTime"`
}
