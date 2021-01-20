// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateInstanceReader is a Reader for the UpdateInstance structure.
type UpdateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInstanceOK creates a UpdateInstanceOK with default headers values
func NewUpdateInstanceOK() *UpdateInstanceOK {
	return &UpdateInstanceOK{}
}

/*UpdateInstanceOK handles this case with default header values.

Success
*/
type UpdateInstanceOK struct {
}

func (o *UpdateInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}][%d] updateInstanceOK ", 200)
}

func (o *UpdateInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstanceInternalServerError creates a UpdateInstanceInternalServerError with default headers values
func NewUpdateInstanceInternalServerError() *UpdateInstanceInternalServerError {
	return &UpdateInstanceInternalServerError{}
}

/*UpdateInstanceInternalServerError handles this case with default header values.

Internal error
*/
type UpdateInstanceInternalServerError struct {
}

func (o *UpdateInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}][%d] updateInstanceInternalServerError ", 500)
}

func (o *UpdateInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
