package restapi

import "testing"

func TestConfigureTerraformServer(t *testing.T) {
	// Our TLS configuration does nothing
	configureTLS(nil)

	//Server config does nothing either
	configureServer(nil, "", "")
}
