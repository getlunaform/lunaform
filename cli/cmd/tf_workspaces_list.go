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
	"github.com/drewsonne/lunaform/client/workspaces"
	"github.com/drewsonne/lunaform/server/models"
)

// tfWorkspacesListCmd represents the tfWorkspacesList command
var tfWorkspacesListCmd = &cobra.Command{
	Use:   "tfWorkspacesList",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		workspaces, err := gocdClient.Workspaces.ListWorkspaces(
			workspaces.NewListWorkspacesParams(),
			authHandler,
		)
		var response models.HalLinkable
		if err != nil {
			response = nil
		} else {
			response = workspaces.Payload
		}
		handleOutput(cmd, response, useHal, err)
	},
}
