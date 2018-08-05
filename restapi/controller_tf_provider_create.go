package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/getlunaform/lunaform/models"
	"net/http"
)

func CreateTfProviderController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.CreateProviderHandlerFunc {
	return func(params operation.CreateProviderParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		if code, err := buildCreateTfProviderResponse(params.TerraformProvider, db, ch); err != nil {
			return NewServerError(code, err.Error())
		}
		return operation.NewCreateProviderCreated().WithPayload(params.TerraformProvider)
	}
}

func buildCreateTfProviderResponse(provider *models.ResourceTfProvider, db database.Database, ch helpers.ContextHelper) (errCode int, err error) {
	provider.Embedded = &models.ResourceListTfStack{
		Stacks: make([]*models.ResourceTfStack, 0),
	}

	if err := db.Create(DB_TABLE_TF_PROVIDER, provider.Name, provider); err != nil {
		return http.StatusInternalServerError, err
	}
	provider.Links = helpers.HalAddCuries(ch, helpers.HalSelfLink(
		helpers.HalDocLink(nil, ch.OperationID),
		ch.Endpoint,
	))

	return
}
