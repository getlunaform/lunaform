package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/getlunaform/lunaform/models"
)

func CreateTfProviderController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.CreateProviderHandlerFunc {
	return func(params operation.CreateProviderParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		provider := buildCreateTfProviderResponse(params.TerraformProvider)
		provider.Links = helpers.HalAddCuries(ch, helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			ch.Endpoint,
		))
		return operation.NewCreateProviderCreated().WithPayload(provider)
	}
}

func buildCreateTfProviderResponse(provider *models.ResourceTfProvider) *models.ResourceTfProvider {
	provider.Embedded = &models.ResourceListTfProvider{
		Providers: make([]*models.ResourceTfProvider, 0),
	}
	provider.ID = idGenerator.MustGenerate()

	return provider
}
