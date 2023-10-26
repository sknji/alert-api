package v1

import (
	"context"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(ctx context.Context) func(r chi.Router) {
	return func(r chi.Router) {
		r.Mount("/alert", RegisterAlertRoutes(ctx))
	}
}
