// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateTableMetadataReader is a Reader for the UpdateTableMetadata structure.
type UpdateTableMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTableMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTableMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateTableMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTableMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTableMetadataOK creates a UpdateTableMetadataOK with default headers values
func NewUpdateTableMetadataOK() *UpdateTableMetadataOK {
	return &UpdateTableMetadataOK{}
}

/*UpdateTableMetadataOK handles this case with default header values.

Success
*/
type UpdateTableMetadataOK struct {
}

func (o *UpdateTableMetadataOK) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/metadataConfigs][%d] updateTableMetadataOK ", 200)
}

func (o *UpdateTableMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTableMetadataNotFound creates a UpdateTableMetadataNotFound with default headers values
func NewUpdateTableMetadataNotFound() *UpdateTableMetadataNotFound {
	return &UpdateTableMetadataNotFound{}
}

/*UpdateTableMetadataNotFound handles this case with default header values.

Table not found
*/
type UpdateTableMetadataNotFound struct {
}

func (o *UpdateTableMetadataNotFound) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/metadataConfigs][%d] updateTableMetadataNotFound ", 404)
}

func (o *UpdateTableMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTableMetadataInternalServerError creates a UpdateTableMetadataInternalServerError with default headers values
func NewUpdateTableMetadataInternalServerError() *UpdateTableMetadataInternalServerError {
	return &UpdateTableMetadataInternalServerError{}
}

/*UpdateTableMetadataInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateTableMetadataInternalServerError struct {
}

func (o *UpdateTableMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/metadataConfigs][%d] updateTableMetadataInternalServerError ", 500)
}

func (o *UpdateTableMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}