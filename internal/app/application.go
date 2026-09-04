package app

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/bigsm0uk/port-viewer/internal/app/config"
	"github.com/bigsm0uk/port-viewer/internal/app/logger"
	"github.com/bigsm0uk/port-viewer/internal/collector"
	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

type Application struct {
	logger     *zap.Logger
	cfg        *config.AppConfig
	httpServer *echo.Echo

	collector        *collector.Collector
	collectorHandler *collector.Handler
}

func New() *Application {
	return &Application{}
}

// Init Инициализирует приложение, решая проблему гонки данных
func (a *Application) Init() *Application {
	a.Collector()
	a.HttpServer()
	return a
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

func (a *Application) CollectorHandler() *collector.Handler {
	if a.collectorHandler == nil {
		a.collectorHandler = collector.NewHandler(a.Logger(), a.Collector())
	}
	return a.collectorHandler
}

func (a *Application) Collector() *collector.Collector {
	if a.collector == nil {
		a.collector = collector.New(a.Logger())
	}
	return a.collector
}

func (a *Application) HttpServer() *echo.Echo {
	if a.httpServer == nil {
		a.httpServer = echo.New()
		a.Routes()
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
