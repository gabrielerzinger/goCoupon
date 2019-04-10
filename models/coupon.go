package models

/*
	DiscountType must be either 'PERCENT', 'FIXED_VALUE'
	Amount must be the discount amount [0,1] for PERCENt
	CartPrice must be the minimum cart price which the coupon aplies
*/

// Coupon model
type Coupon struct {
	DiscountType string  `json:"type"`
	Amount       float32 `json:"amount"`
	CartPrice    float32 `json:"cartPrice"`
	TimesUsed    int32   `json:"timesUsed"`
}
