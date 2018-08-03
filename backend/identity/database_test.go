package identity

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabaseIdP(t *testing.T) {
	idp, err := NewDatabaseIdentityProvider(database.Database{})
	assert.EqualError(t, err, "Database IdP Not Implemented")
	assert.Nil(t, idp)
}
