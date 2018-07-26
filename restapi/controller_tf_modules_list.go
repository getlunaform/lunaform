package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/go-openapi/runtime/middleware"

	"github.com/getlunaform/lunaform/helpers"
	operations "github.com/getlunaform/lunaform/restapi/operations/modules"
	"strings"
	"net/http"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListTfModulesController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListModulesHandlerFunc {
	return operations.ListModulesHandlerFunc(func(params operations.ListModulesParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		modules, err := buildListTfModules(db, ch)
		if err != nil {
			return err
		}
		
		return operations.NewListModulesOK().WithPayload(&models.ResponseListTfModules{
			Links:    helpers.HalRootRscLinks(ch),
			Embedded: modules,
		})
	})
}

func buildListTfModules(db database.Database, ch helpers.ContextHelper) (m *models.ResourceListTfModule, err middleware.Responder) {
	modules := make([]*models.ResourceTfModule, 0)
	if err := db.List(DB_TABLE_TF_MODULE, &modules); err != nil {
		return nil, NewServerError(http.StatusInternalServerError, err.Error())
	}
	for _, module := range modules {
		module.GenerateLinks(strings.TrimSuffix(ch.Endpoint, "s"))
		module.Embedded = nil
	}

	return &models.ResourceListTfModule{
		Modules: modules,
	}, nil

}
