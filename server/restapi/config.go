package restapi

import (
	"encoding/json"
	"github.com/go-openapi/swag"
	"github.com/zeebox/terraform-server/server/restapi/operations"
)

var cliconfig = ConfigFileFlags{}

type Configuration struct {
	Identity CfgIdentity `json:"identity"`
	Backend  CfgBackend  `json:"backend"`
}

type CfgIdentity struct {
	Defaults []CfgIdentityDefault `json:"defaults"`
}

type CfgIdentityDefault struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type CfgBackend struct {
	DatabaseType string      `json:"database_type"`
	Database     interface{} `json:"database"`
	IdentityType string      `json:"identity_type"`
	Identity     interface{} `json:"identity"`
}

func (cfg *Configuration) loadFromFile(path string) {
	j, err := swag.YAMLDoc(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(j, cfg)
	if err != nil {
		panic(err)
	}
}

type ConfigFileFlags struct {
	ConfigFile string `short:"c" long:"config" description:"Path to configuration on disk"`
}

func parseCliConfiguration() *Configuration {
	cfg := newDefaultConfiguration()
	if cliconfig.ConfigFile != "" {
		cfg.loadFromFile(cliconfig.ConfigFile)
	}
	return cfg
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
		},
	}
}

func configureFlags(api *operations.TerraformServerAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Config",
			LongDescription:  "Configuration",
			Options:          &cliconfig,
		},
	}
}
