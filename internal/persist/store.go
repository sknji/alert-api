package persist

import (
	"github.com/sknji/alert-api/internal/models"
)

type Store interface {
	Save(alert *models.Alert) error
	Get(alertId string) (alert *models.Alert, err error)
	Search(serviceId, startTs, endTs string) (alerts *models.ServiceAlerts, err error)
}
