package identity

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// NewMemoryIdentityProvider creates a memory provider stored in memory.
// This is very volatile and should only be used for development or testing.
func NewMemoryIdentityProvider() Provider {
	return memoryIdentityProvider{
		users: make(map[string]User),
	}
}

// Memory IdentityProvider will store user details in RAM. Once this
// struct is released, all data is lost. This is really only used for
// development and will probably be deprecated in time.
type memoryIdentityProvider struct {
	users map[string]User
}

func (mip memoryIdentityProvider) IsEditable() (editable bool) {
	return true
}

func (mip memoryIdentityProvider) IsFederated() (federated bool) {
	return false
}

func (mip memoryIdentityProvider) ConsumeEndpoint(payload []byte) (err error) {
	return errors.New("Can not consume endpoint for managed IdP")
}

func (mip memoryIdentityProvider) CreateUser(username string, password string) (user User, err error) {
	if _, exists := mip.users[username]; exists {
		return user, fmt.Errorf("User '%s' already exists", username)
	}

	user = User{
		IsEditable: true,
		Username:   username,
		Idp:        mip,
	}
	user.Password, err = mip.hashPassword(password)

	mip.users[username] = user

	return
}

func (mip memoryIdentityProvider) ReadUser(username string) (user User, err error) {
	user, exists := mip.users[username]
	if username == "admin" && !exists {
		var pwd string
		if pwd, err = mip.hashPassword("password"); err != nil {
			return
		}
		user = User{
			IsEditable: mip.IsEditable(),
			Username:   "admin",
			Password:   pwd,
			Idp:        mip,
		}
		mip.users[user.Username] = user
	}

	user.Logout()

	return
}

func (mip memoryIdentityProvider) LoginUser(user User, password string) (loggedin bool) {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func (mip memoryIdentityProvider) ChangePassword(user User, password string) (err error) {
	user.Password, err = mip.hashPassword(password)
	mip.users[user.Username] = user
	return
}

func (mip memoryIdentityProvider) hashPassword(password string) (pwd string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
