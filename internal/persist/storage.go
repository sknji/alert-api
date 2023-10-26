package persist

import (
	"errors"
	"github.com/sknji/alert-api/internal/models"
)

// Store defines an interfaces that requires all storage systems to implement
type Store interface {
	GetService(serviceId string) (serv *models.Service, err error)
	SaveService(serv *models.Service) error
	GetAlert(alertId string) (alert *models.Alert, err error)
	SaveAlert(alert *models.Alert) error
	SaveAlerts(alert []*models.Alert) error
	FindAlerts(serviceId, startTs, endTs string) (alerts []*models.Alert, err error)
}

var ErrEntityNotFound = errors.New("entity not found")
