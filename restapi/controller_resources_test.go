package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/swag"
)

func Test_buildResourceGroupResponse(t *testing.T) {

	for _, tt := range []struct {
		name        string
		resources   []string
		wantRsclist *models.ResourceList
		ch          helpers.ContextHelper
	}{
		{
			name:      "base",
			resources: []string{"one", "two", "three"},
			wantRsclist: &models.ResourceList{
				Resources: []*models.Resource{
					{Name: swag.String("one"), Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/one"},
						}}},
					{Name: swag.String("two"), Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/two"},
						}}},
					{Name: swag.String("three"), Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/three"},
						}}},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			response := buildResourceGroupResponse(tt.resources, tt.ch)
			assert.Equal(t, response, tt.wantRsclist)
		})
	}
}
