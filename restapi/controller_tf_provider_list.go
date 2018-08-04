package restapi

import (
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
)

func ListTfProvidersController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.ListProvidersHandlerFunc {
	return func(params operation.ListProvidersParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		return operation.NewListProvidersOK().WithPayload(&models.ResponseListTfProviders{
			Embedded: &models.ResourceListTfProvider{
				Providers: buildListTfProvidersResponse(),
			},
			Links: helpers.HalAddCuries(ch, helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)),
		})
	}
}

func buildListTfProvidersResponse() []*models.ResourceTfProvider {
	return make([]*models.ResourceTfProvider, 0)
}
