// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CheckHealthLegacyReader is a Reader for the CheckHealthLegacy structure.
type CheckHealthLegacyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckHealthLegacyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckHealthLegacyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckHealthLegacyOK creates a CheckHealthLegacyOK with default headers values
func NewCheckHealthLegacyOK() *CheckHealthLegacyOK {
	return &CheckHealthLegacyOK{}
}

/*CheckHealthLegacyOK handles this case with default header values.

Good
*/
type CheckHealthLegacyOK struct {
}

func (o *CheckHealthLegacyOK) Error() string {
	return fmt.Sprintf("[GET /pinot-controller/admin][%d] checkHealthLegacyOK ", 200)
}

func (o *CheckHealthLegacyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
