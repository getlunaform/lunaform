package restapi

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestingDB(content []map[string]string) database.Database {
	dbDriver, err := database.NewMemoryDBDriverWithCollection(content)
	if err != nil {
		panic(err)
	}

	return database.NewDatabaseWithDriver(dbDriver)
}

func Test_newContextHelperWithContext(t *testing.T) {

	ctxHelper, err := helpers.NewContextHelper(api.Context())
	assert.NoError(t, err)
	for _, tt := range []struct {
		name string
		ctx  *middleware.Context
		want *helpers.ContextHelper
	}{
		{
			name: "base",
			ctx:  api.Context(),
			want: ctxHelper.WithBasePath("/api"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := helpers.NewContextHelperWithContext(tt.ctx)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
