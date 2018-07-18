package restapi

import (
	"github.com/drewsonne/lunaform/server/models"

	"github.com/teris-io/shortid"
	"github.com/go-openapi/swag"
)

func NewTfDeployment(workspace string) *models.ResourceTfDeployment {
	return &models.ResourceTfDeployment{
		ID:        shortid.MustGenerate(),
		Status:    swag.String(TF_DEPLOYMENT_STATUS_DEPLOYING),
		Workspace: swag.String(workspace),
	}
}
