package controller

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/zeebox/terraform-server/backend/identity"
	"github.com/zeebox/terraform-server/server/models"
	"github.com/zeebox/terraform-server/server/restapi/operations/resources"
)

// ListIdentityResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListResourcesController = func(idp identity.Provider, oh ContextHelper) resources.ListResourcesHandlerFunc {
	return resources.ListResourcesHandlerFunc(func(params resources.ListResourcesParams) middleware.Responder {
		oh.SetRequest(params.HTTPRequest)

		var rsc []string
		switch params.Group {
		case "tf":
			rsc = []string{"modules", "stacks", "state-backends", "workspaces"}
		case "identity":
			rsc = []string{"groups", "providers", "users"}
		case "git":
			rsc = []string{"git"}
		}

		r := resources.NewListResourcesOK()
		r.SetPayload(&models.ResponseListResources{
			Links:    halRootRscLinks(oh),
			Embedded: buildResourceGroupResponse(rsc, oh),
		})

		return r
	})
}

// ListResourceGroupsController provides a list of resource groups. This is an exploratory read-only endpoint.
var ListResourceGroupsController = func(idp identity.Provider, oh ContextHelper, ctx *middleware.Context) resources.ListResourceGroupsHandlerFunc {
	return resources.ListResourceGroupsHandlerFunc(func(params resources.ListResourceGroupsParams) middleware.Responder {
		parts := apiParts(params.HTTPRequest, ctx.BasePath())
		parts.OperationID = oh.GetOperationID(params.HTTPRequest, ctx)

		rg := buildResourceGroupResponse([]string{"tf", "identity", "vcs"}, parts)

		r := resources.NewListResourceGroupsOK()
		r.SetPayload(&models.ResponseListResources{
			Links:    halRootRscLinks(parts),
			Embedded: rg,
		})

		return r
	})
}
