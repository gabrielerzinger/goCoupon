package app

import (
	"net/http"
)

// HealthcheckHandler struct
type HealthcheckHandler struct {
	App *App
}

// NewHealthcheckHandler ctor
func NewHealthcheckHandler(a *App) *HealthcheckHandler {
	m := &HealthcheckHandler{
		App: a,
	}
	return m
}

// ServeHTTP method
func (s *HealthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pong, err := s.App.Storage.Redis.Ping().Result()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to stablish connection"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(pong))
}
