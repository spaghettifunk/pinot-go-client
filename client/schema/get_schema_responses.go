// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSchemaReader is a Reader for the GetSchema structure.
type GetSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchemaOK creates a GetSchemaOK with default headers values
func NewGetSchemaOK() *GetSchemaOK {
	return &GetSchemaOK{}
}

/*GetSchemaOK handles this case with default header values.

Success
*/
type GetSchemaOK struct {
}

func (o *GetSchemaOK) Error() string {
	return fmt.Sprintf("[GET /schemas/{schemaName}][%d] getSchemaOK ", 200)
}

func (o *GetSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaNotFound creates a GetSchemaNotFound with default headers values
func NewGetSchemaNotFound() *GetSchemaNotFound {
	return &GetSchemaNotFound{}
}

/*GetSchemaNotFound handles this case with default header values.

Schema not found
*/
type GetSchemaNotFound struct {
}

func (o *GetSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /schemas/{schemaName}][%d] getSchemaNotFound ", 404)
}

func (o *GetSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaInternalServerError creates a GetSchemaInternalServerError with default headers values
func NewGetSchemaInternalServerError() *GetSchemaInternalServerError {
	return &GetSchemaInternalServerError{}
}

/*GetSchemaInternalServerError handles this case with default header values.

Internal error
*/
type GetSchemaInternalServerError struct {
}

func (o *GetSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schemas/{schemaName}][%d] getSchemaInternalServerError ", 500)
}

func (o *GetSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
