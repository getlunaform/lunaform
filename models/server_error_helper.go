package models

import (
	"github.com/go-openapi/swag"
	"net/http"
)

func NewServerError(code int32, errorString string) *ServerError {
	return &ServerError{
		Message: swag.String(errorString),
		Status: swag.String(http.StatusText(
			int(code),
		)),
		StatusCode: &code,
	}
}
