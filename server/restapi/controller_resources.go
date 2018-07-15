package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/server/models"
	"github.com/drewsonne/lunaform/server/restapi/operations/resources"
	"github.com/drewsonne/lunaform/server/helpers"
	"net/http"
)

const (
	DB_TABLE_TF_WORKSPACE = "lf-workspace"
	DB_TABLE_TF_MODULE    = "lf-module"
	DB_TABLE_TF_STACK     = "lf-stack"
	DB_TABLE_AUTH_USER    = "lf-auth-user"
)

var (
	// 400
	HTTP_BAD_REQUEST        = helpers.Int64(http.StatusBadRequest)
	HTTP_BAD_REQUEST_STATUS = helpers.String(http.StatusText(http.StatusBadRequest))
	// 404
	HTTP_NOT_FOUND        = helpers.Int64(http.StatusNotFound)
	HTTP_NOT_FOUND_STATUS = helpers.String(http.StatusText(http.StatusNotFound))
	// 500
	HTTP_INTERNAL_SERVER_ERROR        = helpers.Int64(http.StatusInternalServerError)
	HTTP_INTERNAL_SERVER_ERROR_STATUS = helpers.String(http.StatusText(http.StatusInternalServerError))
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListResourcesController = func(idp identity.Provider, ch helpers.ContextHelper) resources.ListResourcesHandlerFunc {
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
				Links:    helpers.HalRootRscLinks(ch),
				Embedded: buildResourceGroupResponse(rsc, ch),
			})
			return r
		}

		return resources.NewListResourceGroupsNotFound()
	})
}

// ListResourceGroupsController provides a list of resource groups. This is an exploratory read-only endpoint.
var ListResourceGroupsController = func(idp identity.Provider, ch helpers.ContextHelper) resources.ListResourceGroupsHandlerFunc {
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
			Name:  helpers.String(rsc),
			Links: helpers.HalSelfLink(ch.FQEndpoint + "/" + rsc),
		}
	}
	return
}
