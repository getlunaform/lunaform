package restapi

import (
	"testing"
	"github.com/go-openapi/loads"
	"net/http"
	"github.com/getlunaform/lunaform/restapi/operations"
)

func Test_configureTerraformServer(t *testing.T) {
	// Our TLS configuration does nothing
	configureTLS(nil)

	//Server config does nothing either
	configureServer(nil, "", "")
}

func getAPI() (*operations.LunaformAPI, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewLunaformAPI(swaggerSpec)
	return api, nil
}

func GetAPIHandler() (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}
	h := configureAPI(api)
	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return h, nil
}
