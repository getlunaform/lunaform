// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/getlunaform/lunaform/models"
)

// ListProvidersHandlerFunc turns a function with the right signature into a list providers handler
type ListProvidersHandlerFunc func(ListProvidersParams, *models.ResourceAuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn ListProvidersHandlerFunc) Handle(params ListProvidersParams, principal *models.ResourceAuthUser) middleware.Responder {
	return fn(params, principal)
}

// ListProvidersHandler interface for that can handle valid list providers params
type ListProvidersHandler interface {
	Handle(ListProvidersParams, *models.ResourceAuthUser) middleware.Responder
}

// NewListProviders creates a new http.Handler for the list providers operation
func NewListProviders(ctx *middleware.Context, handler ListProvidersHandler) *ListProviders {
	return &ListProviders{Context: ctx, Handler: handler}
}

/*ListProviders swagger:route GET /tf/providers/ providers listProviders

List Terraform Providers

*/
type ListProviders struct {
	Context *middleware.Context
	Handler ListProvidersHandler
}

func (o *ListProviders) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListProvidersParams()

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
