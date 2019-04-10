package repositories

import (
	"github.com/gabrielerzinger/goCoupon/models"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// Storage struct
type RedisStorage struct {
	Redis *redis.Client
}

// NewStorage ctor
func NewStorage() *RedisStorage {
	return &RedisStorage{}
}

// Connect connects to redis db
func (s *RedisStorage) Connect(config *viper.Viper) error {
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
func (s *RedisStorage) Find(name string) (*models.Coupon, error) { return nil, nil }

// Update impl
func (s *RedisStorage) Update(name string, coupon *models.Coupon) error { return nil }

// Store impl
func (s *RedisStorage) Store(name string, coupon *models.Coupon) error { return nil }

// Ping implementation
func (s *RedisStorage) Ping() error {
	_, err := s.Redis.Ping().Result()

	return err
}
