package persist

import (
	log "github.com/sirupsen/logrus"
	"github.com/sknji/alert-api/internal/models"
)

type InMemoryStorage struct {
	services map[string]*models.Service
	alerts   map[string]*models.Alert
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		services: make(map[string]*models.Service),
		alerts:   make(map[string]*models.Alert),
	}
}

func (ims *InMemoryStorage) SaveService(serv *models.Service) error {
	log.Infoln("InMemoryStorage saving Service", serv)
	ims.services[serv.ServiceId] = serv
	return nil
}

func (ims *InMemoryStorage) GetService(serviceId string) (serv *models.Service, err error) {
	log.Infoln("InMemoryStorage fetching Service", serviceId)
	serv, ok := ims.services[serviceId]
	if !ok {
		return nil, ErrEntityNotFound
	}
	return serv, nil
}

func (ims *InMemoryStorage) SaveAlert(alert *models.Alert) error {
	log.Infoln("InMemoryStorage saving Alert", alert)
	ims.alerts[alert.AlertId] = alert
	return nil
}
func (ims *InMemoryStorage) SaveAlerts(alerts []*models.Alert) error {
	log.Infoln("InMemoryStorage saving Alerts", alerts)
	for _, alert := range alerts {
		ims.alerts[alert.AlertId] = alert
	}
	return nil
}
func (ims *InMemoryStorage) GetAlert(alertId string) (alert *models.Alert, err error) {
	log.Infoln("InMemoryStorage fetching Alert", alertId)
	alert, ok := ims.alerts[alertId]
	if !ok {
		return nil, ErrEntityNotFound
	}
	return alert, nil
}

func (ims *InMemoryStorage) FindAlerts(serviceId, startTs, endTs string) (alerts []*models.Alert, err error) {
	log.Infoln("InMemoryStorage Searching service alerts", serviceId, startTs, endTs)
	for _, alert := range ims.alerts {
		if alert.ServiceId != serviceId {
			continue
		}

		if alert.AlertTs >= startTs && alert.AlertTs <= endTs {
			alerts = append(alerts, alert)
		}
	}

	if len(alerts) == 0 {
		return nil, ErrEntityNotFound
	}
	return
}
