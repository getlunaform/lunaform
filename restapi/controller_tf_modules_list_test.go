package restapi

import (
	"reflect"
	"testing"

	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
)

func Test_buildListTfModules(t *testing.T) {
	type args struct {
		db database.Database
		ch helpers.ContextHelper
	}
	tests := []struct {
		name    string
		args    args
		wantM   *models.ResourceListTfModule
		wantErr middleware.Responder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, gotErr := buildListTfModules(tt.args.db, tt.args.ch)
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("buildListTfModules() gotM = %v, want %v", gotM, tt.wantM)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("buildListTfModules() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
