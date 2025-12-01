package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/lehoangvuvt/go-ent-boilerplate/internal/app"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
	"github.com/lehoangvuvt/go-ent-boilerplate/pkg/logger"
)

func main() {
	slog := logger.GetSugaredLogger()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		slog.Fatalf("loading config: %v", err)
	}

	runner, closer, err := app.InitializeWorker(ctx, cfg)
	if err != nil {
		slog.Fatalf("worker init error: %v", err)
	}
	defer closer()

	slog.Info("Worker started ...")

	if err := runner(ctx); err != nil {
		slog.Fatalf("worker run error: %v", err)
	}
}
