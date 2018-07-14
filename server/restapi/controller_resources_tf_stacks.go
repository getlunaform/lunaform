package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/backend/database"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/stacks"
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/lunaform/server/models"
	"strings"
	"github.com/pborman/uuid"
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

var ListTfStacksController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.ListStacksHandlerFunc {
	return operations.ListStacksHandlerFunc(func(params operations.ListStacksParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		stacks := make([]*models.ResourceTfStack, 0)
		err := db.List("tf-stack",&stacks)
		if err != nil {
			return operations.NewListStacksInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		}

		//for i, record := range records {
		//	stack := models.ResourceTfStack{}
		//	json.Unmarshal([]byte(record.Value), &stack)
		//	stack.Links = halSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + stack.ID)
		//	stacks[i] = &stack
		//}

		return operations.NewListStacksOK().WithPayload(&models.ResponseListTfStacks{
			Links: halRootRscLinks(ch),
			Embedded: &models.ResourceListTfStack{
				Resources: stacks,
			},
		})
	})
}

var CreateTfStackController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.DeployStackHandlerFunc {
	return operations.DeployStackHandlerFunc(func(params operations.DeployStackParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfs := params.TerraformStack
		tfs.ID = uuid.New()

		tfs.Deployments = []*models.ResourceTfDeployment{
			{ID: uuid.New(), Status: TF_DEPLOYMENT_STATUS_DEPLOYING},
		}

		err := db.Create("tf-stack", tfs.ID, tfs)

		if err != nil {
			return operations.NewDeployStackBadRequest()
		}

		response := &models.ResourceTfStack{
			Links: halSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + tfs.ID),
			ID:    tfs.ID,
		}
		response.Links.Doc = halDocLink(ch).Doc

		if tfs == nil {
			return operations.NewDeployStackBadRequest()
		} else {
			response.Name = tfs.Name
			response.Status = TF_STACK_STATUS_WAITING_FOR_DEPLOYMENT
			response.ModuleID = tfs.ModuleID
			response.Deployments = tfs.Deployments
		}
		return operations.NewDeployStackAccepted().WithPayload(response)
	})
}

var GetTfStackController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.GetStackHandlerFunc {
	return operations.GetStackHandlerFunc(func(params operations.GetStackParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		id := params.ID

		var stack *models.ResourceTfStack
		err := db.Read("tf-stack", id, stack)

		if err != nil {
			return operations.NewGetStackInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		} else if stack == nil {
			return operations.NewGetStackNotFound().WithPayload(&models.ServerError{
				StatusCode: Int64(404),
				Status:     String("Not Found"),
				Message:    String("Could not find stack with id '" + id + "'"),
			})
		} else {
			stack.Links = halSelfLink(ch.FQEndpoint)
			stack.Links.Doc = halDocLink(ch).Doc
			return operations.NewGetStackOK().WithPayload(stack)
		}
	})
}
