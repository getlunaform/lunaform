// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListResourcesParams creates a new ListResourcesParams object
// no default values defined in spec.
func NewListResourcesParams() ListResourcesParams {

	return ListResourcesParams{}
}

// ListResourcesParams contains all the bound params for the list resources operation
// typically these are obtained from a http.Request
//
// swagger:parameters list-resources
type ListResourcesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Root level resources
	  Required: true
	  In: path
	*/
	Group string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListResourcesParams() beforehand.
func (o *ListResourcesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGroup, rhkGroup, _ := route.Params.GetOK("group")
	if err := o.bindGroup(rGroup, rhkGroup, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGroup binds and validates parameter Group from path.
func (o *ListResourcesParams) bindGroup(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Group = raw

	if err := o.validateGroup(formats); err != nil {
		return err
	}

	return nil
}

// validateGroup carries on validations for parameter Group
func (o *ListResourcesParams) validateGroup(formats strfmt.Registry) error {

	if err := validate.Enum("group", "path", o.Group, []interface{}{"tf", "identity", "vcs"}); err != nil {
		return err
	}

	return nil
}
