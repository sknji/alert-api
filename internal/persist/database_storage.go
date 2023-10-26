package persist

import (
	"errors"
	"github.com/sknji/alert-api/internal/models"
	"github.com/sknji/alert-api/internal/persist/database"
	"gorm.io/gorm"
	"log"
)

type DatabaseStorage struct {
	handle *gorm.DB
}

func NewDatabaseStorage(conf *database.Config) (*DatabaseStorage, error) {
	db, err := gorm.Open(database.DbConn(conf), database.DbConfig(conf))
	if err != nil {
		return nil, err
	}

	if conf.AutoMigrate {
		_ = db.AutoMigrate(models.Alert{}, models.Service{})
	}

	return &DatabaseStorage{handle: db}, nil
}

func (db *DatabaseStorage) SaveService(serv *models.Service) error {
	log.Println("DatabaseStorage Saving alert", serv)
	gormResult := db.handle.Model(models.Service{}).Save(serv)
	return StandardError(gormResult.Error)
}

func (db *DatabaseStorage) GetService(serviceId string) (*models.Service, error) {
	log.Println("DatabaseStorage Fetching alert", serviceId)
	var serv models.Service
	gormResult := db.handle.Model(models.Service{}).
		Where("service_id = ?", serviceId).First(&serv)
	return &serv, StandardError(gormResult.Error)
}

func (db *DatabaseStorage) SaveAlert(alert *models.Alert) error {
	log.Println("DatabaseStorage Saving alert", alert)
	gormResult := db.handle.Model(models.Alert{}).Create(alert)
	return StandardError(gormResult.Error)
}

func (db *DatabaseStorage) SaveAlerts(alerts []*models.Alert) error {
	log.Println("DatabaseStorage Saving alert", alerts)
	gormResult := db.handle.Model(models.Alert{}).Create(alerts)
	return StandardError(gormResult.Error)
}

func (db *DatabaseStorage) GetAlert(alertId string) (*models.Alert, error) {
	log.Println("DatabaseStorage Fetching alert", alertId)
	var alert models.Alert
	gormResult := db.handle.Model(models.Alert{}).
		Where("alert_id = ?", alertId).First(&alert)
	return &alert, StandardError(gormResult.Error)
}

func (db *DatabaseStorage) FindAlerts(serviceId, startTs, endTs string) ([]*models.Alert, error) {
	log.Println("DatabaseStorage Searching service alerts", serviceId, startTs, endTs)
	gormResult := db.handle.Model(models.Alert{})
	if serviceId != "" {
		gormResult = gormResult.Where("service_id = ?", serviceId)
	}

	if startTs != "" && endTs != "" {
		gormResult = gormResult.
			Where("alert_ts >= ?", startTs).
			Where("alert_ts <= ?", endTs)
	} else if endTs != "" {
		gormResult = gormResult.Where("alert_ts <= ?", endTs)
	} else if startTs != "" {
		gormResult = gormResult.Where("alert_ts >= ?", startTs)
	}

	var alerts []*models.Alert
	gormResult = gormResult.Find(&alerts)
	return alerts, StandardError(gormResult.Error)
}

func StandardError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrEntityNotFound
	}
	return err
}
