package repositories

import (
	"strconv"
	"time"

	"github.com/gabrielerzinger/goCoupon/models"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// RedisStorage struct
type RedisStorage struct {
	Redis *redis.Client
}

// NewRedisStorage ctor
func NewRedisStorage() *RedisStorage {
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

// Find a coupon given its name
func (s *RedisStorage) Find(name string) (*models.Coupon, error) {
	couponMap, err := s.Redis.HGetAll(name).Result()

	if err != nil {
		return nil, err
	}

	amount, _ := strconv.ParseFloat(couponMap["amount"], 64)
	discountType := couponMap["type"]
	cartPrice, _ := strconv.ParseFloat(couponMap["cartPrice"], 64)
	used, _ := strconv.ParseInt(couponMap["used"], 10, 64)
	expiration, _ := time.Parse(time.RFC3339, couponMap["expiration"])

	coupon := &models.Coupon{
		Amount:         amount,
		DiscountType:   discountType,
		CartPrice:      cartPrice,
		ExpirationTime: expiration,
		TimesUsed:      used,
	}
	return coupon, nil
}

// Update impl
func (s *RedisStorage) Update(name string, coupon *models.Coupon) error {
	return s.Store(name, coupon)
}

// Store saves given coupon to Redis
func (s *RedisStorage) Store(name string, coupon *models.Coupon) error {

	expirationTime, err := coupon.ExpirationTime.MarshalText()

	if err != nil {
		return err
	}

	couponMap := map[string]interface{}{
		"amount":     coupon.Amount,
		"type":       coupon.DiscountType,
		"cartPrice":  coupon.CartPrice,
		"used":       coupon.TimesUsed,
		"expiration": expirationTime,
	}

	_, err = s.Redis.HMSet(name, couponMap).Result()

	if err != nil {
		return err
	}

	_ = s.Redis.PExpireAt(name, coupon.ExpirationTime)

	return err
}

// Remove a coupon from the storage
func (s *RedisStorage) Remove(name string) error {
	_, err := s.Redis.HDel(name).Result()

	return err
}

// Ping implementation
func (s *RedisStorage) Ping() error {
	_, err := s.Redis.Ping().Result()

	return err
}
