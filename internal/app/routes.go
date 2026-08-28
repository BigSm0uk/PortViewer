package app

import (
	"net/http"

	"github.com/bigsm0uk/port-viewer/internal/webui"
	"github.com/labstack/echo/v5"
)

func (a *Application) Routes(e *echo.Echo) {
	e.StaticFS("/", echo.MustSubFS(webui.EmbedFS, "dist"))
	e.GET("/ping", a.handlePing)
}

func (a *Application) handlePing(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
