// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateTenantReader is a Reader for the CreateTenant structure.
type CreateTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCreateTenantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTenantOK creates a CreateTenantOK with default headers values
func NewCreateTenantOK() *CreateTenantOK {
	return &CreateTenantOK{}
}

/*CreateTenantOK handles this case with default header values.

Success
*/
type CreateTenantOK struct {
}

func (o *CreateTenantOK) Error() string {
	return fmt.Sprintf("[POST /tenants][%d] createTenantOK ", 200)
}

func (o *CreateTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTenantInternalServerError creates a CreateTenantInternalServerError with default headers values
func NewCreateTenantInternalServerError() *CreateTenantInternalServerError {
	return &CreateTenantInternalServerError{}
}

/*CreateTenantInternalServerError handles this case with default header values.

Error creating tenant
*/
type CreateTenantInternalServerError struct {
}

func (o *CreateTenantInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tenants][%d] createTenantInternalServerError ", 500)
}

func (o *CreateTenantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
