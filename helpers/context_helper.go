package helpers

import (
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"strings"
	"errors"
)


func NewContextHelperWithContext(ctx *middleware.Context) *ContextHelper {
	return NewContextHelper(ctx).
		WithBasePath(ctx.BasePath())
}

// NewContextHelper to easily get URL parts for generating HAL resources
func NewContextHelper(ctx *middleware.Context) *ContextHelper {
	return &ContextHelper{
		ctx: ctx,
	}
}

func (ch *ContextHelper) WithBasePath(basePath string) *ContextHelper {
	ch.BasePath = basePath
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
	PathParts        []string
}

// SetRequest and calculate the api parts from the combination of the request and the API. This is used to generate HAL resources
func (ch *ContextHelper) SetRequest(req *http.Request) (err error) {
	ch.Request = req
	ch.ParseRequest(
		ch.Request.Host,
		ch.BasePath,
		ch.Request.RequestURI,
		ch.Request.TLS != nil,
	)

	r, matched := ch.ctx.LookupRoute(ch.Request)
	if ! matched {
		return errors.New("could not find route for request")
	}
	ch.OperationID = r.Operation.ID
	return
}

func (ch *ContextHelper) ParseRequest(host string, basePath string,
	requestUri string, hasTls bool) {

	ch.ServerURL = ch.urlPrefix(host, basePath, hasTls)
	ch.FQEndpoint = ch.urlPrefix(host, requestUri, hasTls)

	ch.Endpoint = strings.TrimPrefix(ch.FQEndpoint, ch.ServerURL)
	ch.EndpointSingular = ch.Endpoint
	if strings.HasSuffix(ch.Endpoint, "s") {
		ch.EndpointSingular = strings.TrimSuffix(ch.Endpoint, "s")
	}
	ch.PathParts = strings.Split(ch.FQEndpoint, "/")
}

func (ch *ContextHelper) urlPrefix(host string, uri string, https bool) string {
	prefix := "http"
	if https {
		prefix += "s"
	}
	return strings.TrimSuffix(prefix+"://"+host+uri, "/")
}

