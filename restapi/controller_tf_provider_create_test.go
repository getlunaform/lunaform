package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/stretchr/testify/assert"
	"github.com/getlunaform/lunaform/models/hal"
)

func Test_buildCreateTfProviderResponse(t *testing.T) {
	type args struct {
		provider *models.ResourceTfProvider
		db       database.Database
		ch       helpers.ContextHelper
	}
	for _, tt := range []struct {
		name        string
		args        args
		wantErrCode int
		wantErr     error
	}{
		{
			name: "base",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			db := getTestingDB([]map[string]string{})
			prov := models.ResourceTfProvider{Name: "mock-provider"}

			gotErrCode, err := buildCreateTfProviderResponse(&prov, db, tt.args.ch)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantErrCode, gotErrCode)

			readProv := models.ResourceTfProvider{}
			err = db.Read(DB_TABLE_TF_PROVIDER, prov.Name, &readProv)

			readProv.Links = &hal.HalRscLinks{
				Curies: []*hal.HalCurie{
					{Href: "/{rel}", Name: "lf", Templated: true},
					{Href: "/docs#operation/{rel}", Name: "doc", Templated: true},
				},
				HalRscLinks: map[string]*hal.HalHref{
					"doc:":    {Href: "/"},
					"lf:self": {Href: "/"},
				},
			}

			assert.NoError(t, err)
			assert.Equal(t, prov, readProv)
		})
	}
}
