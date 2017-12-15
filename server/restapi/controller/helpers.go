package controller

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/zeebox/terraform-server/server/models"
	"net/http"
	"strings"
)

func str(v string) *string { return &v }

func halRootRscLinks(oh ContextHelper) *models.HalRscLinks {
	lnks := halSelfLink(oh.FQEndpoint)
	lnks.Doc = &models.HalHref{
		Href: strfmt.URI(oh.ServerURL + "/docs#operation/" + oh.OperationID),
	}
	return lnks
}

func halSelfLink(href string) *models.HalRscLinks {
	return &models.HalRscLinks{
		Self: &models.HalHref{Href: strfmt.URI(href)},
	}
}

// NewContextHelper to easily get URL parts for generating HAL resources
func NewContextHelper(ctx *middleware.Context) ContextHelper {
	return ContextHelper{
		ctx: ctx,
	}
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
}

// GetAPIParts from the combination of the request and the API. This is used to generate HAL resources
func (oh ContextHelper) GetAPIParts(basePath string) {

	oh.ServerURL = oh.urlPrefix(oh.Request.Host, basePath, oh.Request.TLS != nil)
	oh.FQEndpoint = oh.urlPrefix(oh.Request.Host, oh.Request.RequestURI, oh.Request.TLS != nil)

	oh.endpoint = strings.TrimPrefix(oh.FQEndpoint, oh.ServerURL)
	r, _ := oh.ctx.LookupRoute(oh.Request)
	oh.OperationID = r.Operation.ID
}

func (oh ContextHelper) urlPrefix(host string, uri string, https bool) string {
	prefix := "http"
	if https {
		prefix += "s"
	}
	return strings.TrimSuffix(prefix+"://"+host+uri, "/")
}
