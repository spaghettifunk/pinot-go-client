// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spaghettifunk/pinot-go-client/models"
)

// UpdateTableConfigReader is a Reader for the UpdateTableConfig structure.
type UpdateTableConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTableConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTableConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTableConfigOK creates a UpdateTableConfigOK with default headers values
func NewUpdateTableConfigOK() *UpdateTableConfigOK {
	return &UpdateTableConfigOK{}
}

/*UpdateTableConfigOK handles this case with default header values.

successful operation
*/
type UpdateTableConfigOK struct {
	Payload *models.SuccessResponse
}

func (o *UpdateTableConfigOK) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}][%d] updateTableConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateTableConfigOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *UpdateTableConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
