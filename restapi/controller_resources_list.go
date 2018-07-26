package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/restapi/operations/resources"
	"github.com/go-openapi/runtime/middleware"
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
