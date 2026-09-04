package app

import (
	"net/http"

	"github.com/bigsm0uk/port-viewer/internal/webui"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func (a *Application) Routes() {
	e := a.httpServer

	e.Use(middleware.Recover())

	api := e.Group("/api")
	api.Use(middleware.RequestLogger())

	api.GET("/ping", a.handlePing)
	api.GET("/listeners", a.CollectorHandler().Listeners)
	e.StaticFS("/", echo.MustSubFS(webui.EmbedFS, "dist"))
}

func (a *Application) handlePing(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
