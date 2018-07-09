// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateModuleParamsBody create module params body
// swagger:model createModuleParamsBody
type CreateModuleParamsBody struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// type
	// Required: true
	Type *string `json:"type"`

	// ID of the VCS source for the module
	// Required: true
	VcsID *string `json:"vcs-id"`
}

// Validate validates this create module params body
func (m *CreateModuleParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateModuleParamsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateModuleParamsBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CreateModuleParamsBody) validateVcsID(formats strfmt.Registry) error {

	if err := validate.Required("vcs-id", "body", m.VcsID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateModuleParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateModuleParamsBody) UnmarshalBinary(b []byte) error {
	var res CreateModuleParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
