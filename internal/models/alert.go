package models

import (
	"net/http"
)

type Alert struct {
	AlertId     string `json:"alert_id"`
	ServiceId   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Model       string `json:"model"`
	AlertType   string `json:"alert_type"`
	AlertTs     string `json:"alert_ts"`
	Severity    string `json:"severity"`
	TeamSlack   string `json:"team_slack"`
}

func (a *Alert) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

type ServiceAlerts struct {
	ServiceId   string  `json:"service_id"`
	ServiceName string  `json:"service_name"`
	Alerts      []Alert `json:"alerts"`
}

func (a *ServiceAlerts) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}
