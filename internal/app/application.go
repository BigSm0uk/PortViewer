package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/bigsm0uk/go-simple-service-template/internal/app/config"
	"github.com/bigsm0uk/go-simple-service-template/internal/app/logger"
	"go.uber.org/zap"
)

type Application struct {
	logger *zap.Logger
	cfg    *config.AppConfig

	httpServer *http.Server
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
func (a *Application) HttpServer(mux *http.ServeMux) *http.Server {
	if a.httpServer == nil {
		a.httpServer = &http.Server{
			Addr:    net.JoinHostPort(a.Config().Server.Host, a.Config().Server.Port),
			Handler: mux,
		}
	}
	return a.httpServer
}
func (a *Application) Run(ctx context.Context) error {
	a.Logger().Info("Server starting", zap.String("Addres", net.JoinHostPort(a.Config().Server.Host, a.Config().Server.Port)))
	if err := a.HttpServer(a.Routes()).ListenAndServe(); err != nil {
		return err
	}
	return nil
}
