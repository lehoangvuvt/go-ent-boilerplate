//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
)

//go:generate go run github.com/google/wire/cmd/wire

var containerSet = wire.NewSet(
	provideEntDB,
	provideEntClient,
	provideUserRepository,
	provideTransactionRepository,
	provideJWTDuration,
	provideJWTService,
	provideRedisOptions,
	provideCache,
	provideIdempotencyStore,
	provideHandlerArgs,
	wire.Struct(new(bootstrap.Repositories), "*"),
	wire.Struct(new(bootstrap.Services), "*"),
	wire.Struct(new(bootstrap.Stores), "*"),
	wire.Struct(new(Container), "*"),
	bootstrap.BootstrapHandler,
)

func InitializeContainer(ctx context.Context, cfg *config.Config) (*Container, error) {
	wire.Build(containerSet)
	return &Container{}, nil
}
