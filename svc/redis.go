package svc

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(options *redis.Options) *redis.Client {
	return redis.NewClient(options)
}
