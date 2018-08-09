package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime"
	"net/http"
)

type CommonServerErrorResponder struct {
	Payload *models.ServerError
	code    int
}

func NewServerErrorResponse(code int, errorString string) (r *CommonServerErrorResponder) {
	return &CommonServerErrorResponder{
		Payload: models.NewServerError(int32(code), errorString),
		code:    int(code),
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

func (cser *CommonServerErrorResponder) Error() {

}
