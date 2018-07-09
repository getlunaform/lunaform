package cmd

import (
	"fmt"
	"encoding/json"
	"os"
	tfs "github.com/drewsonne/terraform-server/cli/library"
	"github.com/spf13/cobra"
	"strings"
)

func handlerError(err error) {
	fmt.Print(
		jsonResponse(
			map[string]string{
				"error": err.Error(),
			},
		),
	)
	os.Exit(1)
}

func handleOutput(action *cobra.Command, v tfs.HalLinked, hal bool, err error) {
	if err != nil {
		handlerError(err)
	} else {

		payload := map[string]interface{}{
			"action": strings.Join(
				buildActionName(action, []string{}),
				".",
			),
		}

		if !hal {
			payload["response"] = v.Clean()
		} else {
			payload["response"] = v
		}

		fmt.Print(jsonResponse(payload))
	}
}

func jsonResponse(r interface{}) string {
	out, err := json.MarshalIndent(r, "", "  ")

	if err != nil {
		panic(err)
	}

	return string(out) + "\n"
}

func buildActionName(c *cobra.Command, names []string) []string {
	if parent := c.Parent(); parent != nil {
		names = append([]string{c.Use}, names...)
		names = buildActionName(parent, names)
	}
	return names
}
