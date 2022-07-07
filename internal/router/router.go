package router

import "net/http"

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

// Endpoint handler
func (r *Router) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, handler)
}
