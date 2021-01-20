// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartReplaceSegmentsRequest start replace segments request
//
// swagger:model StartReplaceSegmentsRequest
type StartReplaceSegmentsRequest struct {

	// segments from
	// Read Only: true
	SegmentsFrom []string `json:"segmentsFrom"`

	// segments to
	// Read Only: true
	SegmentsTo []string `json:"segmentsTo"`
}

// Validate validates this start replace segments request
func (m *StartReplaceSegmentsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StartReplaceSegmentsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartReplaceSegmentsRequest) UnmarshalBinary(b []byte) error {
	var res StartReplaceSegmentsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}