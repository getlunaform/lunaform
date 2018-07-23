// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/getlunaform/lunaform/restapi/operations/modules"
	"github.com/getlunaform/lunaform/restapi/operations/resources"
	"github.com/getlunaform/lunaform/restapi/operations/stacks"
	"github.com/getlunaform/lunaform/restapi/operations/state_backends"
	"github.com/getlunaform/lunaform/restapi/operations/workspaces"

	models "github.com/getlunaform/lunaform/models"
)

// NewLunaformAPI creates a new Lunaform instance
func NewLunaformAPI(spec *loads.Document) *LunaformAPI {
	return &LunaformAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ModulesCreateModuleHandler: modules.CreateModuleHandlerFunc(func(params modules.CreateModuleParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation ModulesCreateModule has not yet been implemented")
		}),
		StateBackendsCreateStateBackendHandler: state_backends.CreateStateBackendHandlerFunc(func(params state_backends.CreateStateBackendParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation StateBackendsCreateStateBackend has not yet been implemented")
		}),
		WorkspacesCreateWorkspaceHandler: workspaces.CreateWorkspaceHandlerFunc(func(params workspaces.CreateWorkspaceParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation WorkspacesCreateWorkspace has not yet been implemented")
		}),
		ModulesDeleteModuleHandler: modules.DeleteModuleHandlerFunc(func(params modules.DeleteModuleParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation ModulesDeleteModule has not yet been implemented")
		}),
		StacksDeployStackHandler: stacks.DeployStackHandlerFunc(func(params stacks.DeployStackParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation StacksDeployStack has not yet been implemented")
		}),
		WorkspacesDescribeWorkspaceHandler: workspaces.DescribeWorkspaceHandlerFunc(func(params workspaces.DescribeWorkspaceParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation WorkspacesDescribeWorkspace has not yet been implemented")
		}),
		ModulesGetModuleHandler: modules.GetModuleHandlerFunc(func(params modules.GetModuleParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation ModulesGetModule has not yet been implemented")
		}),
		StacksGetStackHandler: stacks.GetStackHandlerFunc(func(params stacks.GetStackParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation StacksGetStack has not yet been implemented")
		}),
		StacksListDeploymentsHandler: stacks.ListDeploymentsHandlerFunc(func(params stacks.ListDeploymentsParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation StacksListDeployments has not yet been implemented")
		}),
		ModulesListModulesHandler: modules.ListModulesHandlerFunc(func(params modules.ListModulesParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation ModulesListModules has not yet been implemented")
		}),
		ResourcesListResourceGroupsHandler: resources.ListResourceGroupsHandlerFunc(func(params resources.ListResourceGroupsParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourcesListResourceGroups has not yet been implemented")
		}),
		ResourcesListResourcesHandler: resources.ListResourcesHandlerFunc(func(params resources.ListResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation ResourcesListResources has not yet been implemented")
		}),
		StacksListStacksHandler: stacks.ListStacksHandlerFunc(func(params stacks.ListStacksParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation StacksListStacks has not yet been implemented")
		}),
		StateBackendsListStateBackendsHandler: state_backends.ListStateBackendsHandlerFunc(func(params state_backends.ListStateBackendsParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation StateBackendsListStateBackends has not yet been implemented")
		}),
		WorkspacesListWorkspacesHandler: workspaces.ListWorkspacesHandlerFunc(func(params workspaces.ListWorkspacesParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation WorkspacesListWorkspaces has not yet been implemented")
		}),
		StateBackendsUpdateStateBackendHandler: state_backends.UpdateStateBackendHandlerFunc(func(params state_backends.UpdateStateBackendParams, principal *models.ResourceAuthUser) middleware.Responder {
			return middleware.NotImplemented("operation StateBackendsUpdateStateBackend has not yet been implemented")
		}),

		// Applies when the "x-lunaform-auth" header is set
		APIKeyAuth: func(token string) (*models.ResourceAuthUser, error) {
			return nil, errors.NotImplemented("api key auth (api-key) x-lunaform-auth from header param [x-lunaform-auth] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*LunaformAPI This is a RESTful server for managing Terraform plan and apply jobs and the auditing of actions to approve those apply jobs.
The inspiration for this project is the AWS CloudFormation API's. The intention is to implement a locking mechanism
not only for the terraform state, but for the plan and apply of terraform modules. Once a `module` plan starts, it
is instantiated as a `stack` within the nomencalture of `lunaform`.
*/
type LunaformAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/vnd.lunaform.v1+json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/vnd.lunaform.v1+json" mime type
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-lunaform-auth provided in the header
	APIKeyAuth func(string) (*models.ResourceAuthUser, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ModulesCreateModuleHandler sets the operation handler for the create module operation
	ModulesCreateModuleHandler modules.CreateModuleHandler
	// StateBackendsCreateStateBackendHandler sets the operation handler for the create state backend operation
	StateBackendsCreateStateBackendHandler state_backends.CreateStateBackendHandler
	// WorkspacesCreateWorkspaceHandler sets the operation handler for the create workspace operation
	WorkspacesCreateWorkspaceHandler workspaces.CreateWorkspaceHandler
	// ModulesDeleteModuleHandler sets the operation handler for the delete module operation
	ModulesDeleteModuleHandler modules.DeleteModuleHandler
	// StacksDeployStackHandler sets the operation handler for the deploy stack operation
	StacksDeployStackHandler stacks.DeployStackHandler
	// WorkspacesDescribeWorkspaceHandler sets the operation handler for the describe workspace operation
	WorkspacesDescribeWorkspaceHandler workspaces.DescribeWorkspaceHandler
	// ModulesGetModuleHandler sets the operation handler for the get module operation
	ModulesGetModuleHandler modules.GetModuleHandler
	// StacksGetStackHandler sets the operation handler for the get stack operation
	StacksGetStackHandler stacks.GetStackHandler
	// StacksListDeploymentsHandler sets the operation handler for the list deployments operation
	StacksListDeploymentsHandler stacks.ListDeploymentsHandler
	// ModulesListModulesHandler sets the operation handler for the list modules operation
	ModulesListModulesHandler modules.ListModulesHandler
	// ResourcesListResourceGroupsHandler sets the operation handler for the list resource groups operation
	ResourcesListResourceGroupsHandler resources.ListResourceGroupsHandler
	// ResourcesListResourcesHandler sets the operation handler for the list resources operation
	ResourcesListResourcesHandler resources.ListResourcesHandler
	// StacksListStacksHandler sets the operation handler for the list stacks operation
	StacksListStacksHandler stacks.ListStacksHandler
	// StateBackendsListStateBackendsHandler sets the operation handler for the list state backends operation
	StateBackendsListStateBackendsHandler state_backends.ListStateBackendsHandler
	// WorkspacesListWorkspacesHandler sets the operation handler for the list workspaces operation
	WorkspacesListWorkspacesHandler workspaces.ListWorkspacesHandler
	// StateBackendsUpdateStateBackendHandler sets the operation handler for the update state backend operation
	StateBackendsUpdateStateBackendHandler state_backends.UpdateStateBackendHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *LunaformAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *LunaformAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *LunaformAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *LunaformAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *LunaformAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *LunaformAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *LunaformAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the LunaformAPI
func (o *LunaformAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "XLunaformAuthAuth")
	}

	if o.ModulesCreateModuleHandler == nil {
		unregistered = append(unregistered, "modules.CreateModuleHandler")
	}

	if o.StateBackendsCreateStateBackendHandler == nil {
		unregistered = append(unregistered, "state_backends.CreateStateBackendHandler")
	}

	if o.WorkspacesCreateWorkspaceHandler == nil {
		unregistered = append(unregistered, "workspaces.CreateWorkspaceHandler")
	}

	if o.ModulesDeleteModuleHandler == nil {
		unregistered = append(unregistered, "modules.DeleteModuleHandler")
	}

	if o.StacksDeployStackHandler == nil {
		unregistered = append(unregistered, "stacks.DeployStackHandler")
	}

	if o.WorkspacesDescribeWorkspaceHandler == nil {
		unregistered = append(unregistered, "workspaces.DescribeWorkspaceHandler")
	}

	if o.ModulesGetModuleHandler == nil {
		unregistered = append(unregistered, "modules.GetModuleHandler")
	}

	if o.StacksGetStackHandler == nil {
		unregistered = append(unregistered, "stacks.GetStackHandler")
	}

	if o.StacksListDeploymentsHandler == nil {
		unregistered = append(unregistered, "stacks.ListDeploymentsHandler")
	}

	if o.ModulesListModulesHandler == nil {
		unregistered = append(unregistered, "modules.ListModulesHandler")
	}

	if o.ResourcesListResourceGroupsHandler == nil {
		unregistered = append(unregistered, "resources.ListResourceGroupsHandler")
	}

	if o.ResourcesListResourcesHandler == nil {
		unregistered = append(unregistered, "resources.ListResourcesHandler")
	}

	if o.StacksListStacksHandler == nil {
		unregistered = append(unregistered, "stacks.ListStacksHandler")
	}

	if o.StateBackendsListStateBackendsHandler == nil {
		unregistered = append(unregistered, "state_backends.ListStateBackendsHandler")
	}

	if o.WorkspacesListWorkspacesHandler == nil {
		unregistered = append(unregistered, "workspaces.ListWorkspacesHandler")
	}

	if o.StateBackendsUpdateStateBackendHandler == nil {
		unregistered = append(unregistered, "state_backends.UpdateStateBackendHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *LunaformAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *LunaformAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "api-key":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.APIKeyAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *LunaformAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *LunaformAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/vnd.lunaform.v1+json":
			result["application/vnd.lunaform.v1+json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *LunaformAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/vnd.lunaform.v1+json":
			result["application/vnd.lunaform.v1+json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *LunaformAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the lunaform API
func (o *LunaformAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *LunaformAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tf/modules"] = modules.NewCreateModule(o.context, o.ModulesCreateModuleHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tf/state-backends"] = state_backends.NewCreateStateBackend(o.context, o.StateBackendsCreateStateBackendHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/tf/workspace/{name}"] = workspaces.NewCreateWorkspace(o.context, o.WorkspacesCreateWorkspaceHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tf/module/{id}"] = modules.NewDeleteModule(o.context, o.ModulesDeleteModuleHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tf/stacks"] = stacks.NewDeployStack(o.context, o.StacksDeployStackHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/workspace/{name}"] = workspaces.NewDescribeWorkspace(o.context, o.WorkspacesDescribeWorkspaceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/module/{id}"] = modules.NewGetModule(o.context, o.ModulesGetModuleHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/stack/{id}"] = stacks.NewGetStack(o.context, o.StacksGetStackHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/stack/{id}/deployments"] = stacks.NewListDeployments(o.context, o.StacksListDeploymentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/modules"] = modules.NewListModules(o.context, o.ModulesListModulesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = resources.NewListResourceGroups(o.context, o.ResourcesListResourceGroupsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{group}"] = resources.NewListResources(o.context, o.ResourcesListResourcesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/stacks"] = stacks.NewListStacks(o.context, o.StacksListStacksHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/state-backends"] = state_backends.NewListStateBackends(o.context, o.StateBackendsListStateBackendsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tf/workspaces"] = workspaces.NewListWorkspaces(o.context, o.WorkspacesListWorkspacesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/tf/state-backend/{id}"] = state_backends.NewUpdateStateBackend(o.context, o.StateBackendsUpdateStateBackendHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *LunaformAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *LunaformAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *LunaformAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *LunaformAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
