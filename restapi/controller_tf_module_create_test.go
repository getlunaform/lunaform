package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/models/hal"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func Test_buildTfModuleControllerCreateResponse(t *testing.T) {
	for _, tt := range []struct {
		name string
		ch   *helpers.ContextHelper
	}{
		{
			name: "base",
			ch: &helpers.ContextHelper{
				ServerURL:        "http://example.com/api",
				BasePath:         "/api",
				Endpoint:         "/",
				EndpointSingular: "/",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			db := getTestingDB([]map[string]string{})

			mod := models.ResourceTfModule{}
			_, err := buildTfModuleControllerCreateResponse(&mod, db, tt.ch)
			assert.NoError(t, err)

			createdModule := models.ResourceTfModule{}
			err = db.Read(DB_TABLE_TF_MODULE, mod.ID, &createdModule)
			assert.NoError(t, err)

			assert.Equal(t, []*hal.HalCurie{{
				Href:      strfmt.URI(tt.ch.ServerURL + "/{rel}"),
				Name:      "lf",
				Templated: true,
			}, {
				Href:      strfmt.URI(tt.ch.ServerURL + "/docs#operation/{rel}"),
				Name:      "doc",
				Templated: true,
			}}, mod.Links.Curies)

			assert.Equal(t, map[string]*hal.HalHref{
				"lf:self": {
					Href: "/",
				},
				"doc:": {
					Href: "/",
				},
			}, mod.Links.HalRscLinks)

			createdModule.Links = nil
			mod.Links = nil
			assert.Equal(t, mod, createdModule)

		})
	}
}
