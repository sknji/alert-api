package persist

import (
	"errors"
	"github.com/sknji/alert-api/internal/models"
	"log"
)

type DatabaseStore struct {
}

func NewDatabaseStore() (*DatabaseStore, error) {
	return &DatabaseStore{}, nil
}

func (DatabaseStore) Save(alert *models.Alert) error {
	log.Println("MariaDB Saving alert", alert)
	return nil
}

func (DatabaseStore) Get(uuid string) (alert *models.Alert, err error) {
	log.Println("MariaDB Fetching alert", uuid)
	return &models.Alert{}, nil
}

func (DatabaseStore) Search(serviceId, startTs, endTs string) (alerts *models.ServiceAlerts, err error) {
	log.Println("MariaDB Searching service alerts", serviceId, startTs, endTs)
	return &models.ServiceAlerts{}, errors.New("invalid")
}
