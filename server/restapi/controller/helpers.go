package controller

import (
	"github.com/go-openapi/strfmt"
	"github.com/zeebox/terraform-server/server/models"
	"github.com/zeebox/terraform-server/server/restapi/operations"
	"net/http"
	"strings"
)

func str(v string) *string { return &v }

func halRootRscLinks(parts *apiHostBase) *models.HalRscLinks {
	lnks := halSelfLink(parts.FQEndpoint)
	lnks.Doc = &models.HalHref{
		Href: strfmt.URI(parts.ServerURL + "/docs#operation/" + parts.OperationID),
	}
	return lnks
}
func halSelfLink(href string) *models.HalRscLinks {
	return &models.HalRscLinks{
		Self: &models.HalHref{Href: strfmt.URI(href)},
	}
}

type apiHostBase struct {
	ServerURL   string
	Endpoint    string
	FQEndpoint  string
	OperationID string
}

func apiParts(req *http.Request, api *operations.TerraformServerAPI) *apiHostBase {
	prefix := "http"
	if req.TLS != nil {
		prefix += "s"
	}

	root := strings.TrimSuffix(prefix+"://"+req.Host+api.Context().BasePath(), "/")
	requestURI := strings.TrimSuffix(urlPrefix(req), "/")

	route, _, _ := api.Context().RouteInfo(req)

	return &apiHostBase{
		ServerURL:   root,
		Endpoint:    strings.TrimPrefix(requestURI, root),
		FQEndpoint:  requestURI,
		OperationID: route.Operation.ID,
	}
}

func urlPrefix(req *http.Request) string {
	prefix := "http"
	if req.TLS != nil {
		prefix += "s"
	}
	return prefix + "://" + req.Host + req.RequestURI
}

func buildResourceGroupResponse(rscs []string, parts *apiHostBase) (rg []*models.ResourceGroup) {
	rg = make([]*models.ResourceGroup, len(rscs))
	for i, rsc := range rscs {
		rg[i] = &models.ResourceGroup{
			Name:  str(rsc),
			Links: halSelfLink(parts.FQEndpoint + "/" + rsc),
		}
	}
	return
}
