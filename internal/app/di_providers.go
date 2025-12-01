package app

import (
	"context"
	"time"

	"github.com/lehoangvuvt/go-ent-boilerplate/ent"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap"
	bootstrapstack "github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap/stack"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
	rediscache "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/cache/redis"
	entdb "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/ent"
	jwtinfra "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/jwt"
	smtpmail "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/mail/smtp"
	transactionrepository "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/repository/transaction"
	userrepository "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/repository/user"
	cacheports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/cache"
	idempotencyports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/idempotency"
	mailports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/mail"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	securityports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/security"
	"github.com/redis/go-redis/v9"
)

func provideEntDB(ctx context.Context, cfg *config.Config) (*entdb.Client, error) {
	return bootstrap.BootstrapEntDB(ctx, cfg)
}

func provideEntClient(db *entdb.Client) *ent.Client {
	return db.Client()
}

func provideUserRepository(client *ent.Client) repositoryports.UserRepository {
	return userrepository.NewUserRepository(client)
}

func provideTransactionRepository(client *ent.Client) repositoryports.TransactionRepository {
	return transactionrepository.NewTransactionRepository(client)
}

func provideJWTDuration(cfg *config.Config) time.Duration {
	return time.Duration(cfg.JWT.Duration) * time.Second
}

func provideJWTService(cfg *config.Config, duration time.Duration) securityports.JWTService {
	return jwtinfra.NewService(cfg.JWT.Secret, duration)
}

func provideRedisOptions(cfg *config.Config) *redis.Options {
	return &redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
	}
}

func provideCache(ctx context.Context, opts *redis.Options) (cacheports.Cache, error) {
	cache := rediscache.NewRedisCache(opts)
	if err := cache.Ping(ctx); err != nil {
		return nil, err
	}
	return cache, nil
}

func provideIdempotencyStore(cfg *config.Config) idempotencyports.IdempotencyStore {
	return bootstrapstack.BuildIdempotencyStore(
		bootstrapstack.BuildIdempotencyStoreArgs{
			RedisAddr: cfg.Redis.Address,
			Password:  cfg.Redis.Password,
			TTL:       10 * time.Minute,
		},
	)
}

func provideHandlerArgs(
	repos bootstrap.Repositories,
	services bootstrap.Services,
	stores bootstrap.Stores,
) bootstrap.HandlerBootstrapArgs {
	return bootstrap.HandlerBootstrapArgs{
		Repositories: repos,
		Services:     services,
		Stores:       stores,
	}
}

func provideMailService(cfg *config.Config) mailports.MailService {
	return smtpmail.NewSMTPMailService(
		cfg.Mail.Host,
		cfg.Mail.Port,
		cfg.Mail.User,
		cfg.Mail.Pass,
	)
}
