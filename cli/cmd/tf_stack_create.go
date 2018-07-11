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
	"github.com/spf13/cobra"
	"github.com/drewsonne/terraform-server/client/client/tf"
	"github.com/drewsonne/terraform-server/server/models"
)

var flagModule string
var flagModuleId string
var flagName string

// tfStackCreateCmd represents the tfStackCreate command
var tfStackCreateCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var module *models.ResourceTfModule
		if flagModuleId == "" && flagModule != "" {
			modules, err := gocdClient.Tf.ListModules(
				tf.NewListModulesParams(),
			)
			if err != nil {
				handleOutput(cmd, nil, useHal, err)
			}
			for _, module = range modules.Payload.Embedded.Resources {
				if *module.Name == flagModule {
					break
				}
			}
		} else if flagModuleId != "" {
			moduleResponse, err := gocdClient.Tf.GetModule(tf.NewGetModuleParams().WithID(
				flagModuleId,
			))
			if err != nil {
				handleOutput(cmd, nil, useHal, err)
			}
			module = moduleResponse.Payload
		}

		tf.NewDeployStackParams().WithTerraformStack(
			&models.ResourceTfStack{},
		)
		stack, err := gocdClient.Tf.DeployStack(tf.NewDeployStackParams().WithTerraformStack(
			&models.ResourceTfStack{
				ModuleID: String(module.VcsID),
				Name:     String(flagName),
			},
		))
		if err == nil {
			handleOutput(cmd, stack.Payload, useHal, err)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}

	},
}

func init() {
	tfStackCmd.AddCommand(tfStackCreateCmd)

	flags := tfStackCreateCmd.Flags()
	flags.StringVar(&flagModule, "module", "", "Name of the terraform module to deploy")
	flags.StringVar(&flagModuleId, "module-id", "", "ID of the terraform module to deploy")
	flags.StringVar(&flagName, "name", "", "Name of the deployed terraform module")

	tfStackCreateCmd.MarkFlagRequired("name")
}
