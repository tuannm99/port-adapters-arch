package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/tuannm99/edge-platform/apps/control-plane/internal/adapters/http/handler"
)

func New(websiteHandler *handler.WebsiteHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	r.Route("/websites", func(r chi.Router) {
		r.Post("/", websiteHandler.Create)
		r.Get("/", websiteHandler.List)
		r.Get("/{id}/nginx-config", websiteHandler.GetNginxConfig)
	})

	return r
}
