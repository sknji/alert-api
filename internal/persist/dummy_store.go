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

func (DummyStore) Save(alert *models.Alert) error {
	log.Println("DummyStore saving Alert", alert)
	return nil
}

func (DummyStore) Get(uuid string) (alert *models.Alert, err error) {
	log.Println("DummyStore fetching Alert", uuid)
	return &models.Alert{}, nil
}

func (DummyStore) Search(serviceId, startTs, endTs string) (alerts *models.ServiceAlerts, err error) {
	log.Println("DummyStore Searching service alerts", serviceId, startTs, endTs)
	return &models.ServiceAlerts{}, nil
}
