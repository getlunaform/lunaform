package restapi

import (
	"encoding/json"

	"github.com/getlunaform/lunaform/server/restapi/operations"
	"github.com/go-openapi/swag"
)

var cliconfig = ConfigFileFlags{}

// Configuration describes the structure of options in the server config file
type Configuration struct {
	Identity CfgIdentity `json:"identity"`
	Backend  CfgBackend  `json:"backend"`
}

// CfgIdentity describes the structure of options for Identity Providers
type CfgIdentity struct {
	Defaults []CfgIdentityDefault `json:"defaults"`
}

// CfgIdentityDefault allows the setting of a username and password for a default user. This value will only be used
// when initialising a new managed Identity Provider, and will be ignored on subsequent boots.
// @TODO Restrict this to only be for the `admin` user
// @TODO Allow a force cli option when booting to reset the password
type CfgIdentityDefault struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

// CfgBackend describes how the server can load the backend database and the primary managed Identity Provider
type CfgBackend struct {
	DatabaseType string      `json:"database_type"`
	Database     interface{} `json:"database"`
	IdentityType string      `json:"identity_type"`
	Identity     interface{} `json:"identity"`
}

func (cfg *Configuration) loadFromFile(path string) (err error) {
	var j json.RawMessage
	if j, err = swag.YAMLDoc(path); err == nil {
		err = json.Unmarshal(j, cfg)
	}
	return
}

//ConfigFileFlags for loading settings for the server
type ConfigFileFlags struct {
	ConfigFile  string `short:"c" long:"config" description:"Path to configuration on disk"`
	Version     bool   `short:"V" long:"version" description:"Print lunarform version and quit"`
	AdminApiKey string `long:"api-key" description:"Override the admin user's api key.'"`
}

func parseCliConfiguration() (cfg *Configuration, err error) {
	cfg = newDefaultConfiguration()
	if cliconfig.ConfigFile != "" {
		err = cfg.loadFromFile(cliconfig.ConfigFile)
	}
	return
}

func newDefaultConfiguration() *Configuration {
	return &Configuration{
		Identity: CfgIdentity{
			Defaults: []CfgIdentityDefault{
				{User: "admin", Password: "password"},
			},
		},
		Backend: CfgBackend{
			IdentityType: "memory",
			DatabaseType: "memory",
		},
	}
}

func configureFlags(api *operations.LunaformAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Terraform Server",
			LongDescription:  "Server Configuration",
			Options:          &cliconfig,
		},
	}
}
