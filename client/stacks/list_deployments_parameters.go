// Code generated by go-swagger; DO NOT EDIT.

package stacks

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
)

// NewListDeploymentsParams creates a new ListDeploymentsParams object
// with the default values initialized.
func NewListDeploymentsParams() *ListDeploymentsParams {
	var ()
	return &ListDeploymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListDeploymentsParamsWithTimeout creates a new ListDeploymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDeploymentsParamsWithTimeout(timeout time.Duration) *ListDeploymentsParams {
	var ()
	return &ListDeploymentsParams{

		timeout: timeout,
	}
}

// NewListDeploymentsParamsWithContext creates a new ListDeploymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListDeploymentsParamsWithContext(ctx context.Context) *ListDeploymentsParams {
	var ()
	return &ListDeploymentsParams{

		Context: ctx,
	}
}

// NewListDeploymentsParamsWithHTTPClient creates a new ListDeploymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDeploymentsParamsWithHTTPClient(client *http.Client) *ListDeploymentsParams {
	var ()
	return &ListDeploymentsParams{
		HTTPClient: client,
	}
}

/*ListDeploymentsParams contains all the parameters to send to the API endpoint
for the list deployments operation typically these are written to a http.Request
*/
type ListDeploymentsParams struct {

	/*ID
	  Unique identifier for this stack

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list deployments params
func (o *ListDeploymentsParams) WithTimeout(timeout time.Duration) *ListDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list deployments params
func (o *ListDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list deployments params
func (o *ListDeploymentsParams) WithContext(ctx context.Context) *ListDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list deployments params
func (o *ListDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list deployments params
func (o *ListDeploymentsParams) WithHTTPClient(client *http.Client) *ListDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list deployments params
func (o *ListDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list deployments params
func (o *ListDeploymentsParams) WithID(id string) *ListDeploymentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list deployments params
func (o *ListDeploymentsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}