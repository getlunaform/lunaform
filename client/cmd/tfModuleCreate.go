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
	"fmt"
)

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
		module, resp, err := gocdClient.TfApi.CreateModule(ctx, map[string]interface{}{
			"terraformModule": map[string]interface{}{
				"name": "test",
				"type": "test",
			},
		})

		fmt.Print(resp)

		handleOutput(cmd, &module, useHal, err)
	},
}

func init() {
	tfModulesCmd.AddCommand(tfModuleCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tfModuleCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tfModuleCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
