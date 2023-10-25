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
