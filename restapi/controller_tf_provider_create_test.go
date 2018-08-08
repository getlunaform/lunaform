package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/stretchr/testify/assert"
	"github.com/getlunaform/lunaform/models/hal"
	"github.com/go-openapi/swag"
)

func Test_buildCreateTfProviderResponse(t *testing.T) {
	type args struct {
		provider *models.ResourceTfProvider
		db       database.Database
		ch       *helpers.ContextHelper
	}
	for _, tt := range []struct {
		name        string
		args        args
		wantErrCode int
		wantErr     error
	}{
		{
			name: "base",
			args: args{
				ch: &helpers.ContextHelper{
					EndpointSingular: "/",
					Endpoint:         "/",
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			db := getTestingDB([]map[string]string{})
			wantProvider := models.ResourceTfProvider{Name: swag.String("mock-provider")}

			gotErrCode, err := buildCreateTfProviderResponse(&wantProvider, db, tt.args.ch)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantErrCode, gotErrCode)

			gotProvider := models.ResourceTfProvider{}
			err = db.Read(DB_TABLE_TF_PROVIDER, *wantProvider.Name, &gotProvider)

			gotProvider.Links = &hal.HalRscLinks{
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
			assert.Equal(t, wantProvider, gotProvider)
		})
	}
}
