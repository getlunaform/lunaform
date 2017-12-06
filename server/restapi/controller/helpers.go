package controller

import (
	"github.com/go-openapi/strfmt"
	"github.com/zeebox/terraform-server/server/models"
	"github.com/zeebox/terraform-server/server/restapi/operations"
	"net/http"
	"strings"
)

func str(v string) *string { return &v }

func HALRootRscLinks(parts *APIHostBase) *models.HalRscLinks {
	lnks := HALSelfLink(parts.FQEndpoint)
	lnks.Doc = &models.HalHref{
		Href: strfmt.URI(parts.ServerURL + "/docs#operation/" + parts.OperationId),
	}
	return lnks
}
func HALSelfLink(href string) *models.HalRscLinks {
	return &models.HalRscLinks{
		Self: &models.HalHref{Href: strfmt.URI(href)},
	}
}

type APIHostBase struct {
	ServerURL   string
	Endpoint    string
	FQEndpoint  string
	OperationId string
}

func apiParts(req *http.Request, api *operations.TerraformServerAPI) *APIHostBase {
	prefix := "http"
	if req.TLS != nil {
		prefix += "s"
	}

	root := strings.TrimSuffix(prefix+"://"+req.Host+api.Context().BasePath(), "/")
	requestUri := strings.TrimSuffix(urlPrefix(req), "/")

	route, _, _ := api.Context().RouteInfo(req)

	return &APIHostBase{
		ServerURL:   root,
		Endpoint:    strings.TrimPrefix(requestUri, root),
		FQEndpoint:  requestUri,
		OperationId: route.Operation.ID,
	}
}

func urlPrefix(req *http.Request) string {
	prefix := "http"
	if req.TLS != nil {
		prefix += "s"
	}
	return prefix + "://" + req.Host + req.RequestURI
}

func buildResourceGroupResponse(rscs []string, parts *APIHostBase) (rg []*models.ResourceGroup) {
	rg = make([]*models.ResourceGroup, len(rscs))
	for i, rsc := range rscs {
		rg[i] = &models.ResourceGroup{
			Name:  str(rsc),
			Links: HALSelfLink(parts.FQEndpoint + "/" + rsc),
		}
	}
	return
}
