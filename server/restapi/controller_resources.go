package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/lunarform/backend/identity"
	"github.com/drewsonne/lunarform/server/models"
	"github.com/drewsonne/lunarform/server/restapi/operations/resources"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListResourcesController = func(idp identity.Provider, ch ContextHelper) resources.ListResourcesHandlerFunc {
	return resources.ListResourcesHandlerFunc(func(params resources.ListResourcesParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		var rsc []string
		switch params.Group {
		case "tf":
			rsc = []string{"modules", "stacks", "state-backends", "workspaces"}
		case "identity":
			rsc = []string{"groups", "providers", "users"}
		case "vcs":
			rsc = []string{"git"}
		}

		if len(rsc) > 0 {
			r := resources.NewListResourcesOK()
			r.SetPayload(&models.ResponseListResources{
				Links:    halRootRscLinks(ch),
				Embedded: buildResourceGroupResponse(rsc, ch),
			})
			return r
		}

		return resources.NewListResourceGroupsNotFound()
	})
}

// ListResourceGroupsController provides a list of resource groups. This is an exploratory read-only endpoint.
var ListResourceGroupsController = func(idp identity.Provider, ch ContextHelper) resources.ListResourceGroupsHandlerFunc {
	return resources.ListResourceGroupsHandlerFunc(func(params resources.ListResourceGroupsParams) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		rg := buildResourceGroupResponse([]string{"tf", "identity", "vcs"}, ch)

		r := resources.NewListResourceGroupsOK()
		r.SetPayload(&models.ResponseListResources{
			Links:    halRootRscLinks(ch),
			Embedded: rg,
		})

		return r
	})
}

func buildResourceGroupResponse(rscs []string, ch ContextHelper) (rsclist *models.ResourceList) {
	rsclist = &models.ResourceList{
		Resources: make([]*models.Resource, len(rscs)),
	}
	for i, rsc := range rscs {
		rsclist.Resources[i] = &models.Resource{
			Name:  str(rsc),
			Links: halSelfLink(ch.FQEndpoint + "/" + rsc),
		}
	}
	return
}
