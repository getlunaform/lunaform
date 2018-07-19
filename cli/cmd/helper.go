package cmd

import (
	"fmt"
	"encoding/json"
	"os"
	"github.com/spf13/cobra"
	"strings"
	models "github.com/getlunaform/lunaform-models-go"
	"github.com/go-openapi/runtime"
)

func handlerError(err error) {

	var errResponse interface{}
	if errApi, isApiError := err.(*runtime.APIError); isApiError {
		response := map[string]interface{}{
			"code":           errApi.Code,
			"operation-name": errApi.OperationName,
		}
		
		errResponse = response
	} else {
		errResponse = err.Error()
	}
	fmt.Print(
		jsonResponse(
			map[string]interface{}{
				"error": errResponse,
			},
		),
	)
	os.Exit(1)
}

func handleOutput(action *cobra.Command, v models.HalLinkable, hal bool, err error) {
	if err != nil {
		handlerError(err)
	} else {

		payload := map[string]interface{}{
			"action": strings.Join(buildActionName(
				action,
				[]string{},
			), "."),
		}

		if hal {
			payload["response"] = v
		} else {
			payload["response"] = v.Clean()
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

func String(s string) *string {
	return &s
}
