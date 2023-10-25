package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/sknji/alert-api/internal/persist"
	"github.com/sknji/alert-api/internal/server"
	"github.com/sknji/alert-api/internal/service"
	"log"
)

func LoadServices(ctx context.Context) context.Context {
	db, err := persist.NewDummyStore()
	if err != nil {
		log.Panicln(err)
	}

	return service.WithStorage(ctx, db)
}

func main() {
	// TODO: read port from env variable -> cli -> file
	port := 8080

	ctx := LoadServices(context.Background())

	r := chi.NewRouter()

	err := server.StartHTTPServer(ctx, r, port)
	if err != nil {
		log.Panicln(err)
	}
}
