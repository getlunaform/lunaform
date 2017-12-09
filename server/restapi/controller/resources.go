package controller

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/zeebox/terraform-server/backend/identity"
	"github.com/zeebox/terraform-server/server/models"
	"github.com/zeebox/terraform-server/server/restapi/operations"
	"github.com/zeebox/terraform-server/server/restapi/operations/resources"
)

var ListIdentityResourcesController = func(api *operations.TerraformServerAPI, idp identity.Provider) resources.ListIdentityResourcesHandlerFunc {
	return resources.ListIdentityResourcesHandlerFunc(func(params resources.ListIdentityResourcesParams) middleware.Responder {
		parts := apiParts(params.HTTPRequest, api)

		ir := buildResourceGroupResponse([]string{"groups", "providers", "users"}, parts)

		r := resources.NewListIdentityResourcesOK()
		r.SetPayload(&models.ResponseListIdentityResources{
			Links:    HALRootRscLinks(parts),
			Embedded: &models.ResponseListIdentityResourcesEmbedded{IdentityResources: ir},
		})

		return r
	})
}

var ListResourceGroupsController = func(api *operations.TerraformServerAPI, idp identity.Provider) resources.ListResourceGroupsHandlerFunc {
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

var ListTerraformResourcesController = func(api *operations.TerraformServerAPI, idp identity.Provider) resources.ListTerraformResourcesHandlerFunc {
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
