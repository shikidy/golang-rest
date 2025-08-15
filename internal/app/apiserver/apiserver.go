package apiserver

import "github.com/sirupsen/logrus"

type ApiServer struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
	}
}

func (a *ApiServer) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}

	a.logger.Info("starting api server")
	return nil
}

func (a *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)

	return nil
}
