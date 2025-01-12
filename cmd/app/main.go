package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	app2 "gtihub.com/blckvia/go-queue/internal/app"
)

func main() {
	ctx := context.Background()

	logger := zap.Must(zap.NewProduction())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("failed to sync logger", zap.Error(err))
		}
	}(logger)

	app := app2.NewApp(ctx, logger)

	go func() {
		if err := app.Run(); err != nil {
			app.Logger.Fatal("failed to run server: %w", zap.Error(err))
		}
	}()

	app.Logger.Info("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	app.Logger.Info("app shutting down")

	if err := app.Shutdown(ctx, logger); err != nil {
		app.Logger.Error("failed to shutdown server: %w", zap.Error(err))
	}
}
