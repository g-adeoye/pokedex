package repository

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(host string) *RedisClient {
	var client *redis.Client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})

	return &RedisClient{client: client}
}
