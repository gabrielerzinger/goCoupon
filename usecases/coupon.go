package usecases

import (
	"time"

	"github.com/gabrielerzinger/goCoupon/repositories"

	"github.com/gabrielerzinger/goCoupon/models"
)

// Coupon usecases interface
type Coupon interface {
	Create(string, string, float64, float64, time.Time) error
}

type coupon struct {
	repo repositories.Repository
}

// CreateCoupon creates a new Coupon
func (c coupon) CreateCoupon(name, discountType string, amount, cartPrice float64, eTime time.Time) error {
	newCoupon := &models.Coupon{DiscountType: discountType, Amount: amount, CartPrice: cartPrice, TimesUsed: 0, ExpirationTime: eTime}
	return c.repo.Store(name, newCoupon)
}
