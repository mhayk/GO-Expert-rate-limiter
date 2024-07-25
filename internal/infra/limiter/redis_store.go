package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr, password string) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return &RedisStore{client: client}
}

func (rs *RedisStore) Get(key string) (string, error) {
	val, err := rs.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (rs *RedisStore) Set(key string, value interface{}, expiration time.Duration) error {
	return rs.client.Set(ctx, key, value, expiration).Err()
}

func (rs *RedisStore) Incr(key string) error {
	return rs.client.Incr(ctx, key).Err()
}

func (rs *RedisStore) Expire(key string, expiration time.Duration) error {
	return rs.client.Expire(ctx, key, expiration).Err()
}
