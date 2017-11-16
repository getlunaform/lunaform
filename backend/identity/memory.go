package identity

import (
	"github.com/drewsonne/terraform-server/backend"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func NewMemoryIdentityProvider() MemoryIdentityProvider {
	return MemoryIdentityProvider{
		users: make(map[string]backend.User),
	}
}

type MemoryIdentityProvider struct {
	users map[string]backend.User
}

func (mip MemoryIdentityProvider) IsEditable() (editable bool) {
	return true
}

func (mip MemoryIdentityProvider) IsFederated() (federated bool) {
	return false
}

func (mip MemoryIdentityProvider) ConsumeEndpoint(payload []byte) (err error) {
	return
}

func (mip MemoryIdentityProvider) CreateUser(username string, password string) (user backend.User, err error) {
	if _, exists := mip.users[username]; exists {
		return user, fmt.Errorf("User '%s' already exists", username)
	}

	user = backend.User{
		IsEditable: true,
		Username:   username,
		Idp:        mip,
	}
	user.Password, err = mip.hashPassword(password)

	return
}

func (mip MemoryIdentityProvider) ReadUser(username string) (user backend.User, err error) {
	user, exists := mip.users[username]
	if username == "admin" && !exists {
		var pwd string
		if pwd, err = mip.hashPassword("password"); err != nil {
			return
		}
		user = backend.User{
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

func (mip MemoryIdentityProvider) UpdateUser(user backend.User) (err error) {
	return
}

func (mip MemoryIdentityProvider) LoginUser(user backend.User, password string) (loggedin bool) {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func (mip MemoryIdentityProvider) ChangePassword(user backend.User, password string) (err error) {
	user.Password, err = mip.hashPassword(password)
	mip.users[user.Username] = user
	return
}

func (mip MemoryIdentityProvider) hashPassword(password string) (pwd string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
