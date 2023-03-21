package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
}
