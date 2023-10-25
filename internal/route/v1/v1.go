package v1

import (
	"context"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(ctx context.Context, r chi.Router) {
	r.Mount("/alert", RegisterAlert(ctx))
}
