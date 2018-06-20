package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/terraform-server/backend/identity"
	"github.com/drewsonne/terraform-server/server/restapi/operations/resources"
	"github.com/drewsonne/terraform-server/server/models"
)

//func(ListTfResourcesParams) middleware.Responder
// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListResourcesTfModuleController = func(idp identity.Provider, ch ContextHelper) resources.ListTfModulesHandlerFunc {
	return resources.ListTfModulesHandlerFunc(func(params resources.ListTfModulesParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tf := resources.NewListTfModulesOK()
		tf.SetPayload(&models.ResponseListTfModules{
			Links: halRootRscLinks(ch),
		})
		return tf

		//return resources.NewListResourceGroupsNotFound()
	})
}
