// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ParameterizedHeader parameterized header
//
// swagger:model ParameterizedHeader
type ParameterizedHeader struct {

	// parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this parameterized header
func (m *ParameterizedHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterizedHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterizedHeader) UnmarshalBinary(b []byte) error {
	var res ParameterizedHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
