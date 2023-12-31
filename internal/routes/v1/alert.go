package v1

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"github.com/sknji/alert-api/internal/models"
	"github.com/sknji/alert-api/internal/service"
	"net/http"
)

type alertResource struct{}

func RegisterAlertRoutes(ctx context.Context) http.Handler {
	r := chi.NewRouter()
	ar := &alertResource{}
	r.Get("/", ar.search(ctx))
	r.Get("/{alertID}", ar.get(ctx))
	r.Post("/", ar.post(ctx))
	return r
}

func (ar *alertResource) post(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var alert models.AlertRequest

		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			_ = render.Render(w, r, models.ErrorInvalidRequest(alert.AlertId, err))
			return
		}
		_ = r.Body.Close()

		errResp := service.AlertService(ctx).CreateAlert(ctx, &alert)
		if errResp != nil {
			_ = render.Render(w, r, errResp)
			return
		}
		_ = render.Render(w, r, models.NewAlertResponse(alert.AlertId))
	}
}

func (ar *alertResource) get(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		alertId := chi.URLParam(r, "alertID")
		log.Infoln("alertID received:", alertId)

		alert, errResp := service.AlertService(ctx).FetchAlert(ctx, alertId)
		if errResp != nil {
			_ = render.Render(w, r, errResp)
			return
		}
		_ = render.Render(w, r, alert)
	}
}

func (ar *alertResource) search(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTs := r.URL.Query().Get("start_ts")
		endTs := r.URL.Query().Get("end_ts")
		serviceId := r.URL.Query().Get("service_id")

		log.Infof("service_id:%s, start_ts:%s, end_ts:%s\n", serviceId, startTs, endTs)

		alerts, errResp := service.AlertService(ctx).SearchServiceAlerts(ctx, serviceId, startTs, endTs)
		if errResp != nil {
			_ = render.Render(w, r, errResp)
			return
		}
		_ = render.Render(w, r, alerts)
	}
}
