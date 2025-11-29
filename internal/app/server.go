package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/config"
	"github.com/lehoangvuvt/go-ent-boilerplate/pkg/logger"
)

func StartServer(ctx context.Context, cfg *config.Config, r *chi.Mux) error {
	slog := logger.GetSugaredLogger()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: r,
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Infof("listening on :%d", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		slog.Info("shutting down server...")
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shutdown: %w", err)
		}
		return nil
	case err := <-errCh:
		return err
	}
}
