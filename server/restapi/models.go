package restapi

import (
	"github.com/drewsonne/lunaform/server/models"
	"github.com/pborman/uuid"
)

func NewTfDeployment(workspace string) *models.ResourceTfDeployment {
	return &models.ResourceTfDeployment{
		ID:        uuid.New(),
		Status:    TF_DEPLOYMENT_STATUS_DEPLOYING,
		Workspace: workspace,
	}
}
