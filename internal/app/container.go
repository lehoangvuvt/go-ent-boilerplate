package app

import (
	"github.com/go-chi/chi/v5"
	entdb "github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/ent"
)

type Container struct {
	Router *chi.Mux
	DB     *entdb.Client
}

func (c *Container) Close() error {
	if c == nil || c.DB == nil {
		return nil
	}
	return c.DB.Close()
}
