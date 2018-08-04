package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/swag"
	"github.com/getlunaform/lunaform/models/hal"
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
					{Name: swag.String("one"), Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/one"},
						}}},
					{Name: swag.String("two"), Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/two"},
						}}},
					{Name: swag.String("three"), Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/three"},
						}}},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t,
				buildResourceGroupResponse(
					tt.resources, tt.ch,
				),
				tt.wantRsclist,
			)
		})
	}
}
