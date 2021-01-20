// Code generated by go-swagger; DO NOT EDIT.

package instance

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

	"github.com/spaghettifunk/pinot-go-client/models"
)

// NewAddInstanceParams creates a new AddInstanceParams object
// with the default values initialized.
func NewAddInstanceParams() *AddInstanceParams {
	var ()
	return &AddInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddInstanceParamsWithTimeout creates a new AddInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddInstanceParamsWithTimeout(timeout time.Duration) *AddInstanceParams {
	var ()
	return &AddInstanceParams{

		timeout: timeout,
	}
}

// NewAddInstanceParamsWithContext creates a new AddInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddInstanceParamsWithContext(ctx context.Context) *AddInstanceParams {
	var ()
	return &AddInstanceParams{

		Context: ctx,
	}
}

// NewAddInstanceParamsWithHTTPClient creates a new AddInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddInstanceParamsWithHTTPClient(client *http.Client) *AddInstanceParams {
	var ()
	return &AddInstanceParams{
		HTTPClient: client,
	}
}

/*AddInstanceParams contains all the parameters to send to the API endpoint
for the add instance operation typically these are written to a http.Request
*/
type AddInstanceParams struct {

	/*Body*/
	Body *models.Instance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add instance params
func (o *AddInstanceParams) WithTimeout(timeout time.Duration) *AddInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add instance params
func (o *AddInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add instance params
func (o *AddInstanceParams) WithContext(ctx context.Context) *AddInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add instance params
func (o *AddInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add instance params
func (o *AddInstanceParams) WithHTTPClient(client *http.Client) *AddInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add instance params
func (o *AddInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add instance params
func (o *AddInstanceParams) WithBody(body *models.Instance) *AddInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add instance params
func (o *AddInstanceParams) SetBody(body *models.Instance) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}