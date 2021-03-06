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
	err := s.App.Storage.Ping()

	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to stablish connection", err)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Alive"))
}
