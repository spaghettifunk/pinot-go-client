// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StartReplaceSegmentsReader is a Reader for the StartReplaceSegments structure.
type StartReplaceSegmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartReplaceSegmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewStartReplaceSegmentsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewStartReplaceSegmentsDefault creates a StartReplaceSegmentsDefault with default headers values
func NewStartReplaceSegmentsDefault(code int) *StartReplaceSegmentsDefault {
	return &StartReplaceSegmentsDefault{
		_statusCode: code,
	}
}

/*StartReplaceSegmentsDefault handles this case with default header values.

successful operation
*/
type StartReplaceSegmentsDefault struct {
	_statusCode int
}

// Code gets the status code for the start replace segments default response
func (o *StartReplaceSegmentsDefault) Code() int {
	return o._statusCode
}

func (o *StartReplaceSegmentsDefault) Error() string {
	return fmt.Sprintf("[POST /segments/{tableName}/startReplaceSegments][%d] startReplaceSegments default ", o._statusCode)
}

func (o *StartReplaceSegmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
