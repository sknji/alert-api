package service

import (
	"context"
	"errors"
	"github.com/sknji/alert-api/internal/models"
	"github.com/sknji/alert-api/internal/persist"
)

type AlertServiceI interface {
	CreateAlert(ctx context.Context, m *models.AlertRequest) *models.ErrorResponse
	FetchAlert(ctx context.Context, alertId string) (*models.Alert, *models.ErrorResponse)
	SearchServiceAlerts(ctx context.Context, serviceId string, startTs string, endTs string) (*models.Service, *models.ErrorResponse)
}

var ErrSearchResultNotFound = errors.New("search results not found for provided serviceId, startTs and endTs")
var ErrAlertNotExist = errors.New("alert record does not exist")
var ErrAlertExist = errors.New("alert record exist")

type AlertServiceImpl struct{}

func (asi *AlertServiceImpl) ValidateNormalizeAlertReq(req *models.AlertRequest) *models.Service {
	var serv models.Service
	serv.ServiceId = req.ServiceId
	serv.ServiceName = req.ServiceName

	alert := req.GetAlert()
	alert.ServiceId = req.ServiceId
	serv.AddAlert(alert)
	return &serv
}

func (asi *AlertServiceImpl) CreateAlert(
	ctx context.Context, alert *models.AlertRequest) *models.ErrorResponse {
	serv := asi.ValidateNormalizeAlertReq(alert)

	_, err := Storage(ctx).GetService(serv.ServiceId)
	if errors.Is(err, persist.ErrEntityNotFound) {
		err = Storage(ctx).SaveService(serv)
	}
	if err != nil {
		return models.ErrorInternalServerError(alert.AlertId, err)
	}

	a := serv.Alerts[0]

	entity, err := Storage(ctx).GetAlert(alert.AlertId)
	if err == nil && entity != nil {
		return models.ErrorInvalidRequest(alert.AlertId, ErrAlertExist)
	}

	err = Storage(ctx).SaveAlert(a)
	return models.ErrorInternalServerError(alert.AlertId, err)
}

func (asi *AlertServiceImpl) FetchAlert(
	ctx context.Context, alertId string) (*models.Alert, *models.ErrorResponse) {
	alert, err := Storage(ctx).GetAlert(alertId)
	if err != nil {
		if errors.Is(err, persist.ErrEntityNotFound) {
			return nil, models.ErrorNotFound(alertId, ErrAlertNotExist)
		}
		return nil, models.ErrorInternalServerError(alertId, err)
	}
	return alert, nil
}

func (asi *AlertServiceImpl) SearchServiceAlerts(
	ctx context.Context, serviceId string, startTs string, endTs string) (*models.Service, *models.ErrorResponse) {
	serv, err := Storage(ctx).GetService(serviceId)
	if err != nil {
		return nil, searchError(err)
	}

	alerts, err := Storage(ctx).FindAlerts(serviceId, startTs, endTs)
	if err != nil {
		return nil, searchError(err)
	}

	if len(alerts) == 0 {
		return nil, searchError(ErrSearchResultNotFound)
	}

	serv.SetAlerts(alerts...)
	return serv, nil
}

func searchError(err error) *models.ErrorResponse {
	if errors.Is(err, persist.ErrEntityNotFound) {
		return models.ErrorNotFound("", ErrSearchResultNotFound)
	}
	return models.ErrorInternalServerError("", err)
}
