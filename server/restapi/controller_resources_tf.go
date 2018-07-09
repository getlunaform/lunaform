package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/terraform-server/backend/identity"
	"github.com/drewsonne/terraform-server/server/models"
	"github.com/drewsonne/terraform-server/server/restapi/operations/tf"
	"github.com/drewsonne/terraform-server/backend/database"
	"github.com/pborman/uuid"
	"encoding/json"
	"strings"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListTfModulesController = func(idp identity.Provider, ch ContextHelper, db database.Database) tf.ListModulesHandlerFunc {
	return tf.ListModulesHandlerFunc(func(params tf.ListModulesParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		records, err := db.List("tf-module")
		if err != nil {
			var statuscode int64 = 500
			return tf.NewListModulesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: &statuscode,
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		}

		modules := make([]*models.ResourceTfModule, len(records))
		for i, record := range records {
			mod := models.ResourceTfModule{}
			json.Unmarshal([]byte(record.Value), &mod)
			mod.Links = halSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + mod.VcsID)
			modules[i] = &mod
		}

		return tf.NewListModulesOK().WithPayload(&models.ResponseListTfModules{
			Links: halRootRscLinks(ch),
			Embedded: &models.ResourceListTfModule{
				Resources: modules,
			},
		})
	})
}

var CreateTfModuleController = func(idp identity.Provider, ch ContextHelper, db database.Database) tf.CreateModuleHandlerFunc {
	return tf.CreateModuleHandlerFunc(func(params tf.CreateModuleParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		newId := uuid.New()
		params.TerraformModule.VcsID = newId
		db.Create("tf-module", newId, params.TerraformModule)

		response := &models.ResourceTfModule{
			Links: halRootRscLinks(ch),
			VcsID: newId,
		}
		if params.TerraformModule == nil {
			return tf.NewCreateModuleBadRequest()
		} else {
			response.Name = params.TerraformModule.Name
			response.Type = params.TerraformModule.Type
			return tf.NewCreateModuleCreated().WithPayload(response)
		}

	})
}
