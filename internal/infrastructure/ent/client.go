package entdb

import (
	"context"
	"fmt"

	"github.com/lehoangvuvt/go-ent-boilerplate/ent"
)

type Config struct {
	Driver      string
	DSN         string
	AutoMigrate bool
}

type Client struct {
	Ent *ent.Client
}

func New(ctx context.Context, cfg Config) (*Client, error) {
	client, err := ent.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed opening ent client. Error: %w", err)
	}

	if cfg.AutoMigrate {
		if err := client.Schema.Create(ctx); err != nil {
			client.Close()
			return nil, fmt.Errorf("failed running migrations. Error: %w", err)
		}
	}

	return &Client{Ent: client}, err
}

func (c *Client) Close() error {
	return c.Close()
}

func (c *Client) Client() *ent.Client {
	return c.Ent
}
