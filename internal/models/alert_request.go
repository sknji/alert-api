package models

import "net/http"

type AlertRequest struct {
	Alert
	ServiceName string `json:"service_name"`
	ServiceId   string `json:"service_id"`
}

func (*AlertRequest) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (a *AlertRequest) GetAlert() *Alert {
	return &a.Alert
}

type AlertResponse struct {
	AlertId string `json:"alert_id,omitempty"`
}

func NewAlertResponse(alertId string) *AlertResponse {
	return &AlertResponse{AlertId: alertId}
}

func (*AlertResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}
