package identity

import (
	"errors"
	"github.com/zeebox/terraform-server/backend/database"
)

func NewDatabaseIdentityProvider(db database.Database) (idp Provider, err error) {
	return nil, errors.New("Database IdP Not Implemented")
}
