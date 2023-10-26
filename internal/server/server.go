package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"
	v1 "github.com/sknji/alert-api/internal/routes/v1"
	"net/http"
)

func StartHTTPServer(ctx context.Context, mux *chi.Mux, port string) error {
	setupCors(mux)
	addMiddlewares(mux)
	registerV1Routes(ctx, mux)

	log.Infof("HTTP Server running at port %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func registerV1Routes(ctx context.Context, mux *chi.Mux) {
	mux.Route("/v1", v1.RegisterRoutes(ctx))
}

func addMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func setupCors(mux *chi.Mux) {
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: false,
	}))
}
