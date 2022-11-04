package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-chat/internal/handler/rest"
	"go-chat/internal/service"
	"go-chat/pkg/logging"
	"net/http"
	"time"
)

type Server interface {
	Start() error
}

type server struct {
	logger     *logging.Logger
	config     *Config
	service    *service.Service
	httpServer *http.Server
}

func New(config *Config, service *service.Service, logger *logging.Logger) Server {
	return &server{
		logger:  logger,
		service: service,
		config:  config,
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
			Handler:      rest.New(httprouter.New(), service, logger),
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
	}
}

func (s *server) Start() error {
	s.logger.Info("Starting server...")
	return s.httpServer.ListenAndServe()
}
