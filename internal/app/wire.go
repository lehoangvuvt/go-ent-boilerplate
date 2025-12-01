//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"
	bootstrapstack "github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap/stack"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
)

//go:generate go run github.com/google/wire/cmd/wire

func InitializeContainer(
	ctx context.Context,
	cfg *config.Config,
) (*Container, error) {

	wire.Build(
		DBSet,
		RepositorySet,
		ServiceSet,
		StoreSet,
		HandlerSet,
		ContainerSet,
	)

	return nil, nil
}

func InitializeWorker(
	ctx context.Context,
	cfg *config.Config,
) (bootstrapstack.WorkerRunner, func(), error) {

	wire.Build(
		WorkerSet,
	)

	return nil, nil, nil
}
