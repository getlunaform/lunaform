package identity

import (
	"errors"
	"github.com/zeebox/terraform-server/backend"
)

func NewDatabaseIdentityProvider(db backend.Database) (idp backend.IdentityProvider, err error) {
	return nil, errors.New("Database IdP Not Implemented")
}
