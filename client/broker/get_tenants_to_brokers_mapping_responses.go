// Code generated by go-swagger; DO NOT EDIT.

package broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTenantsToBrokersMappingReader is a Reader for the GetTenantsToBrokersMapping structure.
type GetTenantsToBrokersMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsToBrokersMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsToBrokersMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTenantsToBrokersMappingOK creates a GetTenantsToBrokersMappingOK with default headers values
func NewGetTenantsToBrokersMappingOK() *GetTenantsToBrokersMappingOK {
	return &GetTenantsToBrokersMappingOK{}
}

/*GetTenantsToBrokersMappingOK handles this case with default header values.

successful operation
*/
type GetTenantsToBrokersMappingOK struct {
	Payload map[string][]string
}

func (o *GetTenantsToBrokersMappingOK) Error() string {
	return fmt.Sprintf("[GET /brokers/tenants][%d] getTenantsToBrokersMappingOK  %+v", 200, o.Payload)
}

func (o *GetTenantsToBrokersMappingOK) GetPayload() map[string][]string {
	return o.Payload
}

func (o *GetTenantsToBrokersMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
