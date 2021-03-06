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

// NewGetBrokersForTenantParams creates a new GetBrokersForTenantParams object
// with the default values initialized.
func NewGetBrokersForTenantParams() *GetBrokersForTenantParams {
	var ()
	return &GetBrokersForTenantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBrokersForTenantParamsWithTimeout creates a new GetBrokersForTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBrokersForTenantParamsWithTimeout(timeout time.Duration) *GetBrokersForTenantParams {
	var ()
	return &GetBrokersForTenantParams{

		timeout: timeout,
	}
}

// NewGetBrokersForTenantParamsWithContext creates a new GetBrokersForTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBrokersForTenantParamsWithContext(ctx context.Context) *GetBrokersForTenantParams {
	var ()
	return &GetBrokersForTenantParams{

		Context: ctx,
	}
}

// NewGetBrokersForTenantParamsWithHTTPClient creates a new GetBrokersForTenantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBrokersForTenantParamsWithHTTPClient(client *http.Client) *GetBrokersForTenantParams {
	var ()
	return &GetBrokersForTenantParams{
		HTTPClient: client,
	}
}

/*GetBrokersForTenantParams contains all the parameters to send to the API endpoint
for the get brokers for tenant operation typically these are written to a http.Request
*/
type GetBrokersForTenantParams struct {

	/*State
	  ONLINE|OFFLINE

	*/
	State *string
	/*TenantName
	  Name of the tenant

	*/
	TenantName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get brokers for tenant params
func (o *GetBrokersForTenantParams) WithTimeout(timeout time.Duration) *GetBrokersForTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get brokers for tenant params
func (o *GetBrokersForTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get brokers for tenant params
func (o *GetBrokersForTenantParams) WithContext(ctx context.Context) *GetBrokersForTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get brokers for tenant params
func (o *GetBrokersForTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get brokers for tenant params
func (o *GetBrokersForTenantParams) WithHTTPClient(client *http.Client) *GetBrokersForTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get brokers for tenant params
func (o *GetBrokersForTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithState adds the state to the get brokers for tenant params
func (o *GetBrokersForTenantParams) WithState(state *string) *GetBrokersForTenantParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get brokers for tenant params
func (o *GetBrokersForTenantParams) SetState(state *string) {
	o.State = state
}

// WithTenantName adds the tenantName to the get brokers for tenant params
func (o *GetBrokersForTenantParams) WithTenantName(tenantName string) *GetBrokersForTenantParams {
	o.SetTenantName(tenantName)
	return o
}

// SetTenantName adds the tenantName to the get brokers for tenant params
func (o *GetBrokersForTenantParams) SetTenantName(tenantName string) {
	o.TenantName = tenantName
}

// WriteToRequest writes these params to a swagger request
func (o *GetBrokersForTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tenantName
	if err := r.SetPathParam("tenantName", o.TenantName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
