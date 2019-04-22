package usecases

import (
	"time"

	"github.com/gabrielerzinger/goCoupon/repositories"

	"github.com/gabrielerzinger/goCoupon/models"
)

// Coupon usecases interface
type Coupon interface {
	CreateCoupon(string, string, float64, float64, time.Time) error
	FindCoupon(string) (*models.Coupon, error)
	RemoveCoupon(string) error
	UpdateCoupon(string, *models.Coupon) error
}

type coupon struct {
	repo repositories.Repository
}

// NewUsecase ctor
func NewUsecase(newRepo repositories.Repository) *coupon {
	return &coupon{repo: newRepo}
}

// CreateCoupon creates a new Coupon
func (c coupon) CreateCoupon(name, discountType string, amount, cartPrice float64, eTime time.Time) error {
	newCoupon := &models.Coupon{DiscountType: discountType, Amount: amount, CartPrice: cartPrice, TimesUsed: 0, ExpirationTime: eTime}
	return c.repo.Store(name, newCoupon)
}

// FindCoupon finds a coupon given its name
func (c coupon) FindCoupon(name string) (*models.Coupon, error) {
	foundCoupon, err := c.repo.Find(name)

	if err != nil {
		return &models.Coupon{}, err
	}

	return foundCoupon, err
}

// RemoveCoupon removes a coupon from the storage
func (c coupon) RemoveCoupon(name string) error {
	return c.repo.Remove(name)
}

// UpdateCoupon updates an already stored coupon
func (c coupon) UpdateCoupon(name string, coupon *models.Coupon) error {
	return c.repo.Update(name, coupon)
}
