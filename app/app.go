package app

import (
	"fmt"
	"net/http"

	"github.com/gabrielerzinger/goCoupon/repositories"
	"github.com/gabrielerzinger/goCoupon/usecases"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// App struct
type App struct {
	Address string
	Logger  logrus.FieldLogger
	Config  *viper.Viper
	Router  *mux.Router
	Server  *http.Server
	Storage repositories.Repository
	Usecase usecases.Coupon
}

// NewApp creates a new app
func NewApp(host string, port int, config *viper.Viper, log logrus.FieldLogger) (*App, error) {
	a := &App{
		Address: fmt.Sprintf("%s:%d", host, port),
		Config:  config,
		Logger:  log,
	}
	a.configureApp()
	return a, nil
}

func (a *App) configureApp() {
	a.Router = a.getRouter()
	a.Storage = repositories.NewRedisStorage()
	a.Usecase = usecases.NewUsecase(a.Storage)
	a.configureStorage()
	a.configureServer()
}

func (a *App) configureServer() {
	a.Server = &http.Server{
		Addr:    a.Address,
		Handler: a.Router,
	}
}

func (a *App) configureStorage() error {
	a.Logger.Info("Connecting to redis url:", a.Config.GetString("redis.url"))
	return a.Storage.Connect(a.Config)
}

func (a *App) getRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/healthcheck", NewHealthcheckHandler(a)).Methods("GET")
	router.Handle("/coupon", NewFindHandler(a)).Methods("GET")
	router.Handle("/coupon", NewSaveHandler(a)).Methods("POST")
	return router
}

// Init starts the app
func (a *App) Init() {
	a.Server.ListenAndServe()
}
