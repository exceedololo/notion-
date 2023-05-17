package redis

import (
	"github.com/go-redis/redis"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) IncrementValue(key string, value int) int {
	// Increment the value in Redis
	newValue, _ := r.client.IncrBy(key, int64(value)).Result()
	return int(newValue)
}
