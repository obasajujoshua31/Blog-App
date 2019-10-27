package server

import (
	"github.com/go-chi/chi"
)

type Router struct {
	*chi.Mux
}

func NewRouter() *Router {
	return &Router{chi.NewRouter()}
}
