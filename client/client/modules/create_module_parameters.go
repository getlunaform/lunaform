// Code generated by go-swagger; DO NOT EDIT.

package modules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/drewsonne/lunaform/server/models"
)

// NewCreateModuleParams creates a new CreateModuleParams object
// with the default values initialized.
func NewCreateModuleParams() *CreateModuleParams {
	var ()
	return &CreateModuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateModuleParamsWithTimeout creates a new CreateModuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateModuleParamsWithTimeout(timeout time.Duration) *CreateModuleParams {
	var ()
	return &CreateModuleParams{

		timeout: timeout,
	}
}

// NewCreateModuleParamsWithContext creates a new CreateModuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateModuleParamsWithContext(ctx context.Context) *CreateModuleParams {
	var ()
	return &CreateModuleParams{

		Context: ctx,
	}
}

// NewCreateModuleParamsWithHTTPClient creates a new CreateModuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateModuleParamsWithHTTPClient(client *http.Client) *CreateModuleParams {
	var ()
	return &CreateModuleParams{
		HTTPClient: client,
	}
}

/*CreateModuleParams contains all the parameters to send to the API endpoint
for the create module operation typically these are written to a http.Request
*/
type CreateModuleParams struct {

	/*TerraformModule
	  A terraform module

	*/
	TerraformModule *models.ResourceTfModule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create module params
func (o *CreateModuleParams) WithTimeout(timeout time.Duration) *CreateModuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create module params
func (o *CreateModuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create module params
func (o *CreateModuleParams) WithContext(ctx context.Context) *CreateModuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create module params
func (o *CreateModuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create module params
func (o *CreateModuleParams) WithHTTPClient(client *http.Client) *CreateModuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create module params
func (o *CreateModuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTerraformModule adds the terraformModule to the create module params
func (o *CreateModuleParams) WithTerraformModule(terraformModule *models.ResourceTfModule) *CreateModuleParams {
	o.SetTerraformModule(terraformModule)
	return o
}

// SetTerraformModule adds the terraformModule to the create module params
func (o *CreateModuleParams) SetTerraformModule(terraformModule *models.ResourceTfModule) {
	o.TerraformModule = terraformModule
}

// WriteToRequest writes these params to a swagger request
func (o *CreateModuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TerraformModule != nil {
		if err := r.SetBodyParam(o.TerraformModule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
