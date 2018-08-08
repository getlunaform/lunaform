package restapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
	"net/http"
)

func Test_buildDeleteTfModuleResponse(t *testing.T) {
	for _, tt := range []struct {
		name     string
		moduleId string
		module   *models.ResourceTfModule
		wantErr  *CommonServerErrorResponder
	}{
		{
			name:     "non-existence",
			moduleId: "",
			wantErr:  nil,
		},
		{
			name:     "missing-delete",
			moduleId: "mock-id",
			module:   nil,
			wantErr:  nil,
		},
		{
			name:     "basic-delete",
			moduleId: "mock-id",
			module: &models.ResourceTfModule{
				ID:   "mock-id",
				Name: swag.String("mock name"),
				Embedded: &models.ResourceListTfStack{
					Stacks: make([]*models.ResourceTfStack, 0),
				},
			},
			wantErr: nil,
		},
		{
			name:     "basic-delete",
			moduleId: "mock-id",
			module: &models.ResourceTfModule{
				ID:   "mock-id",
				Name: swag.String("mock name"),
				Embedded: &models.ResourceListTfStack{
					Stacks: []*models.ResourceTfStack{
						{Name: swag.String("mock-1"), ID: "1"},
						{Name: swag.String("mock-2"), ID: "2"},
						{Name: swag.String("mock-3"), ID: "3"},
					},
				},
			},
			wantErr: NewServerErrorResponse(
				http.StatusUnprocessableEntity,
				"Could not delete module as it is relied up by stacks ['1','2','3']",
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			db := getTestingDB([]map[string]string{})

			if tt.moduleId != "" {
				err := db.Create(DB_TABLE_TF_MODULE, tt.moduleId, tt.module)
				assert.NoError(t, err)
			}

			got := buildDeleteTfModuleResponse(db, tt.moduleId)
			assert.Equal(t, tt.wantErr, got)

			if tt.wantErr == nil {
				readMod := models.ResourceTfModule{}
				err := db.Read(DB_TABLE_TF_MODULE, tt.moduleId, &readMod)
				_, isRecordDoesNotExistError := err.(database.RecordDoesNotExistError)
				assert.True(t, isRecordDoesNotExistError)
			}

		})
	}
}
