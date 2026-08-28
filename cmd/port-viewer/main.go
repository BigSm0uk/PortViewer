package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bigsm0uk/port-viewer/internal/app"
	"go.uber.org/zap"
)

func main() {
	a := app.New()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		err := a.Run(ctx)
		if err != nil {
			a.Logger().Info("failed to run", zap.Error(err))
			os.Exit(1)
		}
	}()
	<-ctx.Done()
	a.Logger().Info("server stopped")
	os.Exit(0)
}
