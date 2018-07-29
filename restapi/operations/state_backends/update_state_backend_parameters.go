// Code generated by go-swagger; DO NOT EDIT.

package state_backends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/getlunaform/lunaform/models"
)

// NewUpdateStateBackendParams creates a new UpdateStateBackendParams object
// no default values defined in spec.
func NewUpdateStateBackendParams() UpdateStateBackendParams {

	return UpdateStateBackendParams{}
}

// UpdateStateBackendParams contains all the bound params for the update state backend operation
// typically these are obtained from a http.Request
//
// swagger:parameters update-state-backend
type UpdateStateBackendParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of a terraform state backend
	  Required: true
	  In: path
	*/
	ID string
	/*A terraform state backend
	  In: body
	*/
	TerraformStateBackend *models.ResourceTfStateBackend
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateStateBackendParams() beforehand.
func (o *UpdateStateBackendParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ResourceTfStateBackend
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("terraformStateBackend", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.TerraformStateBackend = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UpdateStateBackendParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}