package api

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	userHandler *UserHandler
}

func NewRouter(userHandler *UserHandler) *Router {
	return &Router{userHandler: userHandler}
}

func (r *Router) RegisterRoutes(ro chi.Router) {
	ro.Post("/api/v1/auth/register", r.userHandler.Create)
}
