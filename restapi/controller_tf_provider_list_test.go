package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/models"
	"github.com/stretchr/testify/assert"
)

func Test_buildListTfProvidersResponse(t *testing.T) {

	for _, tt := range []struct {
		name        string
		providers   []*models.ResourceTfProvider
		wantErrCode int
		wantErr     error
	}{
		{
			name: "base",
			providers: []*models.ResourceTfProvider{{
				Name: "one",
			}, {
				Name: "two",
			}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			db := getTestingDB([]map[string]string{})
			for _, prov := range tt.providers {
				err := db.Create(DB_TABLE_TF_PROVIDER, prov.Name, prov)
				assert.NoError(t, err)
			}

			providers, gotErrCode, err := buildListTfProvidersResponse(db)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantErrCode, gotErrCode)

			assert.Equal(t, tt.providers, providers)

		})
	}
}
