package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/restapi/operations/resources"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"net/http"
)

const (
	DB_TABLE_TF_WORKSPACE    = "lf-workspace"
	DB_TABLE_TF_MODULE       = "lf-module"
	DB_TABLE_TF_STACK        = "lf-stack"
	DB_TABLE_TF_STATEBACKEND = "lf-statebackend"
	DB_TABLE_AUTH_USER       = "lf-auth-user"
)

var (
	// 400
	HTTP_BAD_REQUEST        = swag.Int64(http.StatusBadRequest)
	HTTP_BAD_REQUEST_STATUS = swag.String(http.StatusText(http.StatusBadRequest))
	// 404
	HTTP_NOT_FOUND        = swag.Int64(http.StatusNotFound)
	HTTP_NOT_FOUND_STATUS = swag.String(http.StatusText(http.StatusNotFound))
	// 422
	HTTP_UNPROCESSABLE_ENTITY        = swag.Int64(http.StatusUnprocessableEntity)
	HTTP_UNPROCESSABLE_ENTITY_STATUS = swag.String(http.StatusText(http.StatusUnprocessableEntity))
	// 500
	HTTP_INTERNAL_SERVER_ERROR        = swag.Int64(http.StatusInternalServerError)
	HTTP_INTERNAL_SERVER_ERROR_STATUS = swag.String(http.StatusText(http.StatusInternalServerError))
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListResourcesController = func(ch helpers.ContextHelper) resources.ListResourcesHandlerFunc {
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
			return resources.NewListResourcesOK().WithPayload(&models.ResponseListResources{
				Links:    helpers.HalRootRscLinks(ch),
				Embedded: buildResourceGroupResponse(rsc, ch),
			})
		}

		return resources.NewListResourceGroupsNotFound()
	})
}

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
