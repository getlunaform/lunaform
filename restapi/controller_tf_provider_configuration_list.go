package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"github.com/go-openapi/swag"
	"strings"
)

func ListTfProviderConfigurationController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.ListProviderConfigurationsHandlerFunc {
	return func(params operation.ListProviderConfigurationsParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)
		providers, code, err := buildListTfProviderConfigurationResponse(db)
		if err != nil {
			return NewServerError(code, err.Error())
		}

		return operation.NewListProviderConfigurationsOK().WithPayload(&models.ResponseListTfProviderConfiguration{
			Embedded: &models.ResourceListTfProviderConfiguration{
				ProviderConfigurations: providers,
				Provider: &models.ResourceTfProvider{
					Name: swag.String(params.ProviderName),
					Links: helpers.HalSelfLink(nil,
						strings.TrimSuffix(ch.EndpointSingular, "/configuration")),
				},
			},
			Links: helpers.HalAddCuries(ch, helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)),
		})

	}
}
func buildListTfProviderConfigurationResponse(db database.Database) (
	confs []*models.ResourceTfProviderConfiguration, errCode int, err error,
) {
	confs = make([]*models.ResourceTfProviderConfiguration, 0)
	if err = db.List(DB_TABLE_TF_PROVIDER_CONFIGURATION, &confs); err != nil {
		if _, notFound := err.(database.RecordDoesNotExistError); notFound {
			errCode = http.StatusNotFound
		} else {
			errCode = http.StatusInternalServerError
		}
		confs = nil
	}
	return
}
