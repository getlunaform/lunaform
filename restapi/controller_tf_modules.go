package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/go-openapi/runtime/middleware"

	"github.com/getlunaform/lunaform/helpers"
	operations "github.com/getlunaform/lunaform/restapi/operations/modules"
	"strings"
	"fmt"
	"net/http"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListTfModulesController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListModulesHandlerFunc {
	return operations.ListModulesHandlerFunc(func(params operations.ListModulesParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		modules := make([]*models.ResourceTfModule, 0)
		if err := db.List(DB_TABLE_TF_MODULE, &modules); err != nil {
			return NewServerError(http.StatusInternalServerError, err.Error())
		}

		for _, module := range modules {
			module.GenerateLinks(strings.TrimSuffix(ch.Endpoint, "s"))
			module.Embedded = nil
		}

		return operations.NewListModulesOK().WithPayload(&models.ResponseListTfModules{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfModule{
				Modules: modules,
			},
		})
	})
}

var CreateTfModuleController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.CreateModuleHandlerFunc {
	return operations.CreateModuleHandlerFunc(func(params operations.CreateModuleParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfm := params.TerraformModule
		tfm.ID = idGenerator.MustGenerate()

		if err := db.Create(DB_TABLE_TF_MODULE, tfm.ID, tfm); err != nil {
			return NewServerError(http.StatusInternalServerError, err.Error())
		}

		tfm.Links = helpers.HalRootRscLinks(ch)
		tfm.Embedded = nil
		//tfm.Embedded = &models.ResourceListTfStack{
		//	Stacks: make([]*models.ResourceTfStack, 0),
		//}
		return operations.NewCreateModuleCreated().WithPayload(tfm)
	})
}

var GetTfModuleController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.GetModuleHandlerFunc {
	return operations.GetModuleHandlerFunc(func(params operations.GetModuleParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		module := &models.ResourceTfModule{}
		if err := db.Read(DB_TABLE_TF_MODULE, params.ID, module); err != nil {
			return NewServerError(http.StatusInternalServerError, err.Error())
		} else if module == nil {
			return NewServerError(http.StatusNotFound, "Could not find module with id '"+params.ID+"'")
		} else {

			module.Embedded = &models.ResourceListTfStack{
				Stacks: make([]*models.ResourceTfStack, 0),
			}

			module.Links = helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)

			return operations.NewGetModuleOK().WithPayload(module)
		}
	})
}

var DeleteTfModuleController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.DeleteModuleHandlerFunc {
	return operations.DeleteModuleHandlerFunc(func(params operations.DeleteModuleParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		module := &models.ResourceTfModule{}
		if err := db.Read(DB_TABLE_TF_MODULE, params.ID, module); err != nil {
			if _, moduleNotFound := err.(database.RecordDoesNotExistError); moduleNotFound {
				return operations.NewDeleteModuleNoContent()
			} else {
				return NewServerError(http.StatusInternalServerError, err.Error())
			}
		}

		if len(module.Embedded.Stacks) > 0 {
			stack_ids := []string{}
			for _, stk := range module.Embedded.Stacks {
				stack_ids = append(stack_ids, stk.ID)
			}
			return NewServerError(
				http.StatusUnprocessableEntity,
				fmt.Sprintf("Could not delete module as it is relied up by stacks ['%s']", strings.Join(stack_ids, "','")),
			)
		}

		db.Delete(DB_TABLE_TF_MODULE, params.ID)

		return NewServerError(http.StatusInternalServerError, "Could not delete module")
	})
}
