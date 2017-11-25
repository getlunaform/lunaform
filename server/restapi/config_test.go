package restapi

import (
	"testing"
	"gopkg.in/yaml.v2"
	"github.com/stretchr/testify/assert"
)

func TestConfigFile(t *testing.T) {

	defaultConfigPayload := []byte(`
---
identity:
  defaults:
    - user: admin
      password: mock_password
backend:
  database:
    type: json
    path: config/db.yaml
  identity:
    - type: json
       path: config/auth-db.yaml`)

	c := cfg{}

	err := yaml.Unmarshal(defaultConfigPayload, &c)
	assert.Nil(t, err)

	assert.NotNil(t, c.Identity.Defaults)
	assert.Len(t, c.Identity.Defaults, 1)

	u := c.Identity.Defaults[0]
	assert.NotNil(t, u)
	assert.Equal(t, "admin", u.User)
	assert.Equal(t, "mock_password", u.Password) // @TODO THis should be a bcrypt

}
