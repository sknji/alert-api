package v1

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sknji/alert-api/internal/models"
	"github.com/sknji/alert-api/internal/service"
	"log"
	"net/http"
)

type alertResource struct{}

func RegisterAlert(ctx context.Context) http.Handler {
	r := chi.NewRouter()
	ar := &alertResource{}
	r.Post("/", ar.post(ctx))
	r.Get("/", ar.search(ctx))
	r.Get("/{alertID}", ar.get(ctx))
	return r
}

func (ar alertResource) post(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var alert models.Alert

		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			_ = render.Render(w, r, models.ErrorInvalidRequest(alert.AlertId, err))
			return
		}

		err = service.Storage(ctx).Save(&alert)
		if err != nil {
			_ = render.Render(w, r, models.ErrorInternalServerError(alert.AlertId, err))
			return
		}

		render.Status(r, http.StatusNoContent)
	}
}

func (ar alertResource) get(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		alertId := chi.URLParam(r, "alertID")
		log.Println("alertID received:", alertId)

		alert, err := service.Storage(ctx).Get(alertId)
		if err != nil {
			_ = render.Render(w, r, models.ErrorNotFound(alertId, err))
			return
		}

		_ = render.Render(w, r, alert)
	}
}

func (ar alertResource) search(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceId := r.URL.Query().Get("service_id")
		startTs := r.URL.Query().Get("start_ts")
		endTs := r.URL.Query().Get("end_ts")

		log.Printf("service_id:%s, start_ts:%s, end_ts:%s", serviceId, startTs, endTs)

		alerts, err := service.Storage(ctx).Search(serviceId, startTs, endTs)
		if err != nil {
			_ = render.Render(w, r, models.ErrorNotFoundSearch(err))
			return
		}

		_ = render.Render(w, r, alerts)
	}
}
