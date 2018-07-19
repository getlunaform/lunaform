package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/backend/database"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/stacks"
	"github.com/go-openapi/runtime/middleware"
	models "github.com/getlunaform/lunaform-models-go"
	"github.com/drewsonne/lunaform/server/helpers"
	"fmt"
	"strings"
	"github.com/drewsonne/lunaform/backend/workers"
	"github.com/go-openapi/swag"
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
				StatusCode: swag.Int64(500),
				Status:     swag.String("Internal Server Error"),
				Message:    swag.String(err.Error()),
			})
		}

		for _, stack := range stacks {
			stack.Embedded = nil
			stack.GenerateLinks(strings.TrimSuffix(ch.FQEndpoint, "s"))
		}

		return operations.NewListStacksOK().WithPayload(&models.ResponseListTfStacks{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfStack{
				Stacks: stacks,
			},
		})
	})
}

var CreateTfStackController = func(
	idp identity.Provider, ch helpers.ContextHelper,
	db database.Database,
	workerPool *workers.TfAgentPool,
) operations.DeployStackHandlerFunc {
	return operations.DeployStackHandlerFunc(func(params operations.DeployStackParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfs := params.TerraformStack
		tfs.ID = idGenerator.MustGenerate()

		workspace := models.ResourceTfWorkspace{
			Name: swag.String(params.TerraformStack.Workspace),
		}
		if err := db.Read(DB_TABLE_TF_WORKSPACE, tfs.Workspace, &workspace); err != nil {
			return operations.NewDeployStackBadRequest().WithPayload(&models.ServerError{
				StatusCode: HTTP_BAD_REQUEST,
				Status:     HTTP_BAD_REQUEST_STATUS,
				Message: swag.String(fmt.Sprintf(
					"Could not find workspace with name'%s'",
					params.TerraformStack.Workspace)),
			})
		}

		dep := NewTfDeployment(
			*workspace.Name,
		)

		tfs.Embedded = &models.ResourceTfStackEmbedded{
			Deployments: []*models.ResourceTfDeployment{dep},
			Workspace:   &workspace,
		}

		workerPool.DoPlan(&workers.TfActionPlan{
			Stack:      tfs,
			Deployment: dep,
		})

		if err := db.Create(DB_TABLE_TF_STACK, tfs.ID, tfs); err != nil {
			return operations.NewDeployStackBadRequest()
		}

		tfs.Links = helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			strings.TrimSuffix(ch.Endpoint, "s")+"/"+tfs.ID,
		)
		tfs.Status = TF_STACK_STATUS_WAITING_FOR_DEPLOYMENT

		return operations.NewDeployStackAccepted().WithPayload(tfs)
	})
}

var GetTfStackController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.GetStackHandlerFunc {
	return operations.GetStackHandlerFunc(func(params operations.GetStackParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		stack := &models.ResourceTfStack{}

		if err := db.Read(DB_TABLE_TF_STACK, id, stack); err != nil {
			return operations.NewGetStackInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		} else if stack == nil {
			return operations.NewGetStackNotFound().WithPayload(&models.ServerError{
				StatusCode: HTTP_NOT_FOUND,
				Status:     HTTP_NOT_FOUND_STATUS,
				Message:    swag.String("Could not find stack with id '" + id + "'"),
			})
		} else {
			stack.Links = helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.FQEndpoint,
			)

			stack.Embedded.Workspace.Modules = nil
			stack.Embedded.Workspace.GenerateLinks(ch.ServerURL + "/tf/workspace")
			for _, dep := range stack.Embedded.Deployments {
				dep.Status = nil
				dep.Workspace = nil
				dep.GenerateLinks(ch.FQEndpoint + "/deployment")
			}

			return operations.NewGetStackOK().WithPayload(stack)
		}
	})
}

var ListTfStackDeploymentsController = func(
	idp identity.Provider, ch helpers.ContextHelper,
	db database.Database,
	workerPool *workers.TfAgentPool,
) operations.ListDeploymentsHandlerFunc {
	return operations.ListDeploymentsHandlerFunc(func(params operations.ListDeploymentsParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		var stack *models.ResourceTfStack
		var deployments *models.ResponseListTfDeployments

		if err := db.Read(DB_TABLE_TF_STACK, id, stack); err != nil {
			return operations.NewListStacksInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		} else if stack == nil {
			return operations.NewGetStackNotFound().WithPayload(&models.ServerError{
				StatusCode: HTTP_NOT_FOUND,
				Status:     HTTP_NOT_FOUND_STATUS,
				Message:    swag.String("Could not find stack with id '" + id + "'"),
			})
		} else {
			deployments.Embedded.Deployments = stack.Embedded.Deployments
			deployments.Embedded.Stack = stack
			stack.Embedded.Deployments = nil

			deployments.Links = helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)

			return operations.NewListDeploymentsOK().WithPayload(deployments)
		}
		return
	})
}
