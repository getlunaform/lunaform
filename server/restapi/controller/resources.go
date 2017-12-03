package controller

import (
	"github.com/zeebox/terraform-server/server/restapi/operations"
	"github.com/zeebox/terraform-server/backend"
	"github.com/zeebox/terraform-server/server/restapi/operations/resources"
	"github.com/go-openapi/runtime/middleware"
	"github.com/zeebox/terraform-server/server/models"
)

var ListResourceGroupsController = func(api *operations.TerraformServerAPI, idp backend.IdentityProvider) resources.ListResourceGroupsHandlerFunc {
	return resources.ListResourceGroupsHandlerFunc(func(params resources.ListResourceGroupsParams) middleware.Responder {

		parts := apiParts(params.HTTPRequest, api)

		rg := buildResourceGroupResponse([]string{"tf", "identity", "vcs"}, parts)

		r := resources.NewListResourceGroupsOK()
		r.SetPayload(&models.ResponseListResourceGroups{
			Links:    HALRootRscLinks(parts),
			Embedded: &models.ResponseListResourceGroupsEmbedded{IdentityResources: rg},
		})

		return r
	})
}

var ListTerraformResourcesController = func(api *operations.TerraformServerAPI, idp backend.IdentityProvider) resources.ListTerraformResourcesHandlerFunc {
	return resources.ListTerraformResourcesHandlerFunc(func(params resources.ListTerraformResourcesParams) middleware.Responder {
		parts := apiParts(params.HTTPRequest, api)

		rg := buildResourceGroupResponse([]string{"modules", "stacks", "state-backends", "workspaces"}, parts)

		r := resources.NewListResourceGroupsOK()
		r.SetPayload(&models.ResponseListResourceGroups{
			Links:    HALRootRscLinks(parts),
			Embedded: &models.ResponseListResourceGroupsEmbedded{IdentityResources: rg},
		})

		return r
	})
}
