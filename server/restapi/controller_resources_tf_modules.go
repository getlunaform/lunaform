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

		var modules []*models.ResourceTfModule
		if err := db.List("tf-module", &modules); err != nil {
			return operations.NewListModulesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		}

		return operations.NewListModulesOK().WithPayload(&models.ResponseListTfModules{
			Links: models.HalRoot(ch),
			Embedded: &models.ResourceListTfModule{
				Resources: modules,
			},
		})
	})
}

var CreateTfModuleController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.CreateModuleHandlerFunc {
	return operations.CreateModuleHandlerFunc(func(params operations.CreateModuleParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfm := params.TerraformModule
		tfm.ID = uuid.New()

		if err := db.Create("tf-module", tfm.ID, tfm); err != nil {
			return operations.NewCreateModuleBadRequest()
		}

		if tfm == nil {
			return operations.NewCreateModuleBadRequest()
		} else {
			tfm.Links = models.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + tfm.ID)
			tfm.Links.Doc = models.HalDocLink(ch).Doc
			return operations.NewCreateModuleCreated().WithPayload(tfm)
		}
	})
}

var GetTfModuleController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.GetModuleHandlerFunc {
	return operations.GetModuleHandlerFunc(func(params operations.GetModuleParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		var module *models.ResourceTfModule
		if err := db.Read("tf-module", params.ID, module); err != nil {
			return operations.NewGetModuleInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		} else if module == nil {
			return operations.NewGetModuleNotFound().WithPayload(&models.ServerError{
				StatusCode: Int64(404),
				Status:     String("Not Found"),
				Message:    String("Could not find module with id '" + params.ID + "'"),
			})
		} else {
			module.Links = models.HalSelfLink(ch.FQEndpoint)
			module.Links.Doc = models.HalDocLink(ch).Doc
			return operations.NewGetModuleOK().WithPayload(module)
		}
	})
}
