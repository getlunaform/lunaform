package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/tylerb/graceful"

	"github.com/zeebox/terraform-server/server/restapi/operations"

	bmw "github.com/zeebox/go-http-middleware"
	"github.com/zeebox/goose4"

	"github.com/zeebox/terraform-server/backend/database"
	"github.com/zeebox/terraform-server/backend/identity"
)

// goose4
var (
	buildNumber  string
	buildMachine string
	builtBy      string
	builtWhen    string
	compiler     string
	sha          string
)

// goose4
const (
	runbookURI = "http://github.com/zeebox/terraform-server"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../server --name TerraformServer --spec ../swagger.yml --principal models.Principal

func configureAPI(api *operations.TerraformServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	var idp identity.Provider
	var dbDriver database.Driver
	var err error

	cfg, err := parseCliConfiguration()
	if err != nil {
		panic(err)
	}

	switch cfg.Backend.DatabaseType {
	case "memory":
		dbDriver, err = database.NewMemoryDBDriver()
	default:
		err = fmt.Errorf("unexpected Database type: '%s'", cfg.Backend.DatabaseType)
	}

	if err != nil {
		panic(err)
	}

	db := database.NewDatabaseWithDriver(dbDriver)

	switch cfg.Backend.IdentityType {
	case "memory":
		idp = identity.NewMemoryIdentityProvider()
	case "database":
		idp, err = identity.NewDatabaseIdentityProvider(db)
	default:
		panic(fmt.Sprintf("Unexpected Identity Provider type: '%s'", cfg.Backend.IdentityType))
	}

	if err != nil {
		panic(err)
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	oh := NewContextHelper(api.Context())

	// Controllers for /api/
	api.ResourcesListResourceGroupsHandler = ListResourceGroupsController(idp, oh)

	// Controllers for /api/{group}
	api.ResourcesListResourcesHandler = ListResourcesController(idp, oh)

	api.ServerShutdown = func() {
		dbDriver.Close()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {

	if builtWhen == "" {
		builtWhen = "-1"
	}
	t, err := strconv.Atoi(builtWhen)
	if err != nil {
		panic(err)
	}

	mw := bmw.NewMiddleware(handler)

	se4, err := goose4.NewGoose4(goose4.Config{
		ArtifactID:      "terraform-server",
		BuildNumber:     buildNumber,
		BuildMachine:    buildMachine,
		BuiltBy:         builtBy,
		BuiltWhen:       time.Unix(int64(t), 0),
		CompilerVersion: compiler,
		GitSha:          sha,
		RunbookURI:      runbookURI,
		Version:         buildNumber,
	})
	if err != nil {
		panic(err)
	}

	gmw := NewMiddleware(mw)
	gmw.SE4 = se4
	return gmw
}
