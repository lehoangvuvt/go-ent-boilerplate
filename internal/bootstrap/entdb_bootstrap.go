package bootstrap

import (
	"context"
	"fmt"

	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
	entdb "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/ent"
)

type BuildDSNParams struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     int
}

func BootstrapEntDB(ctx context.Context, cfg *config.Config) (*entdb.Client, error) {
	dsn := buildDSN(BuildDSNParams{
		User:     cfg.DBConfig.User,
		Password: cfg.DBConfig.Password,
		Name:     cfg.DBConfig.Name,
		Host:     cfg.DBConfig.Host,
		Port:     cfg.DBConfig.Port,
	})

	entDB, err := entdb.New(ctx, entdb.Config{
		Driver:      "postgres",
		DSN:         dsn,
		AutoMigrate: cfg.DBConfig.AutoMigrate,
	})
	if err != nil {
		return nil, fmt.Errorf("bootstrapping ent db: %w", err)
	}

	return entDB, nil
}

func buildDSN(params BuildDSNParams) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.Name)
}
