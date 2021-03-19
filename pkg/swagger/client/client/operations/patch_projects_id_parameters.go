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

	"github.com/Donders-Institute/filer-gateway/pkg/swagger/client/models"
)

// NewPatchProjectsIDParams creates a new PatchProjectsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchProjectsIDParams() *PatchProjectsIDParams {
	return &PatchProjectsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchProjectsIDParamsWithTimeout creates a new PatchProjectsIDParams object
// with the ability to set a timeout on a request.
func NewPatchProjectsIDParamsWithTimeout(timeout time.Duration) *PatchProjectsIDParams {
	return &PatchProjectsIDParams{
		timeout: timeout,
	}
}

// NewPatchProjectsIDParamsWithContext creates a new PatchProjectsIDParams object
// with the ability to set a context for a request.
func NewPatchProjectsIDParamsWithContext(ctx context.Context) *PatchProjectsIDParams {
	return &PatchProjectsIDParams{
		Context: ctx,
	}
}

// NewPatchProjectsIDParamsWithHTTPClient creates a new PatchProjectsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchProjectsIDParamsWithHTTPClient(client *http.Client) *PatchProjectsIDParams {
	return &PatchProjectsIDParams{
		HTTPClient: client,
	}
}

/* PatchProjectsIDParams contains all the parameters to send to the API endpoint
   for the patch projects ID operation.

   Typically these are written to a http.Request.
*/
type PatchProjectsIDParams struct {

	/* ID.

	   project identifier
	*/
	ID string

	/* ProjectUpdateData.

	   data for project update
	*/
	ProjectUpdateData *models.RequestBodyProjectResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch projects ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchProjectsIDParams) WithDefaults() *PatchProjectsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch projects ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchProjectsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch projects ID params
func (o *PatchProjectsIDParams) WithTimeout(timeout time.Duration) *PatchProjectsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch projects ID params
func (o *PatchProjectsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch projects ID params
func (o *PatchProjectsIDParams) WithContext(ctx context.Context) *PatchProjectsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch projects ID params
func (o *PatchProjectsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch projects ID params
func (o *PatchProjectsIDParams) WithHTTPClient(client *http.Client) *PatchProjectsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch projects ID params
func (o *PatchProjectsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch projects ID params
func (o *PatchProjectsIDParams) WithID(id string) *PatchProjectsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch projects ID params
func (o *PatchProjectsIDParams) SetID(id string) {
	o.ID = id
}

// WithProjectUpdateData adds the projectUpdateData to the patch projects ID params
func (o *PatchProjectsIDParams) WithProjectUpdateData(projectUpdateData *models.RequestBodyProjectResource) *PatchProjectsIDParams {
	o.SetProjectUpdateData(projectUpdateData)
	return o
}

// SetProjectUpdateData adds the projectUpdateData to the patch projects ID params
func (o *PatchProjectsIDParams) SetProjectUpdateData(projectUpdateData *models.RequestBodyProjectResource) {
	o.ProjectUpdateData = projectUpdateData
}

// WriteToRequest writes these params to a swagger request
func (o *PatchProjectsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.ProjectUpdateData != nil {
		if err := r.SetBodyParam(o.ProjectUpdateData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
