package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

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
	ctx         *middleware.Context
	Request     *http.Request
	ServerURL   string
	endpoint    string
	FQEndpoint  string
	OperationID string
	BasePath    string
}

// SetRequest and calculate the api parts from the combination of the request and the API. This is used to generate HAL resources
func (oh *ContextHelper) SetRequest(req *http.Request) (err error) {

	oh.Request = req

	oh.ServerURL = oh.urlPrefix(oh.Request.Host, oh.BasePath, oh.Request.TLS != nil)
	oh.FQEndpoint = oh.urlPrefix(oh.Request.Host, oh.Request.RequestURI, oh.Request.TLS != nil)

	oh.endpoint = strings.TrimPrefix(oh.FQEndpoint, oh.ServerURL)
	if r, matched := oh.ctx.LookupRoute(oh.Request); matched {
		oh.OperationID = r.Operation.ID
	} else {
		return errors.New("Could not find route for request")
	}

	return
}

func (oh ContextHelper) urlPrefix(host string, uri string, https bool) string {
	prefix := "http"
	if https {
		prefix += "s"
	}
	return strings.TrimSuffix(prefix+"://"+host+uri, "/")
}

func String(s string) *string {
	return &s
}

func Int64(i int64) *int64 {
	return &i
}