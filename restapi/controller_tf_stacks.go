package restapi

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/backend/workers"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	operations "github.com/getlunaform/lunaform/restapi/operations/stacks"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"strings"
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

var ListTfStacksController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.ListStacksHandlerFunc {
	return operations.ListStacksHandlerFunc(func(params operations.ListStacksParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		stacks := make([]*models.ResourceTfStack, 0)
		err := db.List(DB_TABLE_TF_STACK, &stacks)
		if err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		}

		for _, stack := range stacks {
			stack.GenerateLinks(strings.TrimSuffix(ch.Endpoint, "s"))
			stack.Embedded = nil
		}

		return operations.NewListStacksOK().WithPayload(&models.ResponseListTfStacks{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfStack{
				Stacks: stacks,
			},
		})
	})
}

var GetTfStackController = func(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.GetStackHandlerFunc {
	return operations.GetStackHandlerFunc(func(params operations.GetStackParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		stack := &models.ResourceTfStack{}

		if err := db.Read(DB_TABLE_TF_STACK, id, stack); err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		} else if stack == nil {
			return NewServerErrorResponse(http.StatusNotFound, "Could not find stack with id '"+id+"'")
		} else {
			stack.Links = helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)

			stack.Embedded.Workspace.Modules = nil
			stack.Embedded.Workspace.GenerateLinks(ch.ServerURL + "/tf/workspace")
			for _, dep := range stack.Embedded.Deployments {
				dep.Status = nil
				dep.Workspace = nil
				dep.GenerateLinks(ch.Endpoint + "/deployment")
			}

			return operations.NewGetStackOK().WithPayload(stack)
		}
	})
}

var ListTfStackDeploymentsController = func(
	idp identity.Provider, ch *helpers.ContextHelper,
	db database.Database,
	workerPool *workers.TfAgentPool,
) operations.ListDeploymentsHandlerFunc {
	return operations.ListDeploymentsHandlerFunc(func(params operations.ListDeploymentsParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		stack := &models.ResourceTfStack{}
		deployments := &models.ResponseListTfDeployments{}

		if err := db.Read(DB_TABLE_TF_STACK, id, stack); err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		} else if stack == nil {
			return NewServerErrorResponse(http.StatusNotFound, "Could not find stack with id '"+id+"'")
		}
		deployments.Embedded.Deployments = stack.Embedded.Deployments
		deployments.Embedded.Stack = stack
		stack.Embedded.Deployments = nil

		deployments.Links = helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			ch.Endpoint,
		)

		return operations.NewListDeploymentsOK().WithPayload(deployments)
	})
}
