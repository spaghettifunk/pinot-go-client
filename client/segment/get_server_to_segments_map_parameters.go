// Code generated by go-swagger; DO NOT EDIT.

package segment

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

// NewGetServerToSegmentsMapParams creates a new GetServerToSegmentsMapParams object
// with the default values initialized.
func NewGetServerToSegmentsMapParams() *GetServerToSegmentsMapParams {
	var ()
	return &GetServerToSegmentsMapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerToSegmentsMapParamsWithTimeout creates a new GetServerToSegmentsMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServerToSegmentsMapParamsWithTimeout(timeout time.Duration) *GetServerToSegmentsMapParams {
	var ()
	return &GetServerToSegmentsMapParams{

		timeout: timeout,
	}
}

// NewGetServerToSegmentsMapParamsWithContext creates a new GetServerToSegmentsMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServerToSegmentsMapParamsWithContext(ctx context.Context) *GetServerToSegmentsMapParams {
	var ()
	return &GetServerToSegmentsMapParams{

		Context: ctx,
	}
}

// NewGetServerToSegmentsMapParamsWithHTTPClient creates a new GetServerToSegmentsMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServerToSegmentsMapParamsWithHTTPClient(client *http.Client) *GetServerToSegmentsMapParams {
	var ()
	return &GetServerToSegmentsMapParams{
		HTTPClient: client,
	}
}

/*GetServerToSegmentsMapParams contains all the parameters to send to the API endpoint
for the get server to segments map operation typically these are written to a http.Request
*/
type GetServerToSegmentsMapParams struct {

	/*TableName
	  Name of the table

	*/
	TableName string
	/*Type
	  OFFLINE|REALTIME

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get server to segments map params
func (o *GetServerToSegmentsMapParams) WithTimeout(timeout time.Duration) *GetServerToSegmentsMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server to segments map params
func (o *GetServerToSegmentsMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server to segments map params
func (o *GetServerToSegmentsMapParams) WithContext(ctx context.Context) *GetServerToSegmentsMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server to segments map params
func (o *GetServerToSegmentsMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server to segments map params
func (o *GetServerToSegmentsMapParams) WithHTTPClient(client *http.Client) *GetServerToSegmentsMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server to segments map params
func (o *GetServerToSegmentsMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the get server to segments map params
func (o *GetServerToSegmentsMapParams) WithTableName(tableName string) *GetServerToSegmentsMapParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get server to segments map params
func (o *GetServerToSegmentsMapParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the get server to segments map params
func (o *GetServerToSegmentsMapParams) WithType(typeVar *string) *GetServerToSegmentsMapParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get server to segments map params
func (o *GetServerToSegmentsMapParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerToSegmentsMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}