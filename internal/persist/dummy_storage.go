package persist

import (
	"github.com/sknji/alert-api/internal/models"
	"log"
)

type DummyStore struct {
}

func NewDummyStore() (*DummyStore, error) {
	return &DummyStore{}, nil
}

func (DummyStore) SaveService(serv *models.Service) error {
	log.Println("DummyStore saving Service", serv)
	return nil
}

func (DummyStore) GetService(serviceId string) (serv *models.Service, err error) {
	log.Println("DummyStore fetching Service", serv)
	return nil, ErrEntityNotFound
}

func (DummyStore) SaveAlert(alert *models.Alert) error {
	log.Println("DummyStore saving Alert", alert)
	return nil
}
func (DummyStore) SaveAlerts(alerts []*models.Alert) error {
	log.Println("DummyStore saving Alerts", alerts)
	return nil
}
func (DummyStore) GetAlert(alertId string) (alert *models.Alert, err error) {
	log.Println("DummyStore fetching Alert", alertId)
	return nil, ErrEntityNotFound
}

func (DummyStore) FindAlerts(serviceId, startTs, endTs string) (alerts []*models.Alert, err error) {
	log.Println("DummyStore Searching service alerts", serviceId, startTs, endTs)
	return nil, ErrEntityNotFound
}
