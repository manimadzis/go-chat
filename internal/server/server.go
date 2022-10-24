package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chat/internal/handler"
	"go-chat/pkg/logging"
	"net/http"
)

type Server interface {
	Start() error
}

type server struct {
	router *gin.Engine
	logger *logging.Logger
	config *Config
}

func New(config *Config, logger *logging.Logger) Server {

	s := &server{
		router: gin.New(),
		logger: logger,
		config: config,
	}
	handler.Handler{}.InitRouter(s.router)

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Start() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.config.Host, s.config.Post), s)
}
