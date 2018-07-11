// Copyright © 2018 Drew J. Sonne <drew.sonne@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"context"
	apiclient "github.com/drewsonne/terraform-server/client/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/runtime"
	jww "github.com/spf13/jwalterweatherman"
	"strings"
)

const (
	TERRAFORM_SERVER_TYPE_V1 = "application/vnd.terraform.server.v1+json"
)

var cfgFile string
var useHal bool
var gocdClient *apiclient.TerraformServerClient

var config Configuration
var version string

var logLevelMapping = map[string]jww.Threshold{
	"TRACE":    jww.LevelTrace,
	"DEBUG":    jww.LevelDebug,
	"INFO":     jww.LevelInfo,
	"WARN":     jww.LevelWarn,
	"ERROR":    jww.LevelError,
	"CRITICAL": jww.LevelCritical,
	"FATAL":    jww.LevelFatal,
}

type Configuration struct {
	Host    string
	Port    string
	Schemes []string
	Log struct {
		Level string
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tfs-client",
	Short: "A commandline application to interact with terraform-server",
	Long: `A commandline client to perform operations on 'terraform-server'.
These include module, and stack deployment, as well as user and permission management.
For example:

    $ tfs-client auth users list
    $ tfs-client tf modules list
    $ tfs-client tf modules create \
		--name my-module \
		--type git \
		--source git@github.com:zeebox/my-module.git
`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogging)
	cobra.OnInitialize(initGocdClient)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/tfs-client.yaml)")
	rootCmd.PersistentFlags().BoolVar(&useHal, "hal", false, "draw HAL elements in response")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".client" (without extension).
		viper.AddConfigPath(home + "/.config/")
		viper.SetConfigName("tfs-client")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.Unmarshal(&config)
}

func initLogging() {
	logLevel := strings.ToUpper(config.Log.Level)
	jww.SetLogThreshold(
		logLevelMapping[logLevel],
	)
}

func initGocdClient() {
	cfg := apiclient.DefaultTransportConfig().
		WithHost(config.Host + ":" + config.Port).
		WithSchemes(config.Schemes)
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	transport.Context = context.Background()

	transport.Producers[TERRAFORM_SERVER_TYPE_V1] = runtime.JSONProducer()
	transport.Consumers[TERRAFORM_SERVER_TYPE_V1] = runtime.JSONConsumer()

	gocdClient = apiclient.New(transport, strfmt.Default)
}
