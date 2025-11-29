package idempotencystore

import (
	"context"
	"encoding/json"
	"time"

	idempotencyports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/idempotency"
	"github.com/redis/go-redis/v9"
)

type RedisIdempotencyStore struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewRedisIdempotencyStore(rdb *redis.Client, ttl time.Duration) *RedisIdempotencyStore {
	return &RedisIdempotencyStore{
		rdb: rdb,
		ttl: ttl,
	}
}

func (s *RedisIdempotencyStore) Get(ctx context.Context, key string) (*idempotencyports.IdempotencyRecord, error) {
	val, err := s.rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var rec idempotencyports.IdempotencyRecord
	if err := json.Unmarshal(val, &rec); err != nil {
		return nil, err
	}

	return &rec, nil
}

func (s *RedisIdempotencyStore) SavePending(ctx context.Context, key string, requestHash string) error {
	rec := idempotencyports.IdempotencyRecord{
		Key:         key,
		RequestHash: requestHash,
		Status:      "pending",
	}

	data, _ := json.Marshal(rec)
	return s.rdb.Set(ctx, key, data, s.ttl).Err()
}

func (s *RedisIdempotencyStore) SaveSuccess(ctx context.Context, key string, response []byte) error {
	rec := idempotencyports.IdempotencyRecord{
		Key:      key,
		Response: response,
		Status:   "success",
	}

	data, _ := json.Marshal(rec)
	return s.rdb.Set(ctx, key, data, s.ttl).Err()
}

func (s *RedisIdempotencyStore) SaveFailed(ctx context.Context, key string) error {
	rec := idempotencyports.IdempotencyRecord{
		Key:    key,
		Status: "failed",
	}

	data, _ := json.Marshal(rec)
	return s.rdb.Set(ctx, key, data, s.ttl).Err()
}
