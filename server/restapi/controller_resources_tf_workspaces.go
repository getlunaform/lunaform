package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/backend/database"
	"github.com/go-openapi/runtime/middleware"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/workspaces"
	"github.com/drewsonne/lunaform/server/models"
	"strings"
	"github.com/drewsonne/lunaform/server/helpers"
)

var ListTfWorkspacesController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListWorkspacesHandlerFunc {
	return operations.ListWorkspacesHandlerFunc(func(params operations.ListWorkspacesParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspaces := []*models.ResourceTfWorkspace{}
		err := db.List("tf-workspace", &workspaces)
		if err != nil {
			return operations.NewListWorkspacesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: helpers.Int64(500),
				Status:     helpers.String("Internal Server Error"),
				Message:    helpers.String(err.Error()),
			})
		}

		for _, workspace := range workspaces {
			workspace.Links = helpers.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + *workspace.Name)
			if len(workspace.Modules) == 0 {
				workspace.Modules = []*models.ResourceTfModule{}
			}
		}

		return operations.NewListWorkspacesOK().WithPayload(&models.ResponseListTfWorkspaces{
			Links: helpers.HalRootRscLinks(ch),
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

		if err := db.Create("tf-workspace", *tfw.Name, tfw); err != nil {
			return operations.NewCreateWorkspaceBadRequest()
		}

		if tfw == nil {
			return operations.NewCreateWorkspaceBadRequest()
		} else {
			tfw.Links = helpers.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + *tfw.Name)
			tfw.Links.Doc = helpers.HalDocLink(ch).Doc
			return operations.NewCreateWorkspaceCreated().WithPayload(tfw)
		}
	})
}

var GetTfWorkspaceController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.DescribeWorkspaceHandlerFunc {
	return operations.DescribeWorkspaceHandlerFunc(func(params operations.DescribeWorkspaceParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		var module *models.ResourceTfWorkspace
		if err := db.Read("tf-workspace", params.Name, module); err != nil {
			return operations.NewDescribeWorkspaceInternalServerError().WithPayload(&models.ServerError{
				StatusCode: helpers.Int64(500),
				Status:     helpers.String("Internal Server Error"),
				Message:    helpers.String(err.Error()),
			})
		} else if module == nil {
			return operations.NewDescribeWorkspaceNotFound().WithPayload(&models.ServerError{
				StatusCode: helpers.Int64(404),
				Status:     helpers.String("Not Found"),
				Message:    helpers.String("Could not find workspace with name '" + params.Name + "'"),
			})
		} else {
			module.Links = helpers.HalSelfLink(ch.FQEndpoint)
			module.Links.Doc = helpers.HalDocLink(ch).Doc
			return operations.NewDescribeWorkspaceOK().WithPayload(module)
		}
	})
}
