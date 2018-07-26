package helpers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
)

func NewServerError(code int32, errorString string) *models.ServerError {
	return &models.ServerError{
		Message: swag.String(errorString),
		Status: swag.String(http.StatusText(
			int(code),
		)),
		StatusCode: &code,
	}
}

// NewContextHelper to easily get URL parts for generating HAL resources
func NewContextHelper(ctx *middleware.Context) ContextHelper {
	ch := ContextHelper{
		ctx: ctx,
	}
	ch.BasePath = ctx.BasePath()
	return ch
}

// ContextHelper is split into its own little function, as test it is really difficult due to the un-exported nature
// of the majority of the `MatchedRoute` struct, which means it's very difficult to generate a mock response to
// ctx.LookupRoute. Doing it this way, means we can mock it in tests.
type ContextHelper struct {
	ctx              *middleware.Context
	Request          *http.Request
	ServerURL        string
	Endpoint         string
	EndpointSingular string
	FQEndpoint       string
	OperationID      string
	BasePath         string
}

// SetRequest and calculate the api parts from the combination of the request and the API. This is used to generate HAL resources
func (ch *ContextHelper) SetRequest(req *http.Request) (err error) {

	ch.Request = req

	ch.ServerURL = ch.urlPrefix(ch.Request.Host, ch.BasePath, ch.Request.TLS != nil)
	ch.FQEndpoint = ch.urlPrefix(ch.Request.Host, ch.Request.RequestURI, ch.Request.TLS != nil)

	ch.Endpoint = strings.TrimPrefix(ch.FQEndpoint, ch.ServerURL)
	ch.EndpointSingular = ch.Endpoint
	if strings.HasSuffix(ch.Endpoint, "s") {
		ch.EndpointSingular = strings.TrimSuffix(ch.Endpoint, "s")
	}
	if r, matched := ch.ctx.LookupRoute(ch.Request); matched {
		ch.OperationID = r.Operation.ID
	} else {
		return errors.New("Could not find route for request")
	}

	return
}

func (ch ContextHelper) urlPrefix(host string, uri string, https bool) string {
	prefix := "http"
	if https {
		prefix += "s"
	}
	return strings.TrimSuffix(prefix+"://"+host+uri, "/")
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
