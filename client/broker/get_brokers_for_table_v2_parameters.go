// Code generated by go-swagger; DO NOT EDIT.

package broker

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

// NewGetBrokersForTableV2Params creates a new GetBrokersForTableV2Params object
// with the default values initialized.
func NewGetBrokersForTableV2Params() *GetBrokersForTableV2Params {
	var ()
	return &GetBrokersForTableV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBrokersForTableV2ParamsWithTimeout creates a new GetBrokersForTableV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBrokersForTableV2ParamsWithTimeout(timeout time.Duration) *GetBrokersForTableV2Params {
	var ()
	return &GetBrokersForTableV2Params{

		timeout: timeout,
	}
}

// NewGetBrokersForTableV2ParamsWithContext creates a new GetBrokersForTableV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetBrokersForTableV2ParamsWithContext(ctx context.Context) *GetBrokersForTableV2Params {
	var ()
	return &GetBrokersForTableV2Params{

		Context: ctx,
	}
}

// NewGetBrokersForTableV2ParamsWithHTTPClient creates a new GetBrokersForTableV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBrokersForTableV2ParamsWithHTTPClient(client *http.Client) *GetBrokersForTableV2Params {
	var ()
	return &GetBrokersForTableV2Params{
		HTTPClient: client,
	}
}

/*GetBrokersForTableV2Params contains all the parameters to send to the API endpoint
for the get brokers for table v2 operation typically these are written to a http.Request
*/
type GetBrokersForTableV2Params struct {

	/*State
	  ONLINE|OFFLINE

	*/
	State *string
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

// WithTimeout adds the timeout to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) WithTimeout(timeout time.Duration) *GetBrokersForTableV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) WithContext(ctx context.Context) *GetBrokersForTableV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) WithHTTPClient(client *http.Client) *GetBrokersForTableV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithState adds the state to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) WithState(state *string) *GetBrokersForTableV2Params {
	o.SetState(state)
	return o
}

// SetState adds the state to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) SetState(state *string) {
	o.State = state
}

// WithTableName adds the tableName to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) WithTableName(tableName string) *GetBrokersForTableV2Params {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) WithType(typeVar *string) *GetBrokersForTableV2Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get brokers for table v2 params
func (o *GetBrokersForTableV2Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetBrokersForTableV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

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
