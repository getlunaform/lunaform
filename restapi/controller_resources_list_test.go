package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/stretchr/testify/assert"
	"github.com/go-openapi/swag"
)

func Test_buildResourceGroupRootResponse(t *testing.T) {

	for _, tt := range []struct {
		name        string
		ch          helpers.ContextHelper
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
					Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/tf/modules"}},
					}},
				{Name: swag.String("stacks"),
					Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/tf/stacks"}},
					}},
				{Name: swag.String("state-backends"),
					Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/tf/state-backends"}},
					}},
				{Name: swag.String("workspaces"),
					Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/tf/workspaces"}},
					}},
			}},
			ch: helpers.ContextHelper{
				Endpoint: "/tf",
			},
		},
		{
			name: "identity",
			wantRsclist: &models.ResourceList{Resources: []*models.Resource{
				{Name: swag.String("groups"),
					Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/identity/groups"}},
					}},
				{Name: swag.String("providers"),
					Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/identity/providers"}},
					}},
				{Name: swag.String("users"),
					Links: &models.HalRscLinks{HalRscLinksAdditionalProperties: map[string]interface{}{
						"lf:self": &models.HalHref{Href: "/identity/users"}},
					}},
			}},
			ch: helpers.ContextHelper{
				Endpoint: "/identity",
			},
		},
		{
			name: "vcs",
			wantRsclist: &models.ResourceList{Resources: []*models.Resource{
				{Name: swag.String("git"),
					Links: &models.HalRscLinks{
						HalRscLinksAdditionalProperties: map[string]interface{}{
							"lf:self": &models.HalHref{Href: "/vcs/git"}},
					}},
			}},
			ch: helpers.ContextHelper{
				Endpoint: "/vcs",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotRsclist := buildResourceGroupRootResponse(tt.name, tt.ch);
			assert.Equal(t, tt.wantRsclist, gotRsclist)
		})
	}
}
