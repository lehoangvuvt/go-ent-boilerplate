package app

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
	entdb "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/ent"
	userrepository "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/repository/user"
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

	router := bootstrap.BootstrapHandler(bootstrap.HandlerBootstrapArgs{
		Repositories: bootstrap.Repositories{
			UserRepository: userRepo,
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
