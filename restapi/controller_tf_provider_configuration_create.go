package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"net/http"
	"github.com/go-openapi/swag"
)

func CreateTfProviderConfigurationController(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operation.CreateProviderConfigurationHandlerFunc {
	return func(params operation.CreateProviderConfigurationParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		if params.ProviderConfiguration == nil {
			return NewServerErrorResponse(
				http.StatusBadRequest,
				"Missing configuration body",
			)
		}

		if errCode, err := buildCreateTfProviderConfigurationResponse(
			params.ProviderConfiguration,
			params.ProviderName,
			db, ch,
		); err != nil {
			return NewServerErrorResponse(errCode, err.Error())
		}

		return operation.NewCreateProviderConfigurationCreated().
			WithPayload(params.ProviderConfiguration)
	}
}

func buildCreateTfProviderConfigurationResponse(config *models.ResourceTfProviderConfiguration, providerName string, db database.Database, ch *helpers.ContextHelper) (errCode int, err error) {
	config.Embedded = &models.ResourceTfProviderConfigurationEmbedded{
		Provider: &models.ResourceTfProvider{
			Name: swag.String(providerName),
		},
	}

	if config.Configuration == nil {
		config.Configuration = make(map[string]interface{}, 0)
	}

	config.ID = idGenerator.MustGenerate()

	if err := db.Create(DB_TABLE_TF_PROVIDER_CONFIGURATION, config.ID, config); err != nil {
		return http.StatusInternalServerError, err
	}
	config.Links = helpers.HalAddCuries(ch, helpers.HalSelfLink(
		helpers.HalDocLink(nil, ch.OperationID),
		ch.Endpoint,
	))

	return
}
