// Code generated by go-swagger; DO NOT EDIT.

package table

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

// NewUpdateTableConfigParams creates a new UpdateTableConfigParams object
// with the default values initialized.
func NewUpdateTableConfigParams() *UpdateTableConfigParams {
	var ()
	return &UpdateTableConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTableConfigParamsWithTimeout creates a new UpdateTableConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTableConfigParamsWithTimeout(timeout time.Duration) *UpdateTableConfigParams {
	var ()
	return &UpdateTableConfigParams{

		timeout: timeout,
	}
}

// NewUpdateTableConfigParamsWithContext creates a new UpdateTableConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTableConfigParamsWithContext(ctx context.Context) *UpdateTableConfigParams {
	var ()
	return &UpdateTableConfigParams{

		Context: ctx,
	}
}

// NewUpdateTableConfigParamsWithHTTPClient creates a new UpdateTableConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTableConfigParamsWithHTTPClient(client *http.Client) *UpdateTableConfigParams {
	var ()
	return &UpdateTableConfigParams{
		HTTPClient: client,
	}
}

/*UpdateTableConfigParams contains all the parameters to send to the API endpoint
for the update table config operation typically these are written to a http.Request
*/
type UpdateTableConfigParams struct {

	/*Body*/
	Body string
	/*TableName
	  Name of the table to update

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update table config params
func (o *UpdateTableConfigParams) WithTimeout(timeout time.Duration) *UpdateTableConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update table config params
func (o *UpdateTableConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update table config params
func (o *UpdateTableConfigParams) WithContext(ctx context.Context) *UpdateTableConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update table config params
func (o *UpdateTableConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update table config params
func (o *UpdateTableConfigParams) WithHTTPClient(client *http.Client) *UpdateTableConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update table config params
func (o *UpdateTableConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update table config params
func (o *UpdateTableConfigParams) WithBody(body string) *UpdateTableConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update table config params
func (o *UpdateTableConfigParams) SetBody(body string) {
	o.Body = body
}

// WithTableName adds the tableName to the update table config params
func (o *UpdateTableConfigParams) WithTableName(tableName string) *UpdateTableConfigParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the update table config params
func (o *UpdateTableConfigParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTableConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
