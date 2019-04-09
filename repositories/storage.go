package repositories

import (
	"github.com/go-redis/redis"
)

type Storage struct {
	Redis *redis.Client
}

func NewStorage() *Storage {
	return &Storage{}
}