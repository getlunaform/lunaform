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

		tfm := params.TerraformModule

		newId := uuid.New()
		tfm.VcsID = newId

		err := db.Create("tf-module", newId, tfm)

		if err != nil {
			return tf.NewCreateModuleBadRequest()
		}

		response := &models.ResourceTfModule{
			Links: halSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + tfm.VcsID),
			VcsID: newId,
		}
		response.Links.Doc = halDocLink(ch).Doc

		if tfm == nil {
			return tf.NewCreateModuleBadRequest()
		} else {
			response.Name = tfm.Name
			response.Type = tfm.Type
			response.Source = tfm.Source
			return tf.NewCreateModuleCreated().WithPayload(response)
		}
	})
}

var GetTfModuleController = func(idp identity.Provider, ch ContextHelper, db database.Database) tf.GetModuleHandlerFunc {
	return tf.GetModuleHandlerFunc(func(params tf.GetModuleParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		var module *models.ResourceTfModule
		err := db.Read("tf-module", id, module)

		if err != nil {
			return tf.NewGetModuleInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		} else if module == nil {
			return tf.NewGetModuleNotFound().WithPayload(&models.ServerError{
				StatusCode: Int64(404),
				Status:     String("Not Found"),
				Message:    String("Could not find module with id '" + id + "'"),
			})
		} else {
			module.Links = halSelfLink(ch.FQEndpoint)
			module.Links.Doc = halDocLink(ch).Doc
			return tf.NewGetModuleOK().WithPayload(module)
		}
	})
}
