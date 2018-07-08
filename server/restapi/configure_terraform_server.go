package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/tylerb/graceful"

	"github.com/drewsonne/terraform-server/server/restapi/operations"

	"github.com/drewsonne/terraform-server/backend/database"
	"github.com/drewsonne/terraform-server/backend/identity"
	"log"
	"net/http/httputil"
	"net/http/httptest"
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

	// Controllers for /
	api.ResourcesListResourceGroupsHandler = ListResourceGroupsController(idp, oh)

	// Controllers for /{group}
	api.ResourcesListResourcesHandler = ListResourcesController(idp, oh)

	// Controllers for /tf/modules
	api.TfListModulesHandler = ListTfModulesController(idp, oh, db)
	api.TfCreateModuleHandler = CreateTfModuleController(idp, oh, db)

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
	return logResponse("", logRequest(handler))
}

func logRequest(handler http.Handler) http.Handler {
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	//	handler.ServeHTTP(w, r)
	//})
	return handler
}

func logResponse(prefix string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, false)
		if err != nil {
			log.Println(err)
		}
		log.Println(prefix, string(requestDump))

		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, r)

		dump, err := httputil.DumpResponse(rec.Result(), false)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(prefix, string(dump))

		// we copy the captured response headers to our new response
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		// grab the captured response body
		data := rec.Body.Bytes()
		w.WriteHeader(rec.Code)

		w.Write(data)
	}
}
