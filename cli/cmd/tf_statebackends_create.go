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
	"github.com/getlunaform/lunaform-client-go/state_backends"
	models "github.com/getlunaform/lunaform-models-go"
	"github.com/spf13/cobra"
)

var tfStatebackendCreateNameFlag string
var tfStatebackendCreateConfigFlag string

// tfStatebackendCreateCmd represents the tfStatebackendCreate command
var tfStatebackendCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		config := map[string]interface{}{}
		json.Unmarshal([]byte(tfStatebackendCreateConfigFlag), &config)

		params := state_backends.NewCreateStateBackendParams().WithTerraformStateBackend(
			&models.ResourceTfStateBackend{
				Name:          tfStatebackendCreateNameFlag,
				Configuration: &config,
			},
		)

		backend, err := gocdClient.StateBackends.CreateStateBackend(params, authHandler)
		if err == nil {
			handleOutput(cmd, backend.Payload, useHal, err)
		} else if err1, ok := err.(*state_backends.CreateStateBackendBadRequest); ok {
			handleOutput(cmd, err1.Payload, useHal, nil)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}

	},
}

func init() {
	tfStatebackendCreateCmd.Flags().
		StringVar(&tfStatebackendCreateNameFlag, "name", "", "Name of the terraform workspace")
	tfStatebackendCreateCmd.Flags().
		StringVar(&tfStatebackendCreateConfigFlag, "configuration", "",
			"A JSON string describing the configuration for the state backend")
	tfStatebackendCreateCmd.MarkFlagRequired("name")
	tfStatebackendCreateCmd.MarkFlagRequired("configuration")
}
