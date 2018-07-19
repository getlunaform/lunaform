package restapi

import (
	"github.com/getlunaform/lunaform/models"

	"github.com/go-openapi/swag"
)

func NewTfDeployment(workspace string) *models.ResourceTfDeployment {
	return &models.ResourceTfDeployment{
		ID:        idGenerator.MustGenerate(),
		Status:    swag.String(TF_DEPLOYMENT_STATUS_DEPLOYING),
		Workspace: swag.String(workspace),
	}
}
