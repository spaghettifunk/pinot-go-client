// Code generated by go-swagger; DO NOT EDIT.

package task

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

// NewGetTaskQueuesParams creates a new GetTaskQueuesParams object
// with the default values initialized.
func NewGetTaskQueuesParams() *GetTaskQueuesParams {

	return &GetTaskQueuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskQueuesParamsWithTimeout creates a new GetTaskQueuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTaskQueuesParamsWithTimeout(timeout time.Duration) *GetTaskQueuesParams {

	return &GetTaskQueuesParams{

		timeout: timeout,
	}
}

// NewGetTaskQueuesParamsWithContext creates a new GetTaskQueuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTaskQueuesParamsWithContext(ctx context.Context) *GetTaskQueuesParams {

	return &GetTaskQueuesParams{

		Context: ctx,
	}
}

// NewGetTaskQueuesParamsWithHTTPClient creates a new GetTaskQueuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTaskQueuesParamsWithHTTPClient(client *http.Client) *GetTaskQueuesParams {

	return &GetTaskQueuesParams{
		HTTPClient: client,
	}
}

/*GetTaskQueuesParams contains all the parameters to send to the API endpoint
for the get task queues operation typically these are written to a http.Request
*/
type GetTaskQueuesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get task queues params
func (o *GetTaskQueuesParams) WithTimeout(timeout time.Duration) *GetTaskQueuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task queues params
func (o *GetTaskQueuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task queues params
func (o *GetTaskQueuesParams) WithContext(ctx context.Context) *GetTaskQueuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task queues params
func (o *GetTaskQueuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task queues params
func (o *GetTaskQueuesParams) WithHTTPClient(client *http.Client) *GetTaskQueuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task queues params
func (o *GetTaskQueuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskQueuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}