package restapi

import (
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
)

func buildResourceGroupResponse(rscs []string, ch *helpers.ContextHelper) (rsclist *models.ResourceList) {
	rsclist = &models.ResourceList{
		Resources: make([]*models.Resource, len(rscs)),
	}
	for i, rsc := range rscs {
		rsclist.Resources[i] = &models.Resource{
			Name:  swag.String(rsc),
			Links: helpers.HalSelfLink(nil, ch.EndpointSingular+"/"+rsc),
		}
	}
	return
}
