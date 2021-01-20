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

// NewGetTenantsToBrokersMappingV2Params creates a new GetTenantsToBrokersMappingV2Params object
// with the default values initialized.
func NewGetTenantsToBrokersMappingV2Params() *GetTenantsToBrokersMappingV2Params {
	var ()
	return &GetTenantsToBrokersMappingV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantsToBrokersMappingV2ParamsWithTimeout creates a new GetTenantsToBrokersMappingV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantsToBrokersMappingV2ParamsWithTimeout(timeout time.Duration) *GetTenantsToBrokersMappingV2Params {
	var ()
	return &GetTenantsToBrokersMappingV2Params{

		timeout: timeout,
	}
}

// NewGetTenantsToBrokersMappingV2ParamsWithContext creates a new GetTenantsToBrokersMappingV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantsToBrokersMappingV2ParamsWithContext(ctx context.Context) *GetTenantsToBrokersMappingV2Params {
	var ()
	return &GetTenantsToBrokersMappingV2Params{

		Context: ctx,
	}
}

// NewGetTenantsToBrokersMappingV2ParamsWithHTTPClient creates a new GetTenantsToBrokersMappingV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantsToBrokersMappingV2ParamsWithHTTPClient(client *http.Client) *GetTenantsToBrokersMappingV2Params {
	var ()
	return &GetTenantsToBrokersMappingV2Params{
		HTTPClient: client,
	}
}

/*GetTenantsToBrokersMappingV2Params contains all the parameters to send to the API endpoint
for the get tenants to brokers mapping v2 operation typically these are written to a http.Request
*/
type GetTenantsToBrokersMappingV2Params struct {

	/*State
	  ONLINE|OFFLINE

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) WithTimeout(timeout time.Duration) *GetTenantsToBrokersMappingV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) WithContext(ctx context.Context) *GetTenantsToBrokersMappingV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) WithHTTPClient(client *http.Client) *GetTenantsToBrokersMappingV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithState adds the state to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) WithState(state *string) *GetTenantsToBrokersMappingV2Params {
	o.SetState(state)
	return o
}

// SetState adds the state to the get tenants to brokers mapping v2 params
func (o *GetTenantsToBrokersMappingV2Params) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantsToBrokersMappingV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
