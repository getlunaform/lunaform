package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/server/helpers"
	"github.com/drewsonne/lunaform/backend/database"
	models "github.com/getlunaform/lunaform-models-go"
	"github.com/go-openapi/runtime/middleware"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/state_backends"

	"strings"
	"github.com/go-openapi/swag"
)

var CreateTfStateBackendsController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.CreateStateBackendHandlerFunc {
	return operations.CreateStateBackendHandlerFunc(func(params operations.CreateStateBackendParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		statebackend := params.TerraformStateBackend
		statebackend.ID = idGenerator.MustGenerate()

		if err := db.Create(DB_TABLE_TF_STATEBACKEND, statebackend.ID, statebackend); err != nil {
			return operations.NewCreateStateBackendBadRequest().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		}

		statebackend.Links = helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			strings.TrimSuffix(ch.Endpoint, "s")+"/"+statebackend.ID,
		)
		return operations.NewCreateStateBackendCreated().WithPayload(statebackend)
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
				Message:    swag.String(err.Error()),
			})
		}

		return operations.NewListStateBackendsOK().WithPayload(&models.ResponseListTfStateBackends{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfStateBackend{
				StateBackends: statebackends,
			},
		})

		return nil
	})
}

var UpdateTfStateBackendsController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.UpdateStateBackendHandlerFunc {
	return operations.UpdateStateBackendHandlerFunc(func(params operations.UpdateStateBackendParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		statebackend := &models.ResourceTfStateBackend{}
		if err := db.Read(DB_TABLE_TF_STATEBACKEND, params.ID, statebackend); err != nil {
			if _, notFound := err.(database.RecordDoesNotExistError); notFound {
				return operations.NewUpdateStateBackendNotFound().WithPayload(&models.ServerError{
					StatusCode: HTTP_NOT_FOUND,
					Status:     HTTP_NOT_FOUND_STATUS,
					Message:    swag.String(err.Error()),
				})
			} else {
				return operations.NewCreateStateBackendBadRequest().WithPayload(&models.ServerError{
					StatusCode: HTTP_INTERNAL_SERVER_ERROR,
					Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
					Message:    swag.String(err.Error()),
				})
			}
		}

		if params.TerraformStateBackend.Name != "" {
			statebackend.Name = params.TerraformStateBackend.Name
		}

		if params.TerraformStateBackend.Configuration != nil {
			statebackend.Configuration = params.TerraformStateBackend.Configuration
		}

		if err := db.Update(DB_TABLE_TF_STATEBACKEND, params.ID, statebackend); err != nil {
			return operations.NewCreateStateBackendBadRequest().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		}

		return operations.NewUpdateStateBackendOK().WithPayload(statebackend)
	})
}
