package controller

import (
	"github.com/go-openapi/strfmt"
	"github.com/zeebox/terraform-server/server/models"
	"net/http"
)

type ContextHelper interface {
	SetRequest(req *http.Request)
	GetOperationID() string
	GetFQEndpoint() string
	GetServerURL() string
}

func str(v string) *string { return &v }

func halRootRscLinks(oh ContextHelper) *models.HalRscLinks {
	lnks := halSelfLink(oh.GetFQEndpoint())
	lnks.Doc = &models.HalHref{
		Href: strfmt.URI(oh.GetServerURL() + "/docs#operation/" + oh.GetOperationID()),
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

func buildResourceGroupResponse(rscs []string, ch ContextHelper) (rsclist *models.ResourceList) {
	rsclist = &models.ResourceList{
		Resources: make([]*models.Resource, len(rscs)),
	}
	for i, rsc := range rscs {
		rsclist.Resources[i] = &models.Resource{
			Name:  str(rsc),
			Links: halSelfLink(ch.GetFQEndpoint() + "/" + rsc),
		}
	}
	return
}
