package identity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestManagedIdentityProviders for terraform-server to create, update, and destroy users and groups.
// These providers are managed by terraform-server meaning that terraform-server is responsible for being
// its own identity provider. Therefore, all managed terraform servers should conform to the following conditions
func TestManagedIdentityProvider(t *testing.T) {

	for key, idp := range map[string]Provider{
		"Memory": NewMemoryIdentityProvider(),
	} {

		t.Run(key, func(t *testing.T) {

			t.Run("DefaultLogin", func(t *testing.T) {
				admin, err := idp.ReadUser("admin")
				assert.Nil(t, err)
				assert.Equal(t, "admin", admin.Username)

				loggedIn := admin.Login("password")
				assert.True(t, loggedIn)
			})

			t.Run("CreateUser", func(t *testing.T) {
				user, err := idp.CreateUser("test-user", "test-password")
				assert.NotNil(t, user)
				assert.Nil(t, err)

				assert.Equal(t, user.Username, "test-user")

				user1, err := idp.ReadUser("test-user")
				assert.Nil(t, err)
				assert.False(t, user1.LoggedIn())
			})

			t.Run("ChangePassword", func(t *testing.T) {

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

			})
		})

	}
}
