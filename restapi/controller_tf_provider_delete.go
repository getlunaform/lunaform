package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
)

func DeleteTfProviderController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.DeleteProviderHandlerFunc {
	return func(params operation.DeleteProviderParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		return operation.NewDeleteProviderNoContent()
	}
}
