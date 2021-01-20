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
	"github.com/go-openapi/swag"
)

// NewGetTableSizeParams creates a new GetTableSizeParams object
// with the default values initialized.
func NewGetTableSizeParams() *GetTableSizeParams {
	var (
		detailedDefault = bool(true)
	)
	return &GetTableSizeParams{
		Detailed: &detailedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTableSizeParamsWithTimeout creates a new GetTableSizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTableSizeParamsWithTimeout(timeout time.Duration) *GetTableSizeParams {
	var (
		detailedDefault = bool(true)
	)
	return &GetTableSizeParams{
		Detailed: &detailedDefault,

		timeout: timeout,
	}
}

// NewGetTableSizeParamsWithContext creates a new GetTableSizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTableSizeParamsWithContext(ctx context.Context) *GetTableSizeParams {
	var (
		detailedDefault = bool(true)
	)
	return &GetTableSizeParams{
		Detailed: &detailedDefault,

		Context: ctx,
	}
}

// NewGetTableSizeParamsWithHTTPClient creates a new GetTableSizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTableSizeParamsWithHTTPClient(client *http.Client) *GetTableSizeParams {
	var (
		detailedDefault = bool(true)
	)
	return &GetTableSizeParams{
		Detailed:   &detailedDefault,
		HTTPClient: client,
	}
}

/*GetTableSizeParams contains all the parameters to send to the API endpoint
for the get table size operation typically these are written to a http.Request
*/
type GetTableSizeParams struct {

	/*Detailed
	  Get detailed information

	*/
	Detailed *bool
	/*TableName
	  Table name without type

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get table size params
func (o *GetTableSizeParams) WithTimeout(timeout time.Duration) *GetTableSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get table size params
func (o *GetTableSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get table size params
func (o *GetTableSizeParams) WithContext(ctx context.Context) *GetTableSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get table size params
func (o *GetTableSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get table size params
func (o *GetTableSizeParams) WithHTTPClient(client *http.Client) *GetTableSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get table size params
func (o *GetTableSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetailed adds the detailed to the get table size params
func (o *GetTableSizeParams) WithDetailed(detailed *bool) *GetTableSizeParams {
	o.SetDetailed(detailed)
	return o
}

// SetDetailed adds the detailed to the get table size params
func (o *GetTableSizeParams) SetDetailed(detailed *bool) {
	o.Detailed = detailed
}

// WithTableName adds the tableName to the get table size params
func (o *GetTableSizeParams) WithTableName(tableName string) *GetTableSizeParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get table size params
func (o *GetTableSizeParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *GetTableSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Detailed != nil {

		// query param detailed
		var qrDetailed bool
		if o.Detailed != nil {
			qrDetailed = *o.Detailed
		}
		qDetailed := swag.FormatBool(qrDetailed)
		if qDetailed != "" {
			if err := r.SetQueryParam("detailed", qDetailed); err != nil {
				return err
			}
		}

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
