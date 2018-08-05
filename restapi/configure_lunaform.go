package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/backend/workers"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/restapi/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"time"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../server --name TerraformServer --spec ../swagger.yml --principal models.Principal

var (
	version     string
	idGenerator *shortid.Shortid
)

const (
	DB_TABLE_AUTH_APIKEY = "lf-auth-apikey"
)

func configureAPI(api *operations.LunaformAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	var dbDriver database.Driver
	var err error

	var workerPool *workers.TfAgentPool
	var db database.Database
	var idp identity.Provider

	if idGenerator, err = shortid.New(1, shortid.DEFAULT_ABC, uint64(time.Now().UnixNano())); err != nil {
		panic(err)
	}

	cfg, err := parseCliConfiguration()
	if err != nil {
		panic(err)
	}

	switch cfg.Backend.DatabaseType {
	case "memory":
		dbDriver, err = database.NewMemoryDBDriver()
	case "file":
		var filePath string
		var isString bool
		if filePath, isString = cfg.Backend.Database.(string); !isString {
			panic(fmt.Errorf("DB config is not string. Is '%s'", cfg.Backend.Database))
		}
		dbDriver, err = database.NewJSONDBDriver(filePath)
	default:
		err = fmt.Errorf("unexpected Database type: '%s'", cfg.Backend.DatabaseType)
	}

	if err != nil {
		panic(err)
	}

	db = database.NewDatabaseWithDriver(dbDriver)
	workerPool = workers.NewAgentPool(5).
		WithDB(db).
		Start()

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

	oh := helpers.NewContextHelper(api.Context())

	api.APIKeyAuth = func(s string) (p *models.ResourceAuthUser, err error) {
		user := models.ResourceAuthUser{}
		if err = db.Read(DB_TABLE_AUTH_APIKEY, s, &user); err != nil {
			if _, isErrNotFound := err.(database.RecordDoesNotExistError); isErrNotFound {
				return nil, errors.Unauthenticated("http")
			}
			return
		}

		return &user, nil
	}

	configureRootUser(idp)
	configureDefaultWorkspace(&db)

	// Controllers for /
	api.ResourcesListResourceGroupsHandler = ListResourceGroupsController(oh)
	api.ResourcesListResourcesHandler = ListResourcesController(oh)

	// Controllers for /tf/modules
	api.ModulesListModulesHandler = ListTfModulesController(idp, oh, db)
	api.ModulesCreateModuleHandler = CreateTfModuleController(idp, oh, db)
	api.ModulesGetModuleHandler = GetTfModuleController(idp, oh, db)
	api.ModulesDeleteModuleHandler = DeleteTfModuleController(idp, oh, db)

	// Controllers for /tf/stacks
	api.StacksDeployStackHandler = CreateTfStackController(idp, oh, db, workerPool)
	api.StacksListStacksHandler = ListTfStacksController(idp, oh, db)
	api.StacksGetStackHandler = GetTfStackController(idp, oh, db)
	api.StacksUndeployStackHandler = DeleteTfStackController(idp, oh, db, workerPool)
	api.StacksListDeploymentsHandler = ListTfStackDeploymentsController(idp, oh, db, workerPool)

	// Controllers for /tf/workspaces
	api.WorkspacesDescribeWorkspaceHandler = GetTfWorkspaceController(idp, oh, db)
	api.WorkspacesListWorkspacesHandler = ListTfWorkspacesController(idp, oh, db)
	api.WorkspacesCreateWorkspaceHandler = CreateTfWorkspaceController(idp, oh, db)

	// Controllers for /tf/state-backends
	api.StateBackendsListStateBackendsHandler = ListTfStateBackendsController(idp, oh, db)
	api.StateBackendsCreateStateBackendHandler = CreateTfStateBackendsController(idp, oh, db)
	api.StateBackendsUpdateStateBackendHandler = UpdateTfStateBackendsController(idp, oh, db)

	// Controllers for /tf/providers
	api.ProvidersListProvidersHandler = ListTfProvidersController(idp, oh, db)
	api.ProvidersGetProviderHandler = GetTfProviderController(idp, oh, db)
	api.ProvidersCreateProviderHandler = CreateTfProviderController(idp, oh, db)
	api.ProvidersDeleteProviderHandler = DeleteTfProviderController(idp, oh, db)

	// Controllers for /tf/provider configurations
	api.ProvidersCreateProviderConfigurationHandler = CreateTfProviderConfigurationController(idp, oh, db)
	api.ProvidersListProviderConfigurationsHandler = ListTfProviderConfigurationController(idp, oh, db)
	api.ProvidersGetProviderConfigurationHandler = GetTfProviderConfigurationController(idp, oh, db)

	api.ServerShutdown = shutdownHandler(dbDriver, workerPool)

	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{{
		ShortDescription: "Version",
		Options: map[string]string{
			"one": "two",
		},
	}}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func shutdownHandler(dbDriver database.Driver, workerPool *workers.TfAgentPool) func() {
	return func() {
		fmt.Print("Shutdown handler")
		dbDriver.Close()
		workerPool.Shutdown()
	}
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {

	handler = logRequest(handler)
	if Debug {
		return logResponse("lunaform", handler)
	} else {
		return handler
	}
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

func configureDefaultWorkspace(db *database.Database) (err error) {

	defaultWorkspace := &models.ResourceTfWorkspace{}
	if err = db.Read(DB_TABLE_TF_WORKSPACE, "default", &defaultWorkspace); err != nil {
		if _, noDefaultWorkspace := err.(database.RecordDoesNotExistError); !noDefaultWorkspace {
			return
		} else if err = db.Create(DB_TABLE_TF_WORKSPACE, "default", &models.ResourceTfWorkspace{
			Modules: []*models.ResourceTfModule{},
			Name:    swag.String("default"),
		}); err != nil {
			return
		}
	}

	return
}

func configureRootUser(idp identity.Provider) (err error) {
	var (
		adminUser  *identity.User
		foundAdmin = true
	)
	if adminUser, err = idp.ReadUser("admin"); err != nil {
		if _, userNotFound := err.(identity.UserNotFound); !userNotFound {
			return
		} else {
			foundAdmin = false
		}
	}

	if !foundAdmin || cliconfig.AdminApiKey != "" {

		if cliconfig.AdminApiKey == "" {
			cliconfig.AdminApiKey = idGenerator.MustGenerate()
		}

		if !foundAdmin {
			adminUser = &identity.User{
				IsEditable: false,
				Username:   "admin",
			}
		}

		adminUser.APIKeys = []*identity.APIKey{{Value: cliconfig.AdminApiKey}}

		identity.AdminGroup().AddUser(adminUser)

		if foundAdmin {
			adminUser, err = idp.UpdateUser(adminUser.Username, adminUser)
		} else {
			adminUser, err = idp.CreateUser(adminUser)
		}
		if err != nil {
			return
		}
	}

	return
}
