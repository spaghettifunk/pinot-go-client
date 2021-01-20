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

// NewGetTaskConfigsDeprecatedParams creates a new GetTaskConfigsDeprecatedParams object
// with the default values initialized.
func NewGetTaskConfigsDeprecatedParams() *GetTaskConfigsDeprecatedParams {
	var ()
	return &GetTaskConfigsDeprecatedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskConfigsDeprecatedParamsWithTimeout creates a new GetTaskConfigsDeprecatedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTaskConfigsDeprecatedParamsWithTimeout(timeout time.Duration) *GetTaskConfigsDeprecatedParams {
	var ()
	return &GetTaskConfigsDeprecatedParams{

		timeout: timeout,
	}
}

// NewGetTaskConfigsDeprecatedParamsWithContext creates a new GetTaskConfigsDeprecatedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTaskConfigsDeprecatedParamsWithContext(ctx context.Context) *GetTaskConfigsDeprecatedParams {
	var ()
	return &GetTaskConfigsDeprecatedParams{

		Context: ctx,
	}
}

// NewGetTaskConfigsDeprecatedParamsWithHTTPClient creates a new GetTaskConfigsDeprecatedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTaskConfigsDeprecatedParamsWithHTTPClient(client *http.Client) *GetTaskConfigsDeprecatedParams {
	var ()
	return &GetTaskConfigsDeprecatedParams{
		HTTPClient: client,
	}
}

/*GetTaskConfigsDeprecatedParams contains all the parameters to send to the API endpoint
for the get task configs deprecated operation typically these are written to a http.Request
*/
type GetTaskConfigsDeprecatedParams struct {

	/*TaskName
	  Task name

	*/
	TaskName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) WithTimeout(timeout time.Duration) *GetTaskConfigsDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) WithContext(ctx context.Context) *GetTaskConfigsDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) WithHTTPClient(client *http.Client) *GetTaskConfigsDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskName adds the taskName to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) WithTaskName(taskName string) *GetTaskConfigsDeprecatedParams {
	o.SetTaskName(taskName)
	return o
}

// SetTaskName adds the taskName to the get task configs deprecated params
func (o *GetTaskConfigsDeprecatedParams) SetTaskName(taskName string) {
	o.TaskName = taskName
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskConfigsDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param taskName
	if err := r.SetPathParam("taskName", o.TaskName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
