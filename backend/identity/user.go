package identity

import (
	"fmt"
	"time"
)

// User who can log in, be assigned to groups, and is defined within an IdP.
type User struct {
	IsEditable bool
	Username   string
	Password   string
	APIKeys    []*APIKey
	SSHKeys    []*SSHKey
	Idp        Provider
	loggedin   bool
	groups     []Group
}

// Login a user with a plaintext password
func (u *User) Login(password string) bool {
	u.loggedin = u.Idp.LoginUser(u, password)
	return u.LoggedIn()
}

// Logout a user
func (u *User) Logout() {
	u.loggedin = false
}

// LoggedIn returns true if the user is logged in
func (u *User) LoggedIn() bool {
	return u.loggedin
}

// ChangePassword for a user if the IdP allows password changes
func (u *User) ChangePassword(password string) (err error) {
	if !u.LoggedIn() {
		return fmt.Errorf("Could not change password on '%s' as user is not logged in", u.Username)
	}

	return u.Idp.ChangePassword(u, password)
}

// SSHKey contains both the public and private sections, and the location of the sshkey on disk
type SSHKey struct {
	Public          []byte
	Private         []byte
	PublicPath      string
	PrivatePath     string
	ServerGenerated bool
}

// APIKey for a user, to allow them to log in without a username and password
type APIKey struct {
	Value                string
	DateCreated          time.Time
	DateExpired          time.Time
	ValidationPeriod     time.Duration
	AutomaticallyExpired bool
}
