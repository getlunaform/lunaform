package restapi

import (
	"github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
	"github.com/drewsonne/terraform-server/server/restapi/operations"
	"gopkg.in/yaml.v2"
	"testing"
)

var defaultConfigPayload = `
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
      path: config/auth-db.yaml`

func TestConfigFile(t *testing.T) {

	c := Configuration{}

	err := yaml.Unmarshal([]byte(defaultConfigPayload), &c)
	assert.Nil(t, err)

	assert.NotNil(t, c.Identity.Defaults)
	assert.Len(t, c.Identity.Defaults, 1)

	u := c.Identity.Defaults[0]
	assert.NotNil(t, u)
	assert.Equal(t, "admin", u.User)
	assert.Equal(t, "mock_password", u.Password) // @TODO THis should be a bcrypt
}

func TestCliOptions(t *testing.T) {
	api := operations.TerraformServerAPI{}

	assert.Empty(t, api.CommandLineOptionsGroups)

	configureFlags(&api)

	assert.Len(t, api.CommandLineOptionsGroups, 1)
	opt := api.CommandLineOptionsGroups[0]

	assert.Equal(t, "Config", opt.ShortDescription)
	assert.Equal(t, "Configuration", opt.LongDescription)

	assert.NotNil(t, opt.Options)
	assert.IsType(t, &ConfigFileFlags{}, opt.Options)

}

func TestLoadCliConfiguration(t *testing.T) {
	defer filet.CleanUp(t)
	file := filet.TmpFile(t, "", defaultConfigPayload)

	cliconfig = ConfigFileFlags{
		ConfigFile: file.Name(),
	}

	cfg, err := parseCliConfiguration()

	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	assert.IsType(t, &Configuration{}, cfg)

}
