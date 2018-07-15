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
	"github.com/drewsonne/lunaform/client/modules"
)

var nameFlag string
var typeFlag string
var sourceFlag string

// tfModuleCreateCmd represents the tfModuleCreate command
var tfModuleCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Link a git or terraform module repository to the server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		params := modules.NewCreateModuleParams().WithTerraformModule(
			&models.ResourceTfModule{
				Name:   &nameFlag,
				Type:   &typeFlag,
				Source: &sourceFlag,
			},
		)
		module, err := gocdClient.Modules.CreateModule(params, authHandler)

		if err == nil {
			handleOutput(cmd, module.Payload, useHal, err)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}
	},
}

func init() {

	flags := tfModuleCreateCmd.Flags()
	flags.StringVar(&nameFlag, "name", "", "Name of the terraform module")
	flags.StringVar(&typeFlag, "type", "", "Type of the module. One of {git,registry,enterprise}")
	flags.StringVar(&sourceFlag, "source", "", "Source of the terraform module")

	tfModuleCreateCmd.MarkFlagRequired("name")
	tfModuleCreateCmd.MarkFlagRequired("type")
	tfModuleCreateCmd.MarkFlagRequired("source")
}
