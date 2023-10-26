package persist

import (
	log "github.com/sirupsen/logrus"
	"github.com/sknji/alert-api/internal/models"
)

type DummyStorage struct {
}

func NewDummyStore() *DummyStorage {
	return &DummyStorage{}
}

func (ds *DummyStorage) SaveService(serv *models.Service) error {
	log.Infoln("DummyStorage saving Service", serv)
	return nil
}

func (ds *DummyStorage) GetService(serviceId string) (serv *models.Service, err error) {
	log.Infoln("DummyStorage fetching Service", serv)
	return nil, ErrEntityNotFound
}

func (ds *DummyStorage) SaveAlert(alert *models.Alert) error {
	log.Infoln("DummyStorage saving Alert", alert)
	return nil
}
func (ds *DummyStorage) SaveAlerts(alerts []*models.Alert) error {
	log.Infoln("DummyStorage saving Alerts", alerts)
	return nil
}
func (ds *DummyStorage) GetAlert(alertId string) (alert *models.Alert, err error) {
	log.Infoln("DummyStorage fetching Alert", alertId)
	return nil, ErrEntityNotFound
}

func (ds *DummyStorage) FindAlerts(serviceId, startTs, endTs string) (alerts []*models.Alert, err error) {
	log.Infoln("DummyStorage Searching service alerts", serviceId, startTs, endTs)
	return nil, ErrEntityNotFound
}
