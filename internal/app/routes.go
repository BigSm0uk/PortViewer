package app

import "net/http"

func (a *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", a.handlePing)

	return mux
}

func (a *Application) handlePing(w http.ResponseWriter, r *http.Request) {
	a.Logger().Debug("ping request")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("pong"))
}
