package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/lehoangvuvt/go-ent-boilerplate/internal/app"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
	"github.com/lehoangvuvt/go-ent-boilerplate/pkg/logger"
	_ "github.com/lib/pq"
)

func main() {
	slog := logger.GetSugaredLogger()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	cfg, err := config.Load()
	if err != nil {
		slog.Fatalf("loading config: %v", err)
	}
	c, err := app.Build(ctx, cfg)
	if err != nil {
		slog.Fatalf("building app: %v", err)
	}
	defer c.Close()

	if err := app.StartServer(ctx, cfg, c.Router); err != nil {
		slog.Fatalf("starting server: %v", err)
	}
}
