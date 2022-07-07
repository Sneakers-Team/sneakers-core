package apiserver

import "net/http"

type APIServer struct {
	config *Config
}

func NewServer(cfg *Config) *APIServer {
	return &APIServer{
		config: cfg,
	}
}

func (s *APIServer) Start() error {
	return http.ListenAndServe(s.config.BindAdress, nil)
}
