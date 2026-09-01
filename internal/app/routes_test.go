package app

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

func TestRoutesRegistersListenersEndpoint(t *testing.T) {
	application := &Application{
		httpServer: echo.New(),
		logger:     zap.NewNop(),
	}
	application.Routes()

	for _, route := range application.httpServer.Router().Routes() {
		if route.Method == http.MethodGet && route.Path == "/api/listeners" {
			return
		}
	}

	t.Fatal("GET /api/listeners is not registered")
}
