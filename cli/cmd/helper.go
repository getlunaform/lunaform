package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime"
	"github.com/spf13/cobra"
	"os"
	"strings"
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
	printError(errResponse)
}

func handleServerError(err *models.ServerError) {
	printError(err)
}

func handleApiError(err *runtime.APIError) {

	response := map[string]interface{}{
		"operation-name": err.OperationName,
		"code":           err.Code,
		"response":       err.Response,
	}

	//switch e := err.Response.(type) {
	//case *http.Response:
	//	response["response"] = e
	//default:
	//	response["response"] = e
	//}

	printError(response)
}

func printError(errResponse interface{}) {
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
	if serverError, isServerError := v.(*models.ServerError); isServerError {
		handleServerError(serverError)
	} else if apiError, isApiError := err.(*runtime.APIError); isApiError {
		handleApiError(apiError)
	} else if err != nil {
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
