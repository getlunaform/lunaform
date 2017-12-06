package identity

import (
	"github.com/zeebox/terraform-server/backend"
	"errors"
)

func NewDatabaseIdentityProvider(db backend.Database) (idp backend.IdentityProvider, err error) {
	return nil, errors.New("IdP Not Implemented")
}
