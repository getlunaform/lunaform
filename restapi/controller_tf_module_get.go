package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/go-openapi/runtime/middleware"

	"github.com/getlunaform/lunaform/helpers"
	operations "github.com/getlunaform/lunaform/restapi/operations/modules"
	"net/http"
)

func GetTfModuleController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.GetModuleHandlerFunc {
	return operations.GetModuleHandlerFunc(func(params operations.GetModuleParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		module := &models.ResourceTfModule{}
		if err := buildGetTfModuleResponse(params.ID, module, db, ch); err != nil {
			return err
		}

		return operations.NewGetModuleOK().WithPayload(module)
	})
}

func buildGetTfModuleResponse(moduleId string, module *models.ResourceTfModule, db database.Database, ch helpers.ContextHelper) *CommonServerErrorResponder {
	if err := db.Read(DB_TABLE_TF_MODULE, moduleId, module); err != nil {
		return NewServerError(http.StatusInternalServerError, err.Error())
	} else if module == nil {
		return NewServerError(http.StatusNotFound, "Could not find module with id '"+moduleId+"'")
	} else {

		module.Embedded = &models.ResourceListTfStack{
			Stacks: make([]*models.ResourceTfStack, 0),
		}

		module.Links = helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			ch.Endpoint,
		)
	}
	return nil
}
