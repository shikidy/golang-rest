package apiserver

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *http.ServeMux
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: &http.ServeMux{},
	}
}

func (a *ApiServer) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}
	a.configureRouter()

	a.logger.Info("starting api server")
	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (a *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)

	return nil
}

func (a *ApiServer) configureRouter() {
	a.router.HandleFunc("/hello", a.handleHello())
}

func (a *ApiServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
