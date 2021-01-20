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

// NewGetTaskQueueStateParams creates a new GetTaskQueueStateParams object
// with the default values initialized.
func NewGetTaskQueueStateParams() *GetTaskQueueStateParams {
	var ()
	return &GetTaskQueueStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskQueueStateParamsWithTimeout creates a new GetTaskQueueStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTaskQueueStateParamsWithTimeout(timeout time.Duration) *GetTaskQueueStateParams {
	var ()
	return &GetTaskQueueStateParams{

		timeout: timeout,
	}
}

// NewGetTaskQueueStateParamsWithContext creates a new GetTaskQueueStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTaskQueueStateParamsWithContext(ctx context.Context) *GetTaskQueueStateParams {
	var ()
	return &GetTaskQueueStateParams{

		Context: ctx,
	}
}

// NewGetTaskQueueStateParamsWithHTTPClient creates a new GetTaskQueueStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTaskQueueStateParamsWithHTTPClient(client *http.Client) *GetTaskQueueStateParams {
	var ()
	return &GetTaskQueueStateParams{
		HTTPClient: client,
	}
}

/*GetTaskQueueStateParams contains all the parameters to send to the API endpoint
for the get task queue state operation typically these are written to a http.Request
*/
type GetTaskQueueStateParams struct {

	/*TaskType
	  Task type

	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get task queue state params
func (o *GetTaskQueueStateParams) WithTimeout(timeout time.Duration) *GetTaskQueueStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task queue state params
func (o *GetTaskQueueStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task queue state params
func (o *GetTaskQueueStateParams) WithContext(ctx context.Context) *GetTaskQueueStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task queue state params
func (o *GetTaskQueueStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task queue state params
func (o *GetTaskQueueStateParams) WithHTTPClient(client *http.Client) *GetTaskQueueStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task queue state params
func (o *GetTaskQueueStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskType adds the taskType to the get task queue state params
func (o *GetTaskQueueStateParams) WithTaskType(taskType string) *GetTaskQueueStateParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the get task queue state params
func (o *GetTaskQueueStateParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskQueueStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param taskType
	if err := r.SetPathParam("taskType", o.TaskType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
