package app

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

type App struct {
  Address string
  Logger logrus.FieldLogger
  Config *viper.Viper
  Router *mux.Router
  Server *http.Server
}

func NewApp(host string, port int, config *viper.Viper,) (*App, error) {
    a := &App{
      Address: fmt.Sprintf("%s:%d", host, port),
      Config: config,
    }
    a.configureApp()
    return a, nil
}

func (a *App) configureApp() {
  a.Router = a.getRouter()
  a.configureServer()
}

func (a *App) configureServer() {
	a.Server = &http.Server{
		Addr:    a.Address,
		Handler: a.Router,
	}
}

func (a *App) getRouter() *mux.Router {
  router := mux.NewRouter()
  router.Handle("/healthcheck", NewHealthcheckHandler(a)).Methods("GET")
  return router
}

func (a *App) Init() {
  a.Server.ListenAndServe()
}
