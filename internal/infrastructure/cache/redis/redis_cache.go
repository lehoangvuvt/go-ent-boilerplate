package rediscache

import (
	"context"
	"time"

	cacheports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/cache"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

var _ cacheports.Cache = (*RedisCache)(nil)

func NewRedisCache(opts *redis.Options) *RedisCache {
	client := redis.NewClient(opts)
	return &RedisCache{client: client}
}

func (c *RedisCache) Set(ctx context.Context, key string, value []byte) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func (c *RedisCache) SetTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return val, err
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *RedisCache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}
