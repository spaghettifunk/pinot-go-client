// Code generated by go-swagger; DO NOT EDIT.

package zookeeper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOK creates a DeleteOK with default headers values
func NewDeleteOK() *DeleteOK {
	return &DeleteOK{}
}

/*DeleteOK handles this case with default header values.

Success
*/
type DeleteOK struct {
}

func (o *DeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /zk/delete][%d] deleteOK ", 200)
}

func (o *DeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNoContent creates a DeleteNoContent with default headers values
func NewDeleteNoContent() *DeleteNoContent {
	return &DeleteNoContent{}
}

/*DeleteNoContent handles this case with default header values.

No Content
*/
type DeleteNoContent struct {
}

func (o *DeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /zk/delete][%d] deleteNoContent ", 204)
}

func (o *DeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNotFound creates a DeleteNotFound with default headers values
func NewDeleteNotFound() *DeleteNotFound {
	return &DeleteNotFound{}
}

/*DeleteNotFound handles this case with default header values.

ZK Path not found
*/
type DeleteNotFound struct {
}

func (o *DeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /zk/delete][%d] deleteNotFound ", 404)
}

func (o *DeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInternalServerError creates a DeleteInternalServerError with default headers values
func NewDeleteInternalServerError() *DeleteInternalServerError {
	return &DeleteInternalServerError{}
}

/*DeleteInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteInternalServerError struct {
}

func (o *DeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /zk/delete][%d] deleteInternalServerError ", 500)
}

func (o *DeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
