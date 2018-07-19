package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/server/backend/database"
	"github.com/getlunaform/lunaform/server/backend/identity"
	"github.com/getlunaform/lunaform/server/helpers"
	operations "github.com/getlunaform/lunaform/server/restapi/operations/workspaces"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"strings"
)

var ListTfWorkspacesController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListWorkspacesHandlerFunc {
	return operations.ListWorkspacesHandlerFunc(func(params operations.ListWorkspacesParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspaces := []*models.ResourceTfWorkspace{}
		err := db.List(DB_TABLE_TF_WORKSPACE, &workspaces)
		if err != nil {
			return operations.NewListWorkspacesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		}

		for _, workspace := range workspaces {
			workspace.Links = helpers.HalSelfLink(nil, strings.TrimSuffix(ch.Endpoint, "s")+"/"+*workspace.Name)
			workspace.Links.Curies = nil
		}

		return operations.NewListWorkspacesOK().WithPayload(&models.ResponseListTfWorkspaces{
			Links: helpers.HalAddCuries(ch, helpers.HalRootRscLinks(ch)),
			Embedded: &models.ResourceListTfWorkspace{
				Workspaces: workspaces,
			},
		})
	})
}

var CreateTfWorkspaceController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.CreateWorkspaceHandlerFunc {
	return operations.CreateWorkspaceHandlerFunc(func(params operations.CreateWorkspaceParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfw := params.TerraformWorkspace
		tfw.Name = swag.String(params.Name)
		tfw.Modules = []*models.ResourceTfModule{}

		existingWorkspace := models.ResourceTfWorkspace{}

		//halRscLinks := helpers.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + params.Name)
		//halRscLinks.Doc = helpers.HalDocLink(ch).Doc

		if err := db.Read(DB_TABLE_TF_WORKSPACE, params.Name, &existingWorkspace); err != nil {
			if _, isNotFound := err.(database.RecordDoesNotExistError); isNotFound {
				if err := db.Create(DB_TABLE_TF_WORKSPACE, params.Name, tfw); err == nil {
					r = operations.NewCreateWorkspaceCreated().WithPayload(tfw)
				} else {
					r = operations.NewCreateWorkspaceBadRequest().WithPayload(&models.ServerError{
						StatusCode: HTTP_BAD_REQUEST,
						Status:     HTTP_BAD_REQUEST_STATUS,
						Message:    swag.String(err.Error()),
					})
				}
			} else {
				r = operations.NewCreateWorkspaceInternalServerError().WithPayload(&models.ServerError{
					StatusCode: HTTP_INTERNAL_SERVER_ERROR,
					Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
					Message:    swag.String(err.Error()),
				})
			}
		} else {
			if err := db.Update(DB_TABLE_TF_WORKSPACE, params.Name, existingWorkspace); err != nil {
				r = operations.NewCreateWorkspaceInternalServerError().WithPayload(&models.ServerError{
					StatusCode: HTTP_INTERNAL_SERVER_ERROR,
					Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
					Message:    swag.String(err.Error()),
				})
			} else {
				r = operations.NewCreateWorkspaceOK().WithPayload(tfw)
			}
		}

		return
	})
}

var GetTfWorkspaceController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.DescribeWorkspaceHandlerFunc {
	return operations.DescribeWorkspaceHandlerFunc(func(params operations.DescribeWorkspaceParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspace := &models.ResourceTfWorkspace{}
		if err := db.Read(DB_TABLE_TF_WORKSPACE, params.Name, workspace); err != nil {
			r = operations.NewDescribeWorkspaceInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		} else if workspace == nil {
			r = operations.NewDescribeWorkspaceNotFound().WithPayload(&models.ServerError{
				StatusCode: HTTP_NOT_FOUND,
				Status:     HTTP_NOT_FOUND_STATUS,
				Message:    swag.String("Could not find workspace with name '" + params.Name + "'"),
			})
		} else {
			workspace.Links = helpers.HalRootRscLinks(ch)
			r = operations.NewDescribeWorkspaceOK().WithPayload(workspace)
		}
		return
	})
}
