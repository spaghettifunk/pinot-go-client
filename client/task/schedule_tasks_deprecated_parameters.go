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

// NewScheduleTasksDeprecatedParams creates a new ScheduleTasksDeprecatedParams object
// with the default values initialized.
func NewScheduleTasksDeprecatedParams() *ScheduleTasksDeprecatedParams {

	return &ScheduleTasksDeprecatedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleTasksDeprecatedParamsWithTimeout creates a new ScheduleTasksDeprecatedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScheduleTasksDeprecatedParamsWithTimeout(timeout time.Duration) *ScheduleTasksDeprecatedParams {

	return &ScheduleTasksDeprecatedParams{

		timeout: timeout,
	}
}

// NewScheduleTasksDeprecatedParamsWithContext creates a new ScheduleTasksDeprecatedParams object
// with the default values initialized, and the ability to set a context for a request
func NewScheduleTasksDeprecatedParamsWithContext(ctx context.Context) *ScheduleTasksDeprecatedParams {

	return &ScheduleTasksDeprecatedParams{

		Context: ctx,
	}
}

// NewScheduleTasksDeprecatedParamsWithHTTPClient creates a new ScheduleTasksDeprecatedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScheduleTasksDeprecatedParamsWithHTTPClient(client *http.Client) *ScheduleTasksDeprecatedParams {

	return &ScheduleTasksDeprecatedParams{
		HTTPClient: client,
	}
}

/*ScheduleTasksDeprecatedParams contains all the parameters to send to the API endpoint
for the schedule tasks deprecated operation typically these are written to a http.Request
*/
type ScheduleTasksDeprecatedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schedule tasks deprecated params
func (o *ScheduleTasksDeprecatedParams) WithTimeout(timeout time.Duration) *ScheduleTasksDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule tasks deprecated params
func (o *ScheduleTasksDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule tasks deprecated params
func (o *ScheduleTasksDeprecatedParams) WithContext(ctx context.Context) *ScheduleTasksDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule tasks deprecated params
func (o *ScheduleTasksDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule tasks deprecated params
func (o *ScheduleTasksDeprecatedParams) WithHTTPClient(client *http.Client) *ScheduleTasksDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule tasks deprecated params
func (o *ScheduleTasksDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleTasksDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
