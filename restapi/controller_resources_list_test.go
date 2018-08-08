package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/swag"
	"github.com/getlunaform/lunaform/models/hal"
)

func Test_buildResourceGroupRootResponse(t *testing.T) {

	//curies := []*hal.HalCurie{{
	//	Href:      strfmt.URI("/{rel}"),
	//	Name:      "lf",
	//	Templated: true,
	//}, {
	//	Href:      strfmt.URI("/docs#operation/{rel}"),
	//	Name:      "lf",
	//	Templated: true,
	//}}

	for _, tt := range []struct {
		name        string
		ch          *helpers.ContextHelper
		wantRsclist *models.ResourceList
	}{
		{
			name:        "404",
			wantRsclist: &models.ResourceList{Resources: []*models.Resource{}},
		},
		{
			name: "tf",
			wantRsclist: &models.ResourceList{Resources: []*models.Resource{
				{Name: swag.String("modules"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/tf/modules"}},
					}},
				{Name: swag.String("stacks"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/tf/stacks"}},
					}},
				{Name: swag.String("state-backends"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/tf/state-backends"}},
					}},
				{Name: swag.String("workspaces"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/tf/workspaces"}},
					}},
				{Name: swag.String("providers"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/tf/providers"}},
					}},
			}},
			ch: &helpers.ContextHelper{
				Endpoint:         "/tf",
				EndpointSingular: "/tf",
			},
		},
		{
			name: "identity",
			wantRsclist: &models.ResourceList{Resources: []*models.Resource{
				{Name: swag.String("groups"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/identity/groups"}},
					}},
				{Name: swag.String("providers"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/identity/providers"}},
					}},
				{Name: swag.String("users"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/identity/users"}},
					}},
			}},
			ch: &helpers.ContextHelper{
				Endpoint:         "/identity",
				EndpointSingular: "/identity",
			},
		},
		{
			name: "vcs",
			wantRsclist: &models.ResourceList{Resources: []*models.Resource{
				{Name: swag.String("git"),
					Links: &hal.HalRscLinks{
						HalRscLinks: map[string]*hal.HalHref{
							"lf:self": {Href: "/vcs/git"}},
					}},
			}},
			ch: &helpers.ContextHelper{
				Endpoint:         "/vcs",
				EndpointSingular: "/vcs",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotRsclist := buildResourceGroupRootResponse(tt.name, tt.ch);
			assert.Equal(t, tt.wantRsclist, gotRsclist)
		})
	}
}
