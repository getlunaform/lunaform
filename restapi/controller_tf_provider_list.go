package restapi

import (
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func ListTfProvidersController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.ListProvidersHandlerFunc {
	return func(params operation.ListProvidersParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)
		providers := make([]*models.ResourceTfProvider, 0)
		if code, err := buildListTfProvidersResponse(db, &providers); err != nil {
			return NewServerError(code, err.Error())
		}

		return operation.NewListProvidersOK().WithPayload(&models.ResponseListTfProviders{
			Embedded: &models.ResourceListTfProvider{
				Providers: providers,
			},
			Links: helpers.HalAddCuries(ch, helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)),
		})
	}
}

func buildListTfProvidersResponse(db database.Database, providers *[]*models.ResourceTfProvider) (errCode int, err error) {
	if err := db.List(DB_TABLE_TF_PROVIDER, &providers); err != nil {
		return http.StatusInternalServerError, err
	}
	return
}
