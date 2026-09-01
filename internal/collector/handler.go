package collector

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

type Handler struct {
	Logger    *zap.Logger
	Collector *Collector
}

func NewHandler(logger *zap.Logger, collector *Collector) *Handler {
	return &Handler{
		Logger:    logger,
		Collector: collector,
	}
}

func (h *Handler) Listeners(c *echo.Context) error {
	listeners, err := h.Collector.Listeners()
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, listeners)
}
