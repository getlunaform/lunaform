package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
)

func GetTfProviderController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.GetProviderHandlerFunc {
	return func(params operation.GetProviderParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		return operation.NewGetProviderAccepted().WithPayload(&models.ResourceTfProvider{
			Links: helpers.HalAddCuries(ch, helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)),
		})
	}
}
