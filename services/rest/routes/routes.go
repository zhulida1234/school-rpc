package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/zhulida1234/school-rpc/services/rest/service"
)

type Routes struct {
	router *chi.Mux
	svc    service.Service
}

func NewRoutes(r *chi.Mux, svc service.Service) Routes {
	return Routes{
		router: r,
		svc:    svc,
	}
}
