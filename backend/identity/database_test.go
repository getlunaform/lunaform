package identity

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabaseIdP(t *testing.T) {
	idp, err := NewDatabaseIdentityProvider(database.Database{})
	assert.NoError(t, err)
	assert.NotNil(t, idp)
}
