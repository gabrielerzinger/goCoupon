package repositories

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// ConnectRedis connects to redis db
func (s *Storage) ConnectRedis(config *viper.Viper) error {
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
