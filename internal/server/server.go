package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sknji/alert-api/internal/route/v1"
	"log"
	"net/http"
)

func StartHTTPServer(ctx context.Context, r *chi.Mux, port int) error {
	addMiddlewares(r)

	r.Route("/v1", func(r chi.Router) {
		v1.RegisterRoutes(ctx, r)
	})

	log.Printf("HTTP Server running at port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func addMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}
