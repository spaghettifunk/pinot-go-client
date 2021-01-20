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

// NewReloadAllSegmentsDeprecated1Params creates a new ReloadAllSegmentsDeprecated1Params object
// with the default values initialized.
func NewReloadAllSegmentsDeprecated1Params() *ReloadAllSegmentsDeprecated1Params {
	var ()
	return &ReloadAllSegmentsDeprecated1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewReloadAllSegmentsDeprecated1ParamsWithTimeout creates a new ReloadAllSegmentsDeprecated1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewReloadAllSegmentsDeprecated1ParamsWithTimeout(timeout time.Duration) *ReloadAllSegmentsDeprecated1Params {
	var ()
	return &ReloadAllSegmentsDeprecated1Params{

		timeout: timeout,
	}
}

// NewReloadAllSegmentsDeprecated1ParamsWithContext creates a new ReloadAllSegmentsDeprecated1Params object
// with the default values initialized, and the ability to set a context for a request
func NewReloadAllSegmentsDeprecated1ParamsWithContext(ctx context.Context) *ReloadAllSegmentsDeprecated1Params {
	var ()
	return &ReloadAllSegmentsDeprecated1Params{

		Context: ctx,
	}
}

// NewReloadAllSegmentsDeprecated1ParamsWithHTTPClient creates a new ReloadAllSegmentsDeprecated1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReloadAllSegmentsDeprecated1ParamsWithHTTPClient(client *http.Client) *ReloadAllSegmentsDeprecated1Params {
	var ()
	return &ReloadAllSegmentsDeprecated1Params{
		HTTPClient: client,
	}
}

/*ReloadAllSegmentsDeprecated1Params contains all the parameters to send to the API endpoint
for the reload all segments deprecated1 operation typically these are written to a http.Request
*/
type ReloadAllSegmentsDeprecated1Params struct {

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

// WithTimeout adds the timeout to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) WithTimeout(timeout time.Duration) *ReloadAllSegmentsDeprecated1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) WithContext(ctx context.Context) *ReloadAllSegmentsDeprecated1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) WithHTTPClient(client *http.Client) *ReloadAllSegmentsDeprecated1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) WithTableName(tableName string) *ReloadAllSegmentsDeprecated1Params {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) WithType(typeVar *string) *ReloadAllSegmentsDeprecated1Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the reload all segments deprecated1 params
func (o *ReloadAllSegmentsDeprecated1Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReloadAllSegmentsDeprecated1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
