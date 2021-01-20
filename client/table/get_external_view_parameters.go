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

// NewGetExternalViewParams creates a new GetExternalViewParams object
// with the default values initialized.
func NewGetExternalViewParams() *GetExternalViewParams {
	var ()
	return &GetExternalViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalViewParamsWithTimeout creates a new GetExternalViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalViewParamsWithTimeout(timeout time.Duration) *GetExternalViewParams {
	var ()
	return &GetExternalViewParams{

		timeout: timeout,
	}
}

// NewGetExternalViewParamsWithContext creates a new GetExternalViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalViewParamsWithContext(ctx context.Context) *GetExternalViewParams {
	var ()
	return &GetExternalViewParams{

		Context: ctx,
	}
}

// NewGetExternalViewParamsWithHTTPClient creates a new GetExternalViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalViewParamsWithHTTPClient(client *http.Client) *GetExternalViewParams {
	var ()
	return &GetExternalViewParams{
		HTTPClient: client,
	}
}

/*GetExternalViewParams contains all the parameters to send to the API endpoint
for the get external view operation typically these are written to a http.Request
*/
type GetExternalViewParams struct {

	/*TableName
	  Name of the table

	*/
	TableName string
	/*TableType
	  realtime|offline

	*/
	TableType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get external view params
func (o *GetExternalViewParams) WithTimeout(timeout time.Duration) *GetExternalViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external view params
func (o *GetExternalViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external view params
func (o *GetExternalViewParams) WithContext(ctx context.Context) *GetExternalViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external view params
func (o *GetExternalViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external view params
func (o *GetExternalViewParams) WithHTTPClient(client *http.Client) *GetExternalViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external view params
func (o *GetExternalViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the get external view params
func (o *GetExternalViewParams) WithTableName(tableName string) *GetExternalViewParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get external view params
func (o *GetExternalViewParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithTableType adds the tableType to the get external view params
func (o *GetExternalViewParams) WithTableType(tableType *string) *GetExternalViewParams {
	o.SetTableType(tableType)
	return o
}

// SetTableType adds the tableType to the get external view params
func (o *GetExternalViewParams) SetTableType(tableType *string) {
	o.TableType = tableType
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if o.TableType != nil {

		// query param tableType
		var qrTableType string
		if o.TableType != nil {
			qrTableType = *o.TableType
		}
		qTableType := qrTableType
		if qTableType != "" {
			if err := r.SetQueryParam("tableType", qTableType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}