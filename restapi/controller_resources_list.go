package restapi

import (
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/restapi/operations/resources"
	"github.com/go-openapi/runtime/middleware"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListResourcesController = func(ch *helpers.ContextHelper) resources.ListResourcesHandlerFunc {
	return resources.ListResourcesHandlerFunc(func(params resources.ListResourcesParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		return resources.NewListResourcesOK().WithPayload(&models.ResponseListResources{
			Links:    helpers.HalRootRscLinks(ch),
			Embedded: buildResourceGroupRootResponse(params.Group, ch),
		})
	})
}

func buildResourceGroupRootResponse(group string, ch *helpers.ContextHelper) (rsclist *models.ResourceList) {
	var rsc []string
	switch group {
	case "tf":
		rsc = []string{"modules", "stacks", "state-backends", "workspaces", "providers"}
	case "identity":
		rsc = []string{"groups", "providers", "users"}
	case "vcs":
		rsc = []string{"git"}
	}

	return buildResourceGroupResponse(rsc, ch)
}
