package restapi

import (
	"testing"

	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
)

func Test_buildGetTfProviderResponse(t *testing.T) {
	for _, tt := range []struct {
		name         string
		wantProvider *models.ResourceTfProvider
		wantErrCode  int
		wantErr      error
	}{
		{
			name:         "base",
			wantProvider: &models.ResourceTfProvider{Name: swag.String("mock")},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			db := getTestingDB([]map[string]string{})
			err := db.Create(DB_TABLE_TF_PROVIDER, *tt.wantProvider.Name, tt.wantProvider)
			assert.NoError(t, err)

			gotProvider, gotErrCode, err := buildGetTfProviderResponse(db, *tt.wantProvider.Name)
			assert.Equal(t, tt.wantProvider, gotProvider)
			assert.Equal(t, tt.wantErrCode, gotErrCode)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
