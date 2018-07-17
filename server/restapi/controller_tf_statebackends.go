package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/server/helpers"
	"github.com/drewsonne/lunaform/backend/database"
	"github.com/drewsonne/lunaform/server/models"
	"github.com/go-openapi/runtime/middleware"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/state_backends"
)

var CreateTfStateBackendsController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.CreateStateBackendHandlerFunc {
	return operations.CreateStateBackendHandlerFunc(func(params operations.CreateStateBackendParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		return nil
	})
}

var UpdateTfStateBackendsController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.UpdateStateBackendHandlerFunc {
	return operations.UpdateStateBackendHandlerFunc(func(params operations.UpdateStateBackendParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		return nil
	})
}
