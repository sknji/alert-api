package models

import (
	"net/http"
)

type Alert struct {
	AlertId   string `json:"alert_id" gorm:"primaryKey"`
	ServiceId string `json:"service_id" gorm:"column:service_id"`
	Model     string `json:"model"`
	AlertType string `json:"alert_type"`
	AlertTs   string `json:"alert_ts"`
	Severity  string `json:"severity"`
	TeamSlack string `json:"team_slack"`
}

func (*Alert) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (*Alert) TableName() string {
	return "alert"
}

type Service struct {
	ServiceId   string   `json:"service_id" gorm:"primaryKey"`
	ServiceName string   `json:"service_name"`
	Alerts      []*Alert `json:"alerts" gorm:"foreignKey:ServiceId;references:ServiceId"`
}

func (s *Service) AddAlert(alert *Alert) {
	for _, a := range s.Alerts {
		if a.AlertId == alert.AlertId {
			return
		}
	}
	s.Alerts = append(s.Alerts, alert)
}

func (s *Service) SetAlerts(alerts ...*Alert) {
	s.Alerts = alerts
}

func (*Service) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (*Service) TableName() string {
	return "service_alert"
}
