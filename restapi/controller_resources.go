package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/go-openapi/swag"
)

const (
	DB_TABLE_TF_WORKSPACE    = "lf-workspace"
	DB_TABLE_TF_MODULE       = "lf-module"
	DB_TABLE_TF_STACK        = "lf-stack"
	DB_TABLE_TF_STATEBACKEND = "lf-statebackend"
	DB_TABLE_AUTH_USER       = "lf-auth-user"
)



func buildResourceGroupResponse(rscs []string, ch helpers.ContextHelper) (rsclist *models.ResourceList) {
	rsclist = &models.ResourceList{
		Resources: make([]*models.Resource, len(rscs)),
	}
	for i, rsc := range rscs {
		rsclist.Resources[i] = &models.Resource{
			Name:  swag.String(rsc),
			Links: helpers.HalSelfLink(nil, ch.Endpoint+"/"+rsc),
		}
	}
	return
}
