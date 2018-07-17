package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/server/helpers"
	"github.com/drewsonne/lunaform/backend/database"
	"github.com/drewsonne/lunaform/server/models"
	"github.com/go-openapi/runtime/middleware"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/state_backends"
	"github.com/pborman/uuid"
	"strings"
)

var CreateTfStateBackendsController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.CreateStateBackendHandlerFunc {
	return operations.CreateStateBackendHandlerFunc(func(params operations.CreateStateBackendParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		statebackend := params.TerraformStateBackend
		statebackend.ID = uuid.New()

		if err := db.Create(DB_TABLE_TF_STATEBACKEND, statebackend.ID, statebackend); err != nil {
			return operations.NewCreateStateBackendBadRequest().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    helpers.String(err.Error()),
			})
		}

		statebackend.Links = helpers.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + statebackend.ID)
		statebackend.Links.Doc = helpers.HalDocLink(ch).Doc
		return operations.NewCreateStateBackendCreated().WithPayload(statebackend)

		return nil
	})
}

var ListTfStateBackendsController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListStateBackendsHandlerFunc {
	return operations.ListStateBackendsHandlerFunc(func(params operations.ListStateBackendsParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		statebackends := []*models.ResourceTfStateBackend{}
		if err := db.List(DB_TABLE_TF_STATEBACKEND, &statebackends); err != nil {
			return operations.NewListStateBackendsInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    helpers.String(err.Error()),
			})
		}

		return operations.NewListStateBackendsOK().WithPayload(&models.ResponseListTfStateBackends{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfStateBackend{
				Modules: statebackends,
			},
		})

		return nil
	})
}

var UpdateTfStateBackendsController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.UpdateStateBackendHandlerFunc {
	return operations.UpdateStateBackendHandlerFunc(func(params operations.UpdateStateBackendParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		return nil
	})
}
