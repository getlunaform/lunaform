package identity

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewsonne/lunarform/backend/database"
	"testing"
)

func TestDatabaseIdP(t *testing.T) {
	idp, err := NewDatabaseIdentityProvider(database.Database{})
	assert.EqualError(t, err, "Database IdP Not Implemented")
	assert.Nil(t, idp)
}
