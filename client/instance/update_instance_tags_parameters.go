// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// NewUpdateInstanceTagsParams creates a new UpdateInstanceTagsParams object
// with the default values initialized.
func NewUpdateInstanceTagsParams() *UpdateInstanceTagsParams {
	var ()
	return &UpdateInstanceTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInstanceTagsParamsWithTimeout creates a new UpdateInstanceTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInstanceTagsParamsWithTimeout(timeout time.Duration) *UpdateInstanceTagsParams {
	var ()
	return &UpdateInstanceTagsParams{

		timeout: timeout,
	}
}

// NewUpdateInstanceTagsParamsWithContext creates a new UpdateInstanceTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInstanceTagsParamsWithContext(ctx context.Context) *UpdateInstanceTagsParams {
	var ()
	return &UpdateInstanceTagsParams{

		Context: ctx,
	}
}

// NewUpdateInstanceTagsParamsWithHTTPClient creates a new UpdateInstanceTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInstanceTagsParamsWithHTTPClient(client *http.Client) *UpdateInstanceTagsParams {
	var ()
	return &UpdateInstanceTagsParams{
		HTTPClient: client,
	}
}

/*UpdateInstanceTagsParams contains all the parameters to send to the API endpoint
for the update instance tags operation typically these are written to a http.Request
*/
type UpdateInstanceTagsParams struct {

	/*InstanceName
	  Instance name

	*/
	InstanceName string
	/*Tags
	  Comma separated tags list

	*/
	Tags string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update instance tags params
func (o *UpdateInstanceTagsParams) WithTimeout(timeout time.Duration) *UpdateInstanceTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update instance tags params
func (o *UpdateInstanceTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update instance tags params
func (o *UpdateInstanceTagsParams) WithContext(ctx context.Context) *UpdateInstanceTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update instance tags params
func (o *UpdateInstanceTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update instance tags params
func (o *UpdateInstanceTagsParams) WithHTTPClient(client *http.Client) *UpdateInstanceTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update instance tags params
func (o *UpdateInstanceTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstanceName adds the instanceName to the update instance tags params
func (o *UpdateInstanceTagsParams) WithInstanceName(instanceName string) *UpdateInstanceTagsParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the update instance tags params
func (o *UpdateInstanceTagsParams) SetInstanceName(instanceName string) {
	o.InstanceName = instanceName
}

// WithTags adds the tags to the update instance tags params
func (o *UpdateInstanceTagsParams) WithTags(tags string) *UpdateInstanceTagsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the update instance tags params
func (o *UpdateInstanceTagsParams) SetTags(tags string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInstanceTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instanceName
	if err := r.SetPathParam("instanceName", o.InstanceName); err != nil {
		return err
	}

	// query param tags
	qrTags := o.Tags
	qTags := qrTags
	if qTags != "" {
		if err := r.SetQueryParam("tags", qTags); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}