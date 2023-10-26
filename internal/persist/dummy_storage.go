package persist

import (
	"github.com/sknji/alert-api/internal/models"
	"log"
)

type DummyStorage struct {
}

func NewDummyStore() *DummyStorage {
	return &DummyStorage{}
}

func (ds *DummyStorage) SaveService(serv *models.Service) error {
	log.Println("DummyStorage saving Service", serv)
	return nil
}

func (ds *DummyStorage) GetService(serviceId string) (serv *models.Service, err error) {
	log.Println("DummyStorage fetching Service", serv)
	return nil, ErrEntityNotFound
}

func (ds *DummyStorage) SaveAlert(alert *models.Alert) error {
	log.Println("DummyStorage saving Alert", alert)
	return nil
}
func (ds *DummyStorage) SaveAlerts(alerts []*models.Alert) error {
	log.Println("DummyStorage saving Alerts", alerts)
	return nil
}
func (ds *DummyStorage) GetAlert(alertId string) (alert *models.Alert, err error) {
	log.Println("DummyStorage fetching Alert", alertId)
	return nil, ErrEntityNotFound
}

func (ds *DummyStorage) FindAlerts(serviceId, startTs, endTs string) (alerts []*models.Alert, err error) {
	log.Println("DummyStorage Searching service alerts", serviceId, startTs, endTs)
	return nil, ErrEntityNotFound
}
