package apiserver

import (
	"msqrd/sneakers/internal/router"
	"net/http"
)

type APIServer struct {
	config *Config
	router *router.Router
}

// Configure new APIServer and return pointer of it
func NewServer(cfg *Config) *APIServer {
	return &APIServer{
		config: cfg,
		router: router.NewRouter(),
	}
}

// Start server
func (s *APIServer) Start() error {
	s.ConfigureRouter()

	return http.ListenAndServe(s.config.BindAdress, nil)
}

// Configure router endpoints
func (s *APIServer) ConfigureRouter() {}
