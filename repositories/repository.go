package repositories

import (
	"github.com/gabrielerzinger/goCoupon/models"
	"github.com/spf13/viper"
)

// Repository interface
type Repository interface {
	Connect(config *viper.Viper) error
	Find(name string) (*models.Coupon, error)
	Update(name string, coupon *models.Coupon) error
	Store(name string, coupon *models.Coupon) error
	Ping() error
}
