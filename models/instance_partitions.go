// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstancePartitions instance partitions
//
// swagger:model InstancePartitions
type InstancePartitions struct {

	// instance partitions name
	// Read Only: true
	InstancePartitionsName string `json:"instancePartitionsName,omitempty"`

	// partition to instances map
	// Read Only: true
	PartitionToInstancesMap map[string][]string `json:"partitionToInstancesMap,omitempty"`
}

// Validate validates this instance partitions
func (m *InstancePartitions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstancePartitions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstancePartitions) UnmarshalBinary(b []byte) error {
	var res InstancePartitions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
