package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/backend/workers"
	"github.com/getlunaform/lunaform/helpers"
	operations "github.com/getlunaform/lunaform/restapi/operations/stacks"
	"github.com/go-openapi/runtime/middleware"
	"strings"
	"net/http"
	"fmt"
)

var CreateTfStackController = func(
	idp identity.Provider, ch helpers.ContextHelper,
	db database.Database,
	workerPool *workers.TfAgentPool,
) operations.DeployStackHandlerFunc {
	return operations.DeployStackHandlerFunc(func(params operations.DeployStackParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfs := params.TerraformStack
		tfs.ID = idGenerator.MustGenerate()

		if tfs.Embedded == nil {
			tfs.Embedded = &models.ResourceTfStackEmbedded{
				Deployments:            []*models.ResourceTfDeployment{},
				Module:                 &models.ResourceTfModule{},
				ProviderConfigurations: make([]*models.ResourceTfProviderConfiguration, 0),
				Workspace:              &models.ResourceTfWorkspace{},
			}

		}
		workspace := tfs.Embedded.Workspace
		module := tfs.Embedded.Module

		if err := db.Read(DB_TABLE_TF_WORKSPACE, tfs.WorkspaceName, &workspace); err != nil {
			return NewServerError(
				http.StatusBadRequest,
				fmt.Sprintf("Could not find workspace with name'%s'", *workspace.Name),
			)
		}
		tfs.WorkspaceName = "" // Clear the input values

		if err := db.Read(DB_TABLE_TF_MODULE, tfs.ModuleID, &module); err != nil {
			return NewServerError(
				http.StatusBadRequest,
				fmt.Sprintf("Could not find module with id '%s'", module.ID),
			)
		}
		tfs.ModuleID = ""

		// Do this so that we don't have nested JSON and end up with a stack overflow.
		tfs.Embedded.Module = &models.ResourceTfModule{
			Name: tfs.Embedded.Module.Name,
			ID:   tfs.Embedded.Module.ID,
		}

		dep := NewTfDeployment(tfs.WorkspaceName)

		tfs.Embedded.Deployments = append(tfs.Embedded.Deployments, dep)

		for _, providerConfigurationId := range tfs.ProviderConfigurationsIds {
			provConf := models.ResourceTfProviderConfiguration{}
			if err := db.Read(DB_TABLE_TF_PROVIDER_CONFIGURATION, providerConfigurationId, &provConf); err != nil {
				return NewServerError(
					http.StatusBadRequest,
					fmt.Sprintf("Could not find provider configuration with id '%s'", providerConfigurationId),
				)
			}
			tfs.Embedded.ProviderConfigurations = append(tfs.Embedded.ProviderConfigurations, &provConf)
		}

		go workerPool.DoPlan(&workers.TfActionPlan{
			Stack:      tfs,
			Deployment: dep,
			Module:     module,
			DoInit:     true,
		})

		if err := db.Create(DB_TABLE_TF_STACK, tfs.ID, tfs); err != nil {
			return operations.NewDeployStackBadRequest()
		}

		if module.Embedded == nil {
			module.Embedded = &models.ResourceListTfStack{}
		}
		if module.Embedded.Stacks == nil {
			module.Embedded.Stacks = make([]*models.ResourceTfStack, 0)
		}
		module.Embedded.Stacks = append(module.Embedded.Stacks, tfs)
		if err := db.Update(DB_TABLE_TF_MODULE, module.ID, module); err != nil {
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
