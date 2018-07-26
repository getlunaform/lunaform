package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
	"net/http"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime"
)

type CommonServerErrorResponder struct {
	Payload *models.ServerError
	code    int
}

func NewServerError(code int32, errorString string) (r middleware.Responder) {
	return &CommonServerErrorResponder{
		Payload: &models.ServerError{
			Message: swag.String(errorString),
			Status: swag.String(http.StatusText(
				int(code),
			)),
			StatusCode: &code,
		},
		code: int(code),
	}
}

// WriteResponse to the client
func (cser *CommonServerErrorResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(cser.code)
	if cser.Payload != nil {
		payload := cser.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
