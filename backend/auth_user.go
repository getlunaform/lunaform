package backend

import (
	"fmt"
	"time"
)

type User struct {
	IsEditable bool
	Username   string
	Password   string
	ApiKeys    []*ApiKey
	SSHKeys    []*SSHKey
	Idp        IdentityProvider
	loggedin   bool
	groups     []Group
}

func (u *User) Login(password string) bool {
	u.loggedin = u.Idp.LoginUser(*u, password)
	return u.LoggedIn()
}

func (u *User) Logout() {
	u.loggedin = false
}

func (u *User) LoggedIn() bool {
	return u.loggedin
}

func (u *User) ChangePassword(password string) (err error) {
	if !u.LoggedIn() {
		return fmt.Errorf("Could not change password on '%s' as user is not logged in", u.Username)
	}

	return u.Idp.ChangePassword(*u, password)
}

type SSHKey struct {
	Public          []byte
	Private         []byte
	PublicPath      string
	PrivatePath     string
	ServerGenerated bool
}

type ApiKey struct {
	Value                string
	DateCreated          time.Time
	DateExpired          time.Time
	ValidationPeriod     time.Duration
	AutomaticallyExpired bool
}
