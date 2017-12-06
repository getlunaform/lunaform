package identity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/zeebox/terraform-server/backend"
)

func TestDatabaseIdP(t *testing.T) {
	idp, err := NewDatabaseIdentityProvider(backend.Database{})
	assert.EqualError(t, err, "Database IdP Not Implemented")
	assert.Nil(t, idp)
}
