package models

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	AlertId    string `json:"alert_id,omitempty"`
	StatusCode int    `json:"-"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorInternalServerError(alertId string, err error) *ErrorResponse {
	if err == nil {
		return nil
	}

	msg := err.Error()
	return &ErrorResponse{
		AlertId:    alertId,
		Error:      msg,
		StatusCode: http.StatusInternalServerError,
	}
}

func ErrorInvalidRequest(alertId string, err error) *ErrorResponse {
	if err == nil {
		return nil
	}

	msg := err.Error()
	return &ErrorResponse{
		AlertId:    alertId,
		Error:      msg,
		StatusCode: http.StatusBadRequest,
	}
}

func ErrorNotFound(alertId string, err error) *ErrorResponse {
	if err == nil {
		return nil
	}

	msg := err.Error()
	return &ErrorResponse{
		AlertId:    alertId,
		Error:      msg,
		StatusCode: http.StatusNotFound,
	}
}
