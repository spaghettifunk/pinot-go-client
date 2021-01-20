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

// NewGetTableStatsParams creates a new GetTableStatsParams object
// with the default values initialized.
func NewGetTableStatsParams() *GetTableStatsParams {
	var ()
	return &GetTableStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTableStatsParamsWithTimeout creates a new GetTableStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTableStatsParamsWithTimeout(timeout time.Duration) *GetTableStatsParams {
	var ()
	return &GetTableStatsParams{

		timeout: timeout,
	}
}

// NewGetTableStatsParamsWithContext creates a new GetTableStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTableStatsParamsWithContext(ctx context.Context) *GetTableStatsParams {
	var ()
	return &GetTableStatsParams{

		Context: ctx,
	}
}

// NewGetTableStatsParamsWithHTTPClient creates a new GetTableStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTableStatsParamsWithHTTPClient(client *http.Client) *GetTableStatsParams {
	var ()
	return &GetTableStatsParams{
		HTTPClient: client,
	}
}

/*GetTableStatsParams contains all the parameters to send to the API endpoint
for the get table stats operation typically these are written to a http.Request
*/
type GetTableStatsParams struct {

	/*TableName
	  Name of the table

	*/
	TableName string
	/*Type
	  realtime|offline

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get table stats params
func (o *GetTableStatsParams) WithTimeout(timeout time.Duration) *GetTableStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get table stats params
func (o *GetTableStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get table stats params
func (o *GetTableStatsParams) WithContext(ctx context.Context) *GetTableStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get table stats params
func (o *GetTableStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get table stats params
func (o *GetTableStatsParams) WithHTTPClient(client *http.Client) *GetTableStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get table stats params
func (o *GetTableStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the get table stats params
func (o *GetTableStatsParams) WithTableName(tableName string) *GetTableStatsParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get table stats params
func (o *GetTableStatsParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the get table stats params
func (o *GetTableStatsParams) WithType(typeVar *string) *GetTableStatsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get table stats params
func (o *GetTableStatsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetTableStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
