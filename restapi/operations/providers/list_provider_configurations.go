// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/getlunaform/lunaform/models"
)

// ListProviderConfigurationsHandlerFunc turns a function with the right signature into a list provider configurations handler
type ListProviderConfigurationsHandlerFunc func(ListProviderConfigurationsParams, *models.ResourceAuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn ListProviderConfigurationsHandlerFunc) Handle(params ListProviderConfigurationsParams, principal *models.ResourceAuthUser) middleware.Responder {
	return fn(params, principal)
}

// ListProviderConfigurationsHandler interface for that can handle valid list provider configurations params
type ListProviderConfigurationsHandler interface {
	Handle(ListProviderConfigurationsParams, *models.ResourceAuthUser) middleware.Responder
}

// NewListProviderConfigurations creates a new http.Handler for the list provider configurations operation
func NewListProviderConfigurations(ctx *middleware.Context, handler ListProviderConfigurationsHandler) *ListProviderConfigurations {
	return &ListProviderConfigurations{Context: ctx, Handler: handler}
}

/*ListProviderConfigurations swagger:route GET /tf/provider/{provider-name}/configurations providers listProviderConfigurations

List Configurations for s Terraform Provider

*/
type ListProviderConfigurations struct {
	Context *middleware.Context
	Handler ListProviderConfigurationsHandler
}

func (o *ListProviderConfigurations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListProviderConfigurationsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.ResourceAuthUser
	if uprinc != nil {
		principal = uprinc.(*models.ResourceAuthUser) // this is really a models.ResourceAuthUser, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
