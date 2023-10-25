package main

import (
	"context"
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/sknji/alert-api/internal/config"
	"github.com/sknji/alert-api/internal/persist"
	"github.com/sknji/alert-api/internal/server"
	"github.com/sknji/alert-api/internal/service"
	"log"
)

func LoadServices(ctx context.Context, conf config.Configuration) context.Context {
	db, err := persist.NewDatabaseStorage(&conf.Database)
	if err != nil {
		log.Panicln(err)
	}

	ctx = service.WithStorage(ctx, db)
	ctx = service.WithAlertService(ctx, &service.AlertServiceImpl{})
	return ctx
}

func main() {
	configFile := flag.String("config", "config-local.yaml", "")
	flag.Parse()

	conf, err := config.LoadConfigs(*configFile)
	if err != nil {
		log.Panicln(err)
	}

	ctx := LoadServices(context.Background(), conf)

	r := chi.NewRouter()

	if err := server.StartHTTPServer(ctx, r, conf.Port); err != nil {
		log.Panicln(err)
	}
}
