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
	"github.com/drewsonne/lunaform/server/models"
	"fmt"
	"github.com/drewsonne/lunaform/client/modules"
	"github.com/drewsonne/lunaform/client/stacks"
)

var tfStackCreateCmdModuleFlag string
var tfStackCreateCmdModuleIdFlag string
var tfStackCreateCmdNameFlag string
var tfStackCreateCmdWorkspaceFlag string

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
		var err error
		moduleSrcIsId := true
		if tfStackCreateCmdModuleIdFlag == "" && tfStackCreateCmdModuleFlag != "" {
			moduleSrcIsId = false
			modules, err := gocdClient.Modules.ListModules(
				modules.NewListModulesParams(),
				authHandler,
			)
			if err == nil {
				for _, module = range modules.Payload.Embedded.Modules {
					if *module.Name == tfStackCreateCmdModuleFlag {
						break
					}
				}
			}
		} else if tfStackCreateCmdModuleIdFlag != "" {
			moduleResponse, err := gocdClient.Modules.GetModule(
				modules.NewGetModuleParams().WithID(tfStackCreateCmdModuleIdFlag),
				authHandler,
			)
			if err == nil {
				module = moduleResponse.Payload
			}
		} else {
			err = fmt.Errorf("`--module` or `--module-id` must be provided")
		}

		if module == nil {
			if moduleSrcIsId {
				err = fmt.Errorf("could not find a module with id `" + tfStackCreateCmdModuleIdFlag + "`")
			} else {
				err = fmt.Errorf("could not find a module with name `" + tfStackCreateCmdModuleFlag + "`")
			}
		}

		if err != nil {
			handleOutput(cmd, nil, useHal, err)
		}

		params := stacks.NewDeployStackParams().WithTerraformStack(
			&models.ResourceTfStack{
				ModuleID:  String(module.ID),
				Name:      String(tfStackCreateCmdNameFlag),
				Workspace: tfStackCreateCmdWorkspaceFlag,
			},
		)

		if stack, err := gocdClient.Stacks.DeployStack(params, authHandler); err == nil {
			handleOutput(cmd, stack.Payload, useHal, err)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}

	},
}

func init() {
	flags := tfStackCreateCmd.Flags()
	flags.StringVar(&tfStackCreateCmdModuleFlag, "module", "", "Name of the terraform module to deploy")
	flags.StringVar(&tfStackCreateCmdModuleIdFlag, "module-id", "", "ID of the terraform module to deploy")
	flags.StringVar(&tfStackCreateCmdNameFlag, "name", "", "Name of the deployed terraform module")
	flags.StringVar(&tfStackCreateCmdWorkspaceFlag, "workspace", "", "Terraform workspace to deploy into.")

	tfStackCreateCmd.MarkFlagRequired("name")
	tfStackCreateCmd.MarkFlagRequired("workspace")
}
