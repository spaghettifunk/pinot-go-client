// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceInfo instance info
//
// swagger:model InstanceInfo
type InstanceInfo struct {

	// host
	Host string `json:"host,omitempty"`

	// instance name
	InstanceName string `json:"instanceName,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`
}

// Validate validates this instance info
func (m *InstanceInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceInfo) UnmarshalBinary(b []byte) error {
	var res InstanceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
