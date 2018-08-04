package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
)

func GetTfProviderConfigurationController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.GetProviderConfigurationHandlerFunc {
	return func(params operation.GetProviderConfigurationParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		provider := buildGetTfProviderConfigurationResponse()
		provider.Links = helpers.HalAddCuries(ch, helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			ch.Endpoint,
		))

		return operation.NewGetProviderAccepted().WithPayload(provider)
	}
}

func buildGetTfProviderConfigurationResponse() *models.ResourceTfProvider {
	return &models.ResourceTfProvider{}
}
