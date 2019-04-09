package app

import (
  "net/http"
)

type HealthcheckHandler struct {
  App *App
}

func NewHealthcheckHandler(a *App) *HealthcheckHandler {
  m := &HealthcheckHandler{
    App: a,
  }
  return m
}

func(s *HealthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
  w.Write([]byte("Im alive!"))
}
