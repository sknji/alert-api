package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sknji/alert-api/internal/models"
	"github.com/sknji/alert-api/internal/persist"
	"github.com/sknji/alert-api/internal/service"
	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

var alertId = "b950482e9911ec7e41f7ca5e5d9a427g2"
var invalidAlertId = "b950482e9911ec7e41f7ca5e5_invalid"

var createAlertBody = []byte(
	`{
    "alert_id": "` + alertId + `",
    "service_id": "my_test_service_id_2",
    "service_name": "my_test_service",
    "model": "my_test_model",
    "alert_type": "anomaly",
    "alert_ts": "1695644699",
    "severity": "warning",
    "team_slack": "slack_ch"
}`)

var v1Routes = RegisterAlertRoutes(helperContext())

func execute(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	v1Routes.ServeHTTP(rr, req)
	return rr
}

func helperContext() context.Context {
	ctx := context.Background()
	ctx = service.WithStorage(ctx, persist.NewInMemoryStorage())
	ctx = service.WithAlertService(ctx, &service.AlertServiceImpl{})
	return ctx
}

func TestAlertsRoute_PostAlert_Success_1(t *testing.T) {
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(createAlertBody))
	resp := execute(req)
	respBody := make(map[string]string)
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	assert.Nil(t, err)
	assert.Equal(t, alertId, respBody["alert_id"])
	assert.Equal(t, http.StatusOK, resp.Code, "no content success response expected")
}

func TestAlertsRoute_PostAlert_Failed_1(t *testing.T) {
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(createAlertBody))
	resp := execute(req)
	respBody := make(map[string]string)
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	assert.Nil(t, err)
	assert.Equal(t, alertId, respBody["alert_id"])
	assert.Equal(t, "alert record exist", respBody["error"])
	assert.Equal(t, http.StatusBadRequest, resp.Code, "cannot add since alert already exist")
}

func TestAlertsRoute_GetAlert_Success_2(t *testing.T) {
	req, _ := http.NewRequest("GET", "/"+alertId, nil)
	resp := execute(req)

	var alert models.Alert
	err := json.NewDecoder(resp.Body).Decode(&alert)

	assert.Nil(t, err)
	assert.Equal(t, "my_test_service_id_2", alert.ServiceId)
	assert.Equal(t, http.StatusOK, resp.Code, "alert record should be saved from prev test")
}

func TestAlertsRoute_GetAlert_Failed_3(t *testing.T) {
	req, _ := http.NewRequest("GET", "/"+invalidAlertId, nil)
	resp := execute(req)
	assert.Equal(t, http.StatusNotFound, resp.Code, "alert record for invalid alertId should not exist")
}

func TestAlertsRoute_GetAlertService_Success_4(t *testing.T) {
	queryParams := "?service_id=my_test_service_id_2&start_ts=1695644600&end_ts=1695644800"
	req, _ := http.NewRequest("GET", "/"+queryParams, nil)
	resp := execute(req)

	var serv models.Service
	err := json.NewDecoder(resp.Body).Decode(&serv)

	assert.Nil(t, err)
	assert.Equal(t, "my_test_service_id_2", serv.ServiceId)
	assert.Equal(t, "my_test_service", serv.ServiceName)
	assert.Equal(t, 1, len(serv.Alerts))
	assert.Equal(t, http.StatusOK, resp.Code, "alert record should be saved from prev test")
}

func TestAlertsRoute_GetAlertService_Failed_5(t *testing.T) {
	queryParams := "?service_id=invalid_service_id&start_ts=1695644600&end_ts=1695644800"
	req, _ := http.NewRequest("GET", "/"+queryParams, nil)
	resp := execute(req)

	errResp := make(map[string]string)
	err := json.NewDecoder(resp.Body).Decode(&errResp)

	log.Infoln(errResp)
	assert.Nil(t, err)
	assert.Equal(t, service.ErrSearchResultNotFound.Error(), errResp["error"])
	assert.Equal(t, http.StatusNotFound, resp.Code, "alert record should not exist for invalid service id")
}
