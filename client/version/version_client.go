// Code generated by go-swagger; DO NOT EDIT.

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new version API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for version API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetVersionNumber(params *GetVersionNumberParams) (*GetVersionNumberOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetVersionNumber gets version number of pinot components
*/
func (a *Client) GetVersionNumber(params *GetVersionNumberParams) (*GetVersionNumberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionNumberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVersionNumber",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVersionNumberReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionNumberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVersionNumber: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
