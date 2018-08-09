package restapi

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	operations "github.com/getlunaform/lunaform/restapi/operations/state_backends"
	"github.com/go-openapi/runtime/middleware"

	"net/http"
	"strings"
)

var CreateTfStateBackendsController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.CreateStateBackendHandlerFunc {
	return operations.CreateStateBackendHandlerFunc(func(params operations.CreateStateBackendParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		statebackend := params.TerraformStateBackend
		statebackend.ID = idGenerator.MustGenerate()

		if err := db.Create(DB_TABLE_TF_STATEBACKEND, statebackend.ID, statebackend); err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		}

		statebackend.Links = helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			strings.TrimSuffix(ch.Endpoint, "s")+"/"+statebackend.ID,
		)
		return operations.NewCreateStateBackendCreated().WithPayload(statebackend)
	})
}

var ListTfStateBackendsController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.ListStateBackendsHandlerFunc {
	return operations.ListStateBackendsHandlerFunc(func(params operations.ListStateBackendsParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		statebackends := []*models.ResourceTfStateBackend{}
		if err := db.List(DB_TABLE_TF_STATEBACKEND, &statebackends); err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		}

		return operations.NewListStateBackendsOK().WithPayload(&models.ResponseListTfStateBackends{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfStateBackend{
				StateBackends: statebackends,
			},
		})
	})
}

var UpdateTfStateBackendsController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.UpdateStateBackendHandlerFunc {
	return operations.UpdateStateBackendHandlerFunc(func(params operations.UpdateStateBackendParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		statebackend := &models.ResourceTfStateBackend{}
		if err := db.Read(DB_TABLE_TF_STATEBACKEND, params.ID, statebackend); err != nil {
			if _, notFound := err.(database.RecordDoesNotExistError); notFound {
				return NewServerErrorResponse(http.StatusNotFound, err.Error())
			} else {
				return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
			}
		}

		if params.TerraformStateBackend.Name != "" {
			statebackend.Name = params.TerraformStateBackend.Name
		}

		if params.TerraformStateBackend.Configuration != nil {
			statebackend.Configuration = params.TerraformStateBackend.Configuration
		}

		if err := db.Update(DB_TABLE_TF_STATEBACKEND, params.ID, statebackend); err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		}

		return operations.NewUpdateStateBackendOK().WithPayload(statebackend)
	})
}
