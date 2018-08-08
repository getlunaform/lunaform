package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/getlunaform/lunaform/models/hal"
)

func Test_buildGetTfModuleResponse(t *testing.T) {
	for _, tt := range []struct {
		name     string
		moduleId string
		module   *models.ResourceTfModule
		want     *CommonServerErrorResponder
	}{
		{
			name:     "basic",
			moduleId: "mock-id",
			module: &models.ResourceTfModule{
				Embedded: &models.ResourceListTfStack{
					Stacks: make([]*models.ResourceTfStack, 0),
				},
				Links: &hal.HalRscLinks{
					HalRscLinks: map[string]*hal.HalHref{
						"lf:self": {Href: "/"},
						"doc:":    {Href: "/"},
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			db := getTestingDB([]map[string]string{})
			ch := &helpers.ContextHelper{}

			if tt.module != nil {
				err := db.Create(DB_TABLE_TF_MODULE, tt.moduleId, tt.module)
				assert.NoError(t, err)
			}

			dbMod := models.ResourceTfModule{}
			err := buildGetTfModuleResponse(tt.moduleId, &dbMod, db, ch)
			assert.Equal(t, tt.want, err)
			assert.Equal(t, tt.module, &dbMod)
		})
	}
}
