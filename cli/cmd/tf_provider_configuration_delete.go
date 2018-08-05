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
	"github.com/spf13/cobra"
	"github.com/getlunaform/lunaform/client/providers"
	"github.com/getlunaform/lunaform/models"
)

var tfProviderConfigurationDeleteProviderNameFlag string
var tfProviderConfigurationDeleteProvideConfigurationIdFlag string

// tfProviderConfigurationDeleteCmd represents the tfProviderConfigurationDelete command
var tfProviderConfigurationDeleteCmd = &cobra.Command{
	Use:   "delete-configuration",
	Short: "Delete provider configuration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		params := providers.NewDeleteProviderConfigurationParams().
			WithProviderName(tfProviderConfigurationDeleteProviderNameFlag).
			WithID(tfProviderConfigurationDeleteProvideConfigurationIdFlag)

		_, err := lunaformClient.Providers.DeleteProviderConfiguration(
			params, authHandler)

		if err1, ok := err.(*providers.DeleteProviderConfigurationNotFound); ok {
			handleOutput(cmd, err1.Payload, useHal, nil)
		} else if err == nil {
			handleOutput(cmd, models.StringHalResponse("Successfully deleted"), useHal, err)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}
	},
}

func init() {
	flags := tfProviderConfigurationDeleteCmd.Flags()
	flags.StringVar(&tfProviderConfigurationDeleteProviderNameFlag,
		"provider-name", "", "Provider name")
	flags.StringVar(&tfProviderConfigurationDeleteProvideConfigurationIdFlag,
		"id", "", "Provider configuration is")
	tfProviderConfigurationDeleteCmd.MarkFlagRequired("provider-name")
	tfProviderConfigurationDeleteCmd.MarkFlagRequired("id")
}
