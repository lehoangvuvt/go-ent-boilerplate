package bootstrapstack

import (
	"time"

	idempotencystore "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/idempotency"
	idempotencyports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/idempotency"
	"github.com/redis/go-redis/v9"
)

type BuildIdempotencyStoreArgs struct {
	RedisAddr string
	Password  string
	TTL       time.Duration
}

func BuildIdempotencyStore(args BuildIdempotencyStoreArgs) idempotencyports.IdempotencyStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     args.RedisAddr,
		Password: args.Password,
	})

	store := idempotencystore.NewRedisIdempotencyStore(rdb, args.TTL)
	return store
}
