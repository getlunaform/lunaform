package restapi

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func GetTfProviderConfigurationController(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operation.GetProviderConfigurationHandlerFunc {
	return func(params operation.GetProviderConfigurationParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		provider, errCode, err := buildGetTfProviderConfigurationResponse(db, params.ID)
		if err != nil {
			return NewServerErrorResponse(errCode, err.Error())
		}
		provider.Links = helpers.HalAddCuries(ch, helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			ch.Endpoint,
		))

		return operation.NewGetProviderAccepted().WithPayload(provider)
	}
}

func buildGetTfProviderConfigurationResponse(db database.Database, configId string) (prov *models.ResourceTfProvider, errCode int, err error) {
	prov = &models.ResourceTfProvider{}
	if err := db.Read(DB_TABLE_TF_PROVIDER_CONFIGURATION, configId, prov); err != nil {
		prov = nil
		if _, notFound := err.(database.RecordDoesNotExistError); notFound {
			errCode = http.StatusNotFound
		} else {
			errCode = http.StatusInternalServerError
		}
	}
	return
}
