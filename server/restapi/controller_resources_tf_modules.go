package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/server/models"
	"github.com/drewsonne/lunaform/backend/database"
	"github.com/pborman/uuid"
	"strings"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/modules"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListTfModulesController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.ListModulesHandlerFunc {
	return operations.ListModulesHandlerFunc(func(params operations.ListModulesParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		modules := []*models.ResourceTfModule{}
		err := db.List("tf-module", &modules)
		if err != nil {
			return operations.NewListModulesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		}

		return operations.NewListModulesOK().WithPayload(&models.ResponseListTfModules{
			Links: halRootRscLinks(ch),
			Embedded: &models.ResourceListTfModule{
				Resources: modules,
			},
		})
	})
}

var CreateTfModuleController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.CreateModuleHandlerFunc {
	return operations.CreateModuleHandlerFunc(func(params operations.CreateModuleParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfm := params.TerraformModule

		tfm.VcsID = uuid.New()

		err := db.Create("tf-module", tfm.VcsID, tfm)

		if err != nil {
			return operations.NewCreateModuleBadRequest()
		}

		response := &models.ResourceTfModule{
			Links: halSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + tfm.VcsID),
			VcsID: tfm.VcsID,
		}
		response.Links.Doc = halDocLink(ch).Doc

		if tfm == nil {
			return operations.NewCreateModuleBadRequest()
		} else {
			response.Name = tfm.Name
			response.Type = tfm.Type
			response.Source = tfm.Source
			return operations.NewCreateModuleCreated().WithPayload(response)
		}
	})
}

var GetTfModuleController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.GetModuleHandlerFunc {
	return operations.GetModuleHandlerFunc(func(params operations.GetModuleParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		var module *models.ResourceTfModule
		err := db.Read("tf-module", id, module)

		if err != nil {
			return operations.NewGetModuleInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		} else if module == nil {
			return operations.NewGetModuleNotFound().WithPayload(&models.ServerError{
				StatusCode: Int64(404),
				Status:     String("Not Found"),
				Message:    String("Could not find module with id '" + id + "'"),
			})
		} else {
			module.Links = halSelfLink(ch.FQEndpoint)
			module.Links.Doc = halDocLink(ch).Doc
			return operations.NewGetModuleOK().WithPayload(module)
		}
	})
}
