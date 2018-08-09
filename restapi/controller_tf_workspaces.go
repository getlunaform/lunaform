package restapi

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	operations "github.com/getlunaform/lunaform/restapi/operations/workspaces"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"net/http"
	"strings"
)

var ListTfWorkspacesController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.ListWorkspacesHandlerFunc {
	return operations.ListWorkspacesHandlerFunc(func(params operations.ListWorkspacesParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspaces := []*models.ResourceTfWorkspace{}
		if err := db.List(DB_TABLE_TF_WORKSPACE, &workspaces); err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
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

var CreateTfWorkspaceController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.CreateWorkspaceHandlerFunc {
	return operations.CreateWorkspaceHandlerFunc(func(params operations.CreateWorkspaceParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfw := params.TerraformWorkspace
		tfw.Name = swag.String(params.Name)
		tfw.Modules = []*models.ResourceTfModule{}

		existingWorkspace := models.ResourceTfWorkspace{}

		if err := db.Read(DB_TABLE_TF_WORKSPACE, params.Name, &existingWorkspace); err != nil {
			if _, isNotFound := err.(database.RecordDoesNotExistError); isNotFound {
				if err := db.Create(DB_TABLE_TF_WORKSPACE, params.Name, tfw); err == nil {
					r = operations.NewCreateWorkspaceCreated().WithPayload(tfw)
				} else {
					r = NewServerErrorResponse(http.StatusBadRequest, err.Error())
				}
			} else {
				r = NewServerErrorResponse(http.StatusInternalServerError, err.Error())
			}
		} else {
			if err := db.Update(DB_TABLE_TF_WORKSPACE, params.Name, existingWorkspace); err != nil {
				r = NewServerErrorResponse(http.StatusInternalServerError, err.Error())
			} else {
				r = operations.NewCreateWorkspaceOK().WithPayload(tfw)
			}
		}

		return
	})
}

var GetTfWorkspaceController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.DescribeWorkspaceHandlerFunc {
	return operations.DescribeWorkspaceHandlerFunc(func(params operations.DescribeWorkspaceParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspace := &models.ResourceTfWorkspace{}
		if err := db.Read(DB_TABLE_TF_WORKSPACE, params.Name, workspace); err != nil {
			r = NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		} else if workspace == nil {
			r = NewServerErrorResponse(http.StatusNotFound, "Could not find workspace with name '"+params.Name+"'")
		} else {
			workspace.Links = helpers.HalRootRscLinks(ch)
			r = operations.NewDescribeWorkspaceOK().WithPayload(workspace)
		}
		return
	})
}
