// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ToggleInstanceStateReader is a Reader for the ToggleInstanceState structure.
type ToggleInstanceStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ToggleInstanceStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewToggleInstanceStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewToggleInstanceStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewToggleInstanceStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewToggleInstanceStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewToggleInstanceStateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewToggleInstanceStateOK creates a ToggleInstanceStateOK with default headers values
func NewToggleInstanceStateOK() *ToggleInstanceStateOK {
	return &ToggleInstanceStateOK{}
}

/*ToggleInstanceStateOK handles this case with default header values.

Success
*/
type ToggleInstanceStateOK struct {
}

func (o *ToggleInstanceStateOK) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceName}/state][%d] toggleInstanceStateOK ", 200)
}

func (o *ToggleInstanceStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewToggleInstanceStateBadRequest creates a ToggleInstanceStateBadRequest with default headers values
func NewToggleInstanceStateBadRequest() *ToggleInstanceStateBadRequest {
	return &ToggleInstanceStateBadRequest{}
}

/*ToggleInstanceStateBadRequest handles this case with default header values.

Bad Request
*/
type ToggleInstanceStateBadRequest struct {
}

func (o *ToggleInstanceStateBadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceName}/state][%d] toggleInstanceStateBadRequest ", 400)
}

func (o *ToggleInstanceStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewToggleInstanceStateNotFound creates a ToggleInstanceStateNotFound with default headers values
func NewToggleInstanceStateNotFound() *ToggleInstanceStateNotFound {
	return &ToggleInstanceStateNotFound{}
}

/*ToggleInstanceStateNotFound handles this case with default header values.

Instance not found
*/
type ToggleInstanceStateNotFound struct {
}

func (o *ToggleInstanceStateNotFound) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceName}/state][%d] toggleInstanceStateNotFound ", 404)
}

func (o *ToggleInstanceStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewToggleInstanceStateConflict creates a ToggleInstanceStateConflict with default headers values
func NewToggleInstanceStateConflict() *ToggleInstanceStateConflict {
	return &ToggleInstanceStateConflict{}
}

/*ToggleInstanceStateConflict handles this case with default header values.

Instance cannot be dropped
*/
type ToggleInstanceStateConflict struct {
}

func (o *ToggleInstanceStateConflict) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceName}/state][%d] toggleInstanceStateConflict ", 409)
}

func (o *ToggleInstanceStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewToggleInstanceStateInternalServerError creates a ToggleInstanceStateInternalServerError with default headers values
func NewToggleInstanceStateInternalServerError() *ToggleInstanceStateInternalServerError {
	return &ToggleInstanceStateInternalServerError{}
}

/*ToggleInstanceStateInternalServerError handles this case with default header values.

Internal error
*/
type ToggleInstanceStateInternalServerError struct {
}

func (o *ToggleInstanceStateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /instances/{instanceName}/state][%d] toggleInstanceStateInternalServerError ", 500)
}

func (o *ToggleInstanceStateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
