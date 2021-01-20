// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateTenantReader is a Reader for the UpdateTenant structure.
type UpdateTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateTenantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTenantOK creates a UpdateTenantOK with default headers values
func NewUpdateTenantOK() *UpdateTenantOK {
	return &UpdateTenantOK{}
}

/*UpdateTenantOK handles this case with default header values.

Success
*/
type UpdateTenantOK struct {
}

func (o *UpdateTenantOK) Error() string {
	return fmt.Sprintf("[PUT /tenants][%d] updateTenantOK ", 200)
}

func (o *UpdateTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTenantInternalServerError creates a UpdateTenantInternalServerError with default headers values
func NewUpdateTenantInternalServerError() *UpdateTenantInternalServerError {
	return &UpdateTenantInternalServerError{}
}

/*UpdateTenantInternalServerError handles this case with default header values.

Failed to update the tenant
*/
type UpdateTenantInternalServerError struct {
}

func (o *UpdateTenantInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tenants][%d] updateTenantInternalServerError ", 500)
}

func (o *UpdateTenantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}