// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"github.com/getlunaform/lunaform/client/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/spf13/cobra"
)

var (
	tfProviderConfigurationCreateNameFlag              string
	tfProviderConfigurationCreateProviderNameFlag      string
	tfProviderConfigurationCreateJSONConfigurationFlag string
)

// tfProviderConfigurationCreateCmd represents the tfProviderConfigurationCreate command
var tfProviderConfigurationCreateCmd = &cobra.Command{
	Use:   "create-configuration",
	Short: "Create provider configuration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		configuration := map[string]interface{}{}
		err := json.Unmarshal(
			[]byte(tfProviderConfigurationCreateJSONConfigurationFlag),
			&configuration,
		)
		if err != nil {
			handleOutput(cmd, nil, useHal, err)
		}

		params := providers.NewCreateProviderConfigurationParams().
			WithProviderName(tfProviderConfigurationCreateProviderNameFlag).
			WithProviderConfiguration(
				&models.ResourceTfProviderConfiguration{
					Name:          String(tfProviderConfigurationCreateNameFlag),
					Configuration: configuration,
				})

		provider, err := lunaformClient.Providers.CreateProviderConfiguration(
			params,
			authHandler,
		)

		if err == nil {
			handleOutput(cmd, provider.Payload, useHal, err)
		} else {
			if err1, hasPayload := err.(*providers.CreateProviderConfigurationInternalServerError); hasPayload {
				handleOutput(cmd, err1.Payload, useHal, err)
			} else {
				handleOutput(cmd, nil, useHal, err)
			}
		}

	},
}

func init() {
	flags := tfProviderConfigurationCreateCmd.Flags()

	flags.StringVar(&tfProviderConfigurationCreateNameFlag,
		"name", "", "Configuration Name")
	flags.StringVar(&tfProviderConfigurationCreateProviderNameFlag,
		"provider-name", "", "Terraform provider name")
	flags.StringVar(&tfProviderConfigurationCreateJSONConfigurationFlag,
		"configuration", "", "JSON Encoded Provider configuration")

	tfProviderConfigurationCreateCmd.MarkFlagRequired("name")
	tfProviderConfigurationCreateCmd.MarkFlagRequired("provider-name")
	tfProviderConfigurationCreateCmd.MarkFlagRequired("configuration")
}
