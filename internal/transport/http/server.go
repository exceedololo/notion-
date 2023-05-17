package http

import (
	"log"
	"net/http"

	"github.com/exceedololo/notion-/internal/config"
	"github.com/exceedololo/notion-/internal/transport/http/handlers"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	handlers.RegisterHandlers(router)

	addr := s.config.Host + ":" + s.config.Port
	log.Printf("Server is running on %s", addr)
	return http.ListenAndServe(addr, router)
}
