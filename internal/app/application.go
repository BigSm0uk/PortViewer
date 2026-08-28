package app

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/bigsm0uk/port-viewer/internal/app/config"
	"github.com/bigsm0uk/port-viewer/internal/app/logger"
	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

type Application struct {
	logger *zap.Logger
	cfg    *config.AppConfig

	httpServer *echo.Echo
}

func New() *Application {
	return &Application{}
}

func (a *Application) Config() *config.AppConfig {
	if a.cfg == nil {
		cfg, err := config.ReadConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		a.cfg = cfg
	}
	return a.cfg
}
func (a *Application) Logger() *zap.Logger {
	if a.logger == nil {
		logger, err := logger.NewLogger(a.Config().Logging.Level, a.Config().Logging.Format)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		a.logger = logger
	}
	return a.logger
}
func (a *Application) HttpServer() *echo.Echo {
	if a.httpServer == nil {
		a.httpServer = echo.New()
		a.Routes(a.httpServer)
	}
	return a.httpServer
}
func (a *Application) Run(ctx context.Context) error {
	sc := echo.StartConfig{Address: net.JoinHostPort(a.Config().Server.Host, a.Config().Server.Port), HideBanner: true}

	if err := sc.Start(ctx, a.HttpServer()); err != nil {
		return err
	}
	return nil
}
