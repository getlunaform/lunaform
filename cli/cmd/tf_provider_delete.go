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

var tfProviderDeleteNameFlag string

// tfProviderDeleteCmd represents the tfProviderDelete command
var tfProviderDeleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove provider",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := lunaformClient.Providers.DeleteProvider(
			providers.NewDeleteProviderParams().WithName(tfProviderDeleteNameFlag),
			authHandler,
		)

		if err1, ok := err.(*providers.DeleteProviderNotFound); ok {
			handleOutput(cmd, err1.Payload, useHal, nil)
		} else if err == nil {
			handleOutput(cmd, models.StringHalResponse("Successfully deleted"), useHal, err)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}

	},
}

func init() {
	flags := tfProviderDeleteCmd.Flags()
	flags.StringVar(&tfProviderDeleteNameFlag,
		"name", "", "Terraform provider name")
	tfProviderDeleteCmd.MarkFlagRequired("name")
}
