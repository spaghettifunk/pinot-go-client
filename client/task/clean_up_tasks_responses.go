// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spaghettifunk/pinot-go-client/models"
)

// CleanUpTasksReader is a Reader for the CleanUpTasks structure.
type CleanUpTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CleanUpTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCleanUpTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCleanUpTasksOK creates a CleanUpTasksOK with default headers values
func NewCleanUpTasksOK() *CleanUpTasksOK {
	return &CleanUpTasksOK{}
}

/*CleanUpTasksOK handles this case with default header values.

successful operation
*/
type CleanUpTasksOK struct {
	Payload *models.SuccessResponse
}

func (o *CleanUpTasksOK) Error() string {
	return fmt.Sprintf("[PUT /tasks/{taskType}/cleanup][%d] cleanUpTasksOK  %+v", 200, o.Payload)
}

func (o *CleanUpTasksOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *CleanUpTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}