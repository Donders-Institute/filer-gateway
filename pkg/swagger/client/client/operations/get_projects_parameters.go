// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetProjectsParams creates a new GetProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectsParams() *GetProjectsParams {
	return &GetProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsParamsWithTimeout creates a new GetProjectsParams object
// with the ability to set a timeout on a request.
func NewGetProjectsParamsWithTimeout(timeout time.Duration) *GetProjectsParams {
	return &GetProjectsParams{
		timeout: timeout,
	}
}

// NewGetProjectsParamsWithContext creates a new GetProjectsParams object
// with the ability to set a context for a request.
func NewGetProjectsParamsWithContext(ctx context.Context) *GetProjectsParams {
	return &GetProjectsParams{
		Context: ctx,
	}
}

// NewGetProjectsParamsWithHTTPClient creates a new GetProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectsParamsWithHTTPClient(client *http.Client) *GetProjectsParams {
	return &GetProjectsParams{
		HTTPClient: client,
	}
}

/* GetProjectsParams contains all the parameters to send to the API endpoint
   for the get projects operation.

   Typically these are written to a http.Request.
*/
type GetProjectsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsParams) WithDefaults() *GetProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) WithTimeout(timeout time.Duration) *GetProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects params
func (o *GetProjectsParams) WithContext(ctx context.Context) *GetProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects params
func (o *GetProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) WithHTTPClient(client *http.Client) *GetProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
