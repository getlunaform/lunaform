package restapi

import (
	"github.com/drewsonne/terraform-server/backend/identity"
	"github.com/drewsonne/terraform-server/backend/database"
	operations "github.com/drewsonne/terraform-server/server/restapi/operations/stacks"
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/terraform-server/server/models"
	"encoding/json"
	"strings"
	"github.com/pborman/uuid"
)

var ListTfStacksController = func(idp identity.Provider, ch ContextHelper, db database.Database) operations.ListStacksHandlerFunc {
	return operations.ListStacksHandlerFunc(func(params operations.ListStacksParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		records, err := db.List("tf-stack")
		if err != nil {
			return operations.NewListStacksInternalServerError().WithPayload(&models.ServerError{
				StatusCode: Int64(500),
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		}

		stacks := make([]*models.ResourceTfStack, len(records))
		for i, record := range records {
			stack := models.ResourceTfStack{}
			json.Unmarshal([]byte(record.Value), &stack)
			stack.Links = halSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + stack.ID)
			stacks[i] = &stack
		}

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
