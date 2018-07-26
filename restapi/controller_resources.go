package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/restapi/operations/resources"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

const (
	DB_TABLE_TF_WORKSPACE    = "lf-workspace"
	DB_TABLE_TF_MODULE       = "lf-module"
	DB_TABLE_TF_STACK        = "lf-stack"
	DB_TABLE_TF_STATEBACKEND = "lf-statebackend"
	DB_TABLE_AUTH_USER       = "lf-auth-user"
)

// ListResourceGroupsController provides a list of resource groups. This is an exploratory read-only endpoint.
var ListResourceGroupsController = func(ch helpers.ContextHelper) resources.ListResourceGroupsHandlerFunc {
	return resources.ListResourceGroupsHandlerFunc(func(params resources.ListResourceGroupsParams) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		rg := buildResourceGroupResponse([]string{"tf", "identity", "vcs"}, ch)

		r := resources.NewListResourceGroupsOK()
		r.SetPayload(&models.ResponseListResources{
			Links:    helpers.HalRootRscLinks(ch),
			Embedded: rg,
		})

		return r
	})
}

func buildResourceGroupResponse(rscs []string, ch helpers.ContextHelper) (rsclist *models.ResourceList) {
	rsclist = &models.ResourceList{
		Resources: make([]*models.Resource, len(rscs)),
	}
	for i, rsc := range rscs {
		rsclist.Resources[i] = &models.Resource{
			Name:  swag.String(rsc),
			Links: helpers.HalSelfLink(nil, ch.Endpoint+"/"+rsc),
		}
	}
	return
}
