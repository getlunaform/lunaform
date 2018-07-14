package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/backend/database"
	"github.com/go-openapi/runtime/middleware"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/workspaces"
	"github.com/drewsonne/lunaform/server/models"
	"strings"
)

var ListTfWorkspacesController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.ListWorkspacesHandlerFunc {
	return operations.ListWorkspacesHandlerFunc(func(params operations.ListWorkspacesParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspaces := make([]*models.ResourceListTfWorkspaceWorkspacesItems0, 0)
		err := db.List("tf-workspaces", &workspaces)
		if err != nil {
			return operations.NewListWorkspacesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		}

		return operations.NewListWorkspacesOK().WithPayload(&models.ResponseListTfWorkspaces{
			Links: halRootRscLinks(ch),
			Embedded: &models.ResourceListTfWorkspace{
				Workspaces: workspaces,
			},
		})
	})
}

var CreateTfWorkspaceController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.CreateWorkspaceHandlerFunc {
	return operations.CreateWorkspaceHandlerFunc(func(params operations.CreateWorkspaceParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfw := params.TerraformWorkspace

		if err := db.Create("tf-module", *tfw.Name, tfw); err != nil {
			return operations.NewCreateWorkspaceBadRequest()
		}

		if tfw == nil {
			return operations.NewCreateWorkspaceBadRequest()
		} else {
			tfw.Links = halSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + *tfw.Name)
			tfw.Links.Doc = halDocLink(ch).Doc
			return operations.NewCreateWorkspaceCreated().WithPayload(tfw)
		}
	})
}
