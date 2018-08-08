package restapi

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/go-openapi/runtime/middleware"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/getlunaform/lunaform/helpers"
)

func getTestingDB(content []map[string]string) database.Database {
	dbDriver, err := database.NewMemoryDBDriverWithCollection(content)
	if err != nil {
		panic(err)
	}

	return database.NewDatabaseWithDriver(dbDriver)
}

func Test_newContextHelperWithContext(t *testing.T) {

	for _, tt := range []struct {
		name string
		ctx  *middleware.Context
		want *helpers.ContextHelper
	}{
		{
			name: "base",
			ctx:  api.Context(),
			want: helpers.NewContextHelper(
				api.Context(),
			).WithBasePath("/api"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := helpers.NewContextHelperWithContext(tt.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}
