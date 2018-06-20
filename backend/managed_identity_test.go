package backend

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewsonne/terraform-server/backend/identity"
	"testing"
)

// TestManagedIdentityProviders for terraform-server to create, update, and destroy users and groups.
// These providers are managed by terraform-server meaning that terraform-server is responsible for being
// its own identity provider. Therefore, all managed terraform servers should conform to the following conditions
func TestManagedIdentityProvider(t *testing.T) {

	for key, idp := range map[string]identity.Provider{
		"Memory": identity.NewMemoryIdentityProvider(),
	} {

		t.Run(key, func(t *testing.T) {

			t.Run("I can validate my ability to update my user in an IdP", func(*testing.T) {
				assert.True(t, idp.IsEditable())
			})

			t.Run("I can validate the IdP is not federated", func(*testing.T) {
				assert.False(t, idp.IsFederated())
			})

			t.Run("I get an error when I try to use a federated endpoint", func(t2 *testing.T) {
				assert.EqualError(t, idp.ConsumeEndpoint(nil), "Can not consume endpoint for managed IdP")
			})

			t.Run("I can authenticate a user against an IdP", func(*testing.T) {
				admin, err := idp.ReadUser("admin")
				assert.Nil(t, err)
				assert.Equal(t, "admin", admin.Username)

				loggedIn := admin.Login("password")
				assert.True(t, loggedIn)
			})

			t.Run("I can create a user in my IdP", func(*testing.T) {
				user, err := idp.CreateUser("test-user", "test-password")
				assert.NotNil(t, user)
				assert.Nil(t, err)

				assert.Equal(t, user.Username, "test-user")

				user1, err := idp.ReadUser("test-user")
				assert.Nil(t, err)
				assert.False(t, user1.LoggedIn())
			})

			t.Run("I get an error trying to create a user in my IdP if they already exist", func(*testing.T) {
				user1, _ := idp.CreateUser("test-user", "test-password")
				assert.NotNil(t, user1)

				_, err := idp.CreateUser("test-user", "test-password")
				assert.EqualError(t, err, "User 'test-user' already exists")

			})

			t.Run("I can change a users password", func(*testing.T) {

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
