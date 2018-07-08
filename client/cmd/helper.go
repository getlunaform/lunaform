package cmd

import (
	"fmt"
	"encoding/json"
	"os"
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

func handleOutput(action string, v interface{}, err error) {
	if err != nil {
		handlerError(err)
	} else {
		fmt.Print(jsonResponse(map[string]interface{}{
			"action":   action,
			"response": v,
		}))
	}

}

func jsonResponse(r interface{}) string {
	out, err := json.MarshalIndent(r, "", "  ")

	if err != nil {
		panic(err)
	}

	return string(out)
}
