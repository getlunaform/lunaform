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
	"github.com/getlunaform/lunaform/client/providers"
	"github.com/spf13/cobra"
)

var (
	tfProviderConfigurationGetFlagId           string
	tfProviderConfigurationGetFlagProviderName string
)

// tfProviderConfigurationGetCmd represents the tfProviderConfigurationGet command
var tfProviderConfigurationGetCmd = &cobra.Command{
	Use:   "get-configuration",
	Short: "Describe provider configuration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		params := providers.NewGetProviderConfigurationParams().
			WithProviderName(tfProviderConfigurationGetFlagProviderName).
			WithID(tfProviderConfigurationGetFlagId)
		prov, err := lunaformClient.Providers.GetProviderConfiguration(
			params, authHandler)
		if err == nil {
			handleOutput(cmd, prov.Payload, useHal, err)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}
	},
}

func init() {
	flags := tfProviderConfigurationGetCmd.Flags()
	flags.StringVar(&tfProviderConfigurationGetFlagId,
		"id", "", "Provider Configuration ID")
	flags.StringVar(&tfProviderConfigurationGetFlagProviderName,
		"provider-name", "", "Provider Name")

	tfProviderConfigurationGetCmd.MarkFlagRequired("id")
	tfProviderConfigurationGetCmd.MarkFlagRequired("provider-name")
}
