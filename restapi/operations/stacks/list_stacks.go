// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/getlunaform/lunaform/models"
)

// ListStacksHandlerFunc turns a function with the right signature into a list stacks handler
type ListStacksHandlerFunc func(ListStacksParams, *models.ResourceAuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStacksHandlerFunc) Handle(params ListStacksParams, principal *models.ResourceAuthUser) middleware.Responder {
	return fn(params, principal)
}

// ListStacksHandler interface for that can handle valid list stacks params
type ListStacksHandler interface {
	Handle(ListStacksParams, *models.ResourceAuthUser) middleware.Responder
}

// NewListStacks creates a new http.Handler for the list stacks operation
func NewListStacks(ctx *middleware.Context, handler ListStacksHandler) *ListStacks {
	return &ListStacks{Context: ctx, Handler: handler}
}

/*ListStacks swagger:route GET /tf/stacks stacks listStacks

List deployed TF modules

*/
type ListStacks struct {
	Context *middleware.Context
	Handler ListStacksHandler
}

func (o *ListStacks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStacksParams()

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