package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/backend/database"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/stacks"
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/lunaform/server/models"
	"strings"
	"github.com/pborman/uuid"
	"github.com/drewsonne/lunaform/server/helpers"
	"fmt"
)

const (
	TF_STACK_STATUS_WAITING_FOR_DEPLOYMENT = "waiting_for_deployment"
	TF_STACK_STATUS_DEPLOY_FAIL            = "deployment_failed"
	TF_STACK_STATUS_DEPLOY_SUCEED          = "deployment_succeeded"
	TF_DEPLOYMENT_STATUS_PENDING           = "pending"
	TF_DEPLOYMENT_STATUS_DEPLOYING         = "deploying"
	TF_DEPLOYMENT_STATUS_SUCCESS           = "finished"
	TF_DEPLOYMENT_STATUS_FAIL              = "failed"
)

var ListTfStacksController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListStacksHandlerFunc {
	return operations.ListStacksHandlerFunc(func(params operations.ListStacksParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		stacks := make([]*models.ResourceTfStack, 0)
		err := db.List(DB_TABLE_TF_STACK, &stacks)
		if err != nil {
			return operations.NewListStacksInternalServerError().WithPayload(&models.ServerError{
				StatusCode: helpers.Int64(500),
				Status:     helpers.String("Internal Server Error"),
				Message:    helpers.String(err.Error()),
			})
		}

		return operations.NewListStacksOK().WithPayload(&models.ResponseListTfStacks{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfStack{
				Stacks: stacks,
			},
		})
	})
}

var CreateTfStackController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.DeployStackHandlerFunc {
	return operations.DeployStackHandlerFunc(func(params operations.DeployStackParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfs := params.TerraformStack
		tfs.ID = uuid.New()

		workspace := models.ResourceTfWorkspace{
			Name: params.TerraformStack.Workspace,
		}
		if err := db.Read(DB_TABLE_TF_WORKSPACE, *workspace.Name, &workspace); err != nil {
			return operations.NewDeployStackBadRequest().WithPayload(&models.ServerError{
				StatusCode: HTTP_BAD_REQUEST,
				Status:     HTTP_BAD_REQUEST_STATUS,
				Message: helpers.String(fmt.Sprintf(
					"Could not find workspace with name'%s'",
					params.TerraformStack.Workspace)),
			})
		}

		tfs.Deployments = []*models.ResourceTfDeployment{{
			ID:     uuid.New(),
			Status: TF_DEPLOYMENT_STATUS_DEPLOYING,
		}}

		if err := db.Create(DB_TABLE_TF_STACK, tfs.ID, tfs); err != nil {
			return operations.NewDeployStackBadRequest()
		}

		tfs.Links = helpers.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + tfs.ID)
		tfs.Links.Doc = helpers.HalDocLink(ch).Doc
		tfs.Status = TF_STACK_STATUS_WAITING_FOR_DEPLOYMENT

		return operations.NewDeployStackAccepted().WithPayload(tfs)
	})
}

var GetTfStackController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.GetStackHandlerFunc {
	return operations.GetStackHandlerFunc(func(params operations.GetStackParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		var stack *models.ResourceTfStack
		err := db.Read(DB_TABLE_TF_STACK, id, stack)

		if err != nil {
			return operations.NewGetStackInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    helpers.String(err.Error()),
			})
		} else if stack == nil {
			return operations.NewGetStackNotFound().WithPayload(&models.ServerError{
				StatusCode: HTTP_NOT_FOUND,
				Status:     HTTP_NOT_FOUND_STATUS,
				Message:    helpers.String("Could not find stack with id '" + id + "'"),
			})
		} else {
			stack.Links = helpers.HalSelfLink(ch.FQEndpoint)
			stack.Links.Doc = helpers.HalDocLink(ch).Doc
			return operations.NewGetStackOK().WithPayload(stack)
		}
	})
}
