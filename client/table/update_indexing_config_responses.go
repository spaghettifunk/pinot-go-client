// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateIndexingConfigReader is a Reader for the UpdateIndexingConfig structure.
type UpdateIndexingConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIndexingConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIndexingConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateIndexingConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateIndexingConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateIndexingConfigOK creates a UpdateIndexingConfigOK with default headers values
func NewUpdateIndexingConfigOK() *UpdateIndexingConfigOK {
	return &UpdateIndexingConfigOK{}
}

/*UpdateIndexingConfigOK handles this case with default header values.

Success
*/
type UpdateIndexingConfigOK struct {
}

func (o *UpdateIndexingConfigOK) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigOK ", 200)
}

func (o *UpdateIndexingConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateIndexingConfigNotFound creates a UpdateIndexingConfigNotFound with default headers values
func NewUpdateIndexingConfigNotFound() *UpdateIndexingConfigNotFound {
	return &UpdateIndexingConfigNotFound{}
}

/*UpdateIndexingConfigNotFound handles this case with default header values.

Table not found
*/
type UpdateIndexingConfigNotFound struct {
}

func (o *UpdateIndexingConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigNotFound ", 404)
}

func (o *UpdateIndexingConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateIndexingConfigInternalServerError creates a UpdateIndexingConfigInternalServerError with default headers values
func NewUpdateIndexingConfigInternalServerError() *UpdateIndexingConfigInternalServerError {
	return &UpdateIndexingConfigInternalServerError{}
}

/*UpdateIndexingConfigInternalServerError handles this case with default header values.

Server error updating configuration
*/
type UpdateIndexingConfigInternalServerError struct {
}

func (o *UpdateIndexingConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigInternalServerError ", 500)
}

func (o *UpdateIndexingConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
