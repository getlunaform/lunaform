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
	"github.com/drewsonne/lunaform/client/state_backends"
	"github.com/drewsonne/lunaform/server/models"
)

var tfStateBackendsUpdateIdFlag string
var tfStateBackendsUpdateConfigFlag string
var tfStateBackendsUpdateNameFlag string

// tfStateBackendsUpdateCmd represents the tfStateBackendsUpdate command
var tfStateBackendsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := &models.ResourceTfStateBackend{}
		if tfStateBackendsUpdateConfigFlag != "" {
			payload.Configuration = tfStateBackendsUpdateConfigFlag
		}
		if tfStateBackendsUpdateNameFlag != "" {
			payload.Name = tfStateBackendsUpdateNameFlag
		}
		params := state_backends.NewUpdateStateBackendParams().
			WithID(tfStateBackendsUpdateIdFlag).
			WithTerraformStateBackend(payload)
		backend, err := gocdClient.StateBackends.UpdateStateBackend(params, authHandler)

		var response models.HalLinkable
		if err != nil {
			response = nil
		} else {
			response = backend.Payload
		}
		handleOutput(cmd, response, useHal, err)
	},
}

func init() {
	tfStateBackendsUpdateCmd.Flags().
		StringVar(&tfStateBackendsUpdateIdFlag, "id", "", "ID of the terraform state backend to update")
	tfStateBackendsUpdateCmd.Flags().
		StringVar(&tfStateBackendsUpdateNameFlag, "name", "", "Name of the terraform state backend to update")
	tfStateBackendsUpdateCmd.Flags().
		StringVar(&tfStateBackendsUpdateConfigFlag, "configuration", "",
		"A JSON string describing the configuration for the state backend")
	tfStateBackendsUpdateCmd.MarkFlagRequired("id")
}
