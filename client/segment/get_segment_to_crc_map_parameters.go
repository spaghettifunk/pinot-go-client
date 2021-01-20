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

// NewGetSegmentToCrcMapParams creates a new GetSegmentToCrcMapParams object
// with the default values initialized.
func NewGetSegmentToCrcMapParams() *GetSegmentToCrcMapParams {
	var ()
	return &GetSegmentToCrcMapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSegmentToCrcMapParamsWithTimeout creates a new GetSegmentToCrcMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSegmentToCrcMapParamsWithTimeout(timeout time.Duration) *GetSegmentToCrcMapParams {
	var ()
	return &GetSegmentToCrcMapParams{

		timeout: timeout,
	}
}

// NewGetSegmentToCrcMapParamsWithContext creates a new GetSegmentToCrcMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSegmentToCrcMapParamsWithContext(ctx context.Context) *GetSegmentToCrcMapParams {
	var ()
	return &GetSegmentToCrcMapParams{

		Context: ctx,
	}
}

// NewGetSegmentToCrcMapParamsWithHTTPClient creates a new GetSegmentToCrcMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSegmentToCrcMapParamsWithHTTPClient(client *http.Client) *GetSegmentToCrcMapParams {
	var ()
	return &GetSegmentToCrcMapParams{
		HTTPClient: client,
	}
}

/*GetSegmentToCrcMapParams contains all the parameters to send to the API endpoint
for the get segment to crc map operation typically these are written to a http.Request
*/
type GetSegmentToCrcMapParams struct {

	/*TableName
	  Name of the table

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) WithTimeout(timeout time.Duration) *GetSegmentToCrcMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) WithContext(ctx context.Context) *GetSegmentToCrcMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) WithHTTPClient(client *http.Client) *GetSegmentToCrcMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) WithTableName(tableName string) *GetSegmentToCrcMapParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get segment to crc map params
func (o *GetSegmentToCrcMapParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *GetSegmentToCrcMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
