package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/drewsonne/terraform-server/backend"
)

func TestMemoryIdP(t *testing.T) {
	t.Run("Interface", testMemoryIdPInterface)
	t.Run("DefaultLogin", testMemoryIdPLogin)
	t.Run("CreateUser", testMemoryIdPCreateUser)
	t.Run("ChangePassword", testMemoryIdPChangePassword)
}

func testMemoryIdPInterface(t *testing.T) {
	var idp backend.IdentityProvider
	idp = NewMemoryIdentityProvider()
	assert.NotNil(t, idp)
}

func testMemoryIdPLogin(t *testing.T) {
	var idp backend.IdentityProvider
	idp = NewMemoryIdentityProvider()

	admin, err := idp.ReadUser("admin")
	assert.Nil(t, err)
	assert.Equal(t, "admin", admin.Username)

	loggedIn := admin.Login("password")
	assert.True(t, loggedIn)
}

func testMemoryIdPCreateUser(t *testing.T) {
	var idp backend.IdentityProvider
	idp = NewMemoryIdentityProvider()

	user, err := idp.CreateUser("test-user", "test-password")
	assert.NotNil(t, user)
	assert.Nil(t, err)

	assert.Equal(t, user.Username, "test-user")

	user1, err := idp.ReadUser("test-user")
	assert.Nil(t, err)
	assert.False(t, user1.LoggedIn())
}

func testMemoryIdPChangePassword(t *testing.T) {
	var idp backend.IdentityProvider
	idp = NewMemoryIdentityProvider()

	admin, _ := idp.ReadUser("admin")
	assert.True(t, admin.IsEditable)

	err := admin.ChangePassword("new_password")
	assert.NotNil(t, err)

	admin.Login("password")
	err = admin.ChangePassword("new_password")
	assert.Nil(t, err)

	admin1, _ := idp.ReadUser("admin")
	assert.True(t, admin.LoggedIn())
	assert.False(t, admin1.LoggedIn())
	admin1.Login("new_password")
	assert.True(t, admin1.LoggedIn())

}
