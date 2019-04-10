package repositories

import (
	"github.com/gabrielerzinger/goCoupon/models"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// Storage struct
type Storage struct {
	Redis *redis.Client
}

// NewStorage ctor
func NewStorage() *Storage {
	return &Storage{}
}

// Connect connects to redis db
func (s *Storage) Connect(config *viper.Viper) error {
	client := redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.url"),
		Password: config.GetString("redis.password"),
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic("Failed to connect to redis.")
	}

	s.Redis = client

	return err
}

// Find impl
func (s *Storage) Find(name string) (*models.Coupon, error) { return nil, nil }

// Update impl
func (s *Storage) Update(name string, coupon *models.Coupon) error { return nil }

// Store impl
func (s *Storage) Store(name string, coupon *models.Coupon) error { return nil }

// Ping impl
func (s *Storage) Ping() error { return nil }
