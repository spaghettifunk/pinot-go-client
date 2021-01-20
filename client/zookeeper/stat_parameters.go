// Code generated by go-swagger; DO NOT EDIT.

package zookeeper

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

// NewStatParams creates a new StatParams object
// with the default values initialized.
func NewStatParams() *StatParams {
	var (
		pathDefault = string("/")
	)
	return &StatParams{
		Path: pathDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewStatParamsWithTimeout creates a new StatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatParamsWithTimeout(timeout time.Duration) *StatParams {
	var (
		pathDefault = string("/")
	)
	return &StatParams{
		Path: pathDefault,

		timeout: timeout,
	}
}

// NewStatParamsWithContext creates a new StatParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatParamsWithContext(ctx context.Context) *StatParams {
	var (
		pathDefault = string("/")
	)
	return &StatParams{
		Path: pathDefault,

		Context: ctx,
	}
}

// NewStatParamsWithHTTPClient creates a new StatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatParamsWithHTTPClient(client *http.Client) *StatParams {
	var (
		pathDefault = string("/")
	)
	return &StatParams{
		Path:       pathDefault,
		HTTPClient: client,
	}
}

/*StatParams contains all the parameters to send to the API endpoint
for the stat operation typically these are written to a http.Request
*/
type StatParams struct {

	/*Path
	  Zookeeper Path, must start with /

	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stat params
func (o *StatParams) WithTimeout(timeout time.Duration) *StatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stat params
func (o *StatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stat params
func (o *StatParams) WithContext(ctx context.Context) *StatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stat params
func (o *StatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stat params
func (o *StatParams) WithHTTPClient(client *http.Client) *StatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stat params
func (o *StatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the stat params
func (o *StatParams) WithPath(path string) *StatParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the stat params
func (o *StatParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *StatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {
		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
