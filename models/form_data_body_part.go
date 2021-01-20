// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FormDataBodyPart form data body part
//
// swagger:model FormDataBodyPart
type FormDataBodyPart struct {

	// content disposition
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`

	// entity
	Entity interface{} `json:"entity,omitempty"`

	// form data content disposition
	FormDataContentDisposition *FormDataContentDisposition `json:"formDataContentDisposition,omitempty"`

	// headers
	Headers map[string][]string `json:"headers,omitempty"`

	// media type
	MediaType *MediaType `json:"mediaType,omitempty"`

	// message body workers
	MessageBodyWorkers MessageBodyWorkers `json:"messageBodyWorkers,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameterized headers
	ParameterizedHeaders map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`

	// parent
	Parent *MultiPart `json:"parent,omitempty"`

	// providers
	Providers Providers `json:"providers,omitempty"`

	// simple
	Simple bool `json:"simple,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this form data body part
func (m *FormDataBodyPart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormDataContentDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterizedHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FormDataBodyPart) validateContentDisposition(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentDisposition) { // not required
		return nil
	}

	if m.ContentDisposition != nil {
		if err := m.ContentDisposition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentDisposition")
			}
			return err
		}
	}

	return nil
}

func (m *FormDataBodyPart) validateFormDataContentDisposition(formats strfmt.Registry) error {

	if swag.IsZero(m.FormDataContentDisposition) { // not required
		return nil
	}

	if m.FormDataContentDisposition != nil {
		if err := m.FormDataContentDisposition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formDataContentDisposition")
			}
			return err
		}
	}

	return nil
}

func (m *FormDataBodyPart) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	if m.MediaType != nil {
		if err := m.MediaType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

func (m *FormDataBodyPart) validateParameterizedHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.ParameterizedHeaders) { // not required
		return nil
	}

	for k := range m.ParameterizedHeaders {

		if err := validate.Required("parameterizedHeaders"+"."+k, "body", m.ParameterizedHeaders[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.ParameterizedHeaders[k]); i++ {

			if err := m.ParameterizedHeaders[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterizedHeaders" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *FormDataBodyPart) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FormDataBodyPart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormDataBodyPart) UnmarshalBinary(b []byte) error {
	var res FormDataBodyPart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
