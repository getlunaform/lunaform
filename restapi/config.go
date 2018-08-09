package restapi

import (
	"encoding/json"
	jww "github.com/spf13/jwalterweatherman"

	"fmt"
	"github.com/getlunaform/lunaform/restapi/operations"
	"github.com/go-openapi/swag"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
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
	jww.SetLogThreshold(jww.LevelDebug)

	cfg = newDefaultConfiguration()
	if cliconfig.ConfigFile == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".client" (without extension).
		viper.AddConfigPath(home + "/.config/")
		viper.SetConfigName("lunaform-server")

	} else {
		viper.SetConfigFile(cliconfig.ConfigFile)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.Unmarshal(&cfg)

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
