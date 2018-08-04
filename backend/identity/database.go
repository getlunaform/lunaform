package identity

import (
	"github.com/getlunaform/lunaform/backend/database"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

const (
	DB_TABLE_AUTH_USER   = database.DBTableRecordType("lf-auth-user")
	DB_TABLE_AUTH_APIKEY = database.DBTableRecordType("lf-auth-apikey")
)

// NewDatabaseIdentityProvider is not yet implemented and will return an error
func NewDatabaseIdentityProvider(db database.Database) (idp Provider, err error) {
	return dbIdentityProvider{
		db: db,
	}, nil
}

type dbIdentityProvider struct {
	db database.Database
}

func (dbidp dbIdentityProvider) IsEditable() (editable bool) {
	return true
}

func (dbidp dbIdentityProvider) IsFederated() (federated bool) {
	return false
}

func (dbidp dbIdentityProvider) ConsumeEndpoint(payload []byte) (err error) {
	return errors.New("Can not consume endpoint for db based IdP")
}

func (dbidp dbIdentityProvider) CreateUser(newUser *User) (user *User, err error) {
	if err = dbidp.db.Create(DB_TABLE_AUTH_USER, newUser.Username, newUser); err != nil {
		return
	}
	for _, key := range newUser.APIKeys {
		if err = dbidp.db.Create(DB_TABLE_AUTH_APIKEY, key.Value, newUser); err != nil {
			return
		}
		fmt.Printf("Generated api-key for '%s' user '%s'\n", newUser.Username, key.Value)
	}

	return dbidp.ReadUser(newUser.Username)
}

func (dbidp dbIdentityProvider) ReadUser(username string) (user *User, err error) {
	user = &User{}
	if err = dbidp.db.Read(DB_TABLE_AUTH_USER, username, user); err != nil {
		if _, userRecordNotFound := err.(database.RecordDoesNotExistError); userRecordNotFound {
			err = UserNotFound(err)
			user = nil
		} else {
			return
		}
	}
	return user, err
}
func (dbidp dbIdentityProvider) UpdateUser(username string, user *User) (updatedUser *User, err error) {
	if err = dbidp.db.Update(DB_TABLE_AUTH_USER, username, user); err != nil {
		if _, userRecordNotFound := err.(database.RecordDoesNotExistError); userRecordNotFound {
			err = UserNotFound(err)
		}
	}
	return user, nil
}

func (dbidp dbIdentityProvider) LoginUser(user *User, password string) (loggedin bool) {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func (dbidp dbIdentityProvider) ChangePassword(user *User, password string) (err error) {
	user.Password = password
	return dbidp.db.Update(DB_TABLE_AUTH_USER, user.Username, user)
}
