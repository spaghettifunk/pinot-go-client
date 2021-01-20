// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DateTimeFieldSpec date time field spec
//
// swagger:model DateTimeFieldSpec
type DateTimeFieldSpec struct {

	// data type
	// Enum: [INT LONG FLOAT DOUBLE BOOLEAN STRING BYTES STRUCT MAP LIST]
	DataType string `json:"dataType,omitempty"`

	// default null value
	DefaultNullValue interface{} `json:"defaultNullValue,omitempty"`

	// default null value string
	DefaultNullValueString string `json:"defaultNullValueString,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// granularity
	Granularity string `json:"granularity,omitempty"`

	// max length
	MaxLength int32 `json:"maxLength,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// single value field
	SingleValueField bool `json:"singleValueField,omitempty"`

	// transform function
	TransformFunction string `json:"transformFunction,omitempty"`

	// virtual column provider
	VirtualColumnProvider string `json:"virtualColumnProvider,omitempty"`
}

// Validate validates this date time field spec
func (m *DateTimeFieldSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dateTimeFieldSpecTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INT","LONG","FLOAT","DOUBLE","BOOLEAN","STRING","BYTES","STRUCT","MAP","LIST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dateTimeFieldSpecTypeDataTypePropEnum = append(dateTimeFieldSpecTypeDataTypePropEnum, v)
	}
}

const (

	// DateTimeFieldSpecDataTypeINT captures enum value "INT"
	DateTimeFieldSpecDataTypeINT string = "INT"

	// DateTimeFieldSpecDataTypeLONG captures enum value "LONG"
	DateTimeFieldSpecDataTypeLONG string = "LONG"

	// DateTimeFieldSpecDataTypeFLOAT captures enum value "FLOAT"
	DateTimeFieldSpecDataTypeFLOAT string = "FLOAT"

	// DateTimeFieldSpecDataTypeDOUBLE captures enum value "DOUBLE"
	DateTimeFieldSpecDataTypeDOUBLE string = "DOUBLE"

	// DateTimeFieldSpecDataTypeBOOLEAN captures enum value "BOOLEAN"
	DateTimeFieldSpecDataTypeBOOLEAN string = "BOOLEAN"

	// DateTimeFieldSpecDataTypeSTRING captures enum value "STRING"
	DateTimeFieldSpecDataTypeSTRING string = "STRING"

	// DateTimeFieldSpecDataTypeBYTES captures enum value "BYTES"
	DateTimeFieldSpecDataTypeBYTES string = "BYTES"

	// DateTimeFieldSpecDataTypeSTRUCT captures enum value "STRUCT"
	DateTimeFieldSpecDataTypeSTRUCT string = "STRUCT"

	// DateTimeFieldSpecDataTypeMAP captures enum value "MAP"
	DateTimeFieldSpecDataTypeMAP string = "MAP"

	// DateTimeFieldSpecDataTypeLIST captures enum value "LIST"
	DateTimeFieldSpecDataTypeLIST string = "LIST"
)

// prop value enum
func (m *DateTimeFieldSpec) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dateTimeFieldSpecTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DateTimeFieldSpec) validateDataType(formats strfmt.Registry) error {

	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("dataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DateTimeFieldSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DateTimeFieldSpec) UnmarshalBinary(b []byte) error {
	var res DateTimeFieldSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
