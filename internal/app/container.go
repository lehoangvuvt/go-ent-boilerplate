package app

import (
	"context"
	"fmt"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap"
	bootstrapstack "github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap/stack"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
	rediscache "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/cache/redis"
	entdb "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/ent"
	jwtinfra "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/jwt"
	transactionrepository "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/repository/transaction"
	userrepository "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/repository/user"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	Router *chi.Mux
	DB     *entdb.Client
}

func Build(ctx context.Context, cfg *config.Config) (*Container, error) {
	entDB, err := bootstrap.BootstrapEntDB(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("bootstrapping ent DB: %w", err)
	}

	userRepo := userrepository.NewUserRepository(entDB.Client())
	transactionRepo := transactionrepository.NewTransactionRepository(entDB.Client())

	jwtDuration := time.Duration(cfg.JWTConfig.Duration) * time.Second
	jwtService := jwtinfra.NewService(cfg.JWTConfig.Secret, jwtDuration)

	cacheService := rediscache.NewRedisCache(&redis.Options{
		Addr:     cfg.RedisConfig.Address,
		Password: cfg.RedisConfig.Password,
	})

	err = cacheService.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("pinging Redis: %w", err)
	}

	idempotencyStore := bootstrapstack.BuildIdempotencyStore(
		bootstrapstack.BuildIdempotencyStoreArgs{
			RedisAddr: cfg.RedisConfig.Address,
			Password:  cfg.RedisConfig.Password,
			TTL:       10 * time.Minute,
		},
	)

	router := bootstrap.BootstrapHandler(bootstrap.HandlerBootstrapArgs{
		Repositories: bootstrap.Repositories{
			UserRepository:        userRepo,
			TransactionRepository: transactionRepo,
		},
		Services: bootstrap.Services{
			JWTService:   jwtService,
			JWTDuration:  jwtDuration,
			CacheService: cacheService,
		},
		Stores: bootstrap.Stores{
			IdempotencyStore: idempotencyStore,
		},
	})

	return &Container{
		Router: router,
		DB:     entDB,
	}, nil
}

func (c *Container) Close() error {
	if c == nil || c.DB == nil {
		return nil
	}
	return c.DB.Close()
}
