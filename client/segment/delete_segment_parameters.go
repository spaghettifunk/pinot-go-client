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

// NewDeleteSegmentParams creates a new DeleteSegmentParams object
// with the default values initialized.
func NewDeleteSegmentParams() *DeleteSegmentParams {
	var ()
	return &DeleteSegmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSegmentParamsWithTimeout creates a new DeleteSegmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSegmentParamsWithTimeout(timeout time.Duration) *DeleteSegmentParams {
	var ()
	return &DeleteSegmentParams{

		timeout: timeout,
	}
}

// NewDeleteSegmentParamsWithContext creates a new DeleteSegmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSegmentParamsWithContext(ctx context.Context) *DeleteSegmentParams {
	var ()
	return &DeleteSegmentParams{

		Context: ctx,
	}
}

// NewDeleteSegmentParamsWithHTTPClient creates a new DeleteSegmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSegmentParamsWithHTTPClient(client *http.Client) *DeleteSegmentParams {
	var ()
	return &DeleteSegmentParams{
		HTTPClient: client,
	}
}

/*DeleteSegmentParams contains all the parameters to send to the API endpoint
for the delete segment operation typically these are written to a http.Request
*/
type DeleteSegmentParams struct {

	/*SegmentName
	  Name of the segment

	*/
	SegmentName string
	/*TableName
	  Name of the table

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete segment params
func (o *DeleteSegmentParams) WithTimeout(timeout time.Duration) *DeleteSegmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete segment params
func (o *DeleteSegmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete segment params
func (o *DeleteSegmentParams) WithContext(ctx context.Context) *DeleteSegmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete segment params
func (o *DeleteSegmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete segment params
func (o *DeleteSegmentParams) WithHTTPClient(client *http.Client) *DeleteSegmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete segment params
func (o *DeleteSegmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSegmentName adds the segmentName to the delete segment params
func (o *DeleteSegmentParams) WithSegmentName(segmentName string) *DeleteSegmentParams {
	o.SetSegmentName(segmentName)
	return o
}

// SetSegmentName adds the segmentName to the delete segment params
func (o *DeleteSegmentParams) SetSegmentName(segmentName string) {
	o.SegmentName = segmentName
}

// WithTableName adds the tableName to the delete segment params
func (o *DeleteSegmentParams) WithTableName(tableName string) *DeleteSegmentParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the delete segment params
func (o *DeleteSegmentParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSegmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param segmentName
	if err := r.SetPathParam("segmentName", o.SegmentName); err != nil {
		return err
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
