// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	hal "github.com/getlunaform/lunaform/models/hal"
)

// ResourceTfProviderConfiguration A Terraform provider configuration
// swagger:model resource-tf-provider-configuration
type ResourceTfProviderConfiguration struct {

	// embedded
	Embedded *ResourceTfProviderConfigurationEmbedded `json:"_embedded,omitempty"`

	// links
	Links *hal.HalRscLinks `json:"_links,omitempty"`

	// configuration
	// Required: true
	Configuration interface{} `json:"configuration"`

	// id
	ID string `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this resource tf provider configuration
func (m *ResourceTfProviderConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceTfProviderConfiguration) validateEmbedded(formats strfmt.Registry) error {

	if swag.IsZero(m.Embedded) { // not required
		return nil
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceTfProviderConfiguration) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceTfProviderConfiguration) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	return nil
}

func (m *ResourceTfProviderConfiguration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceTfProviderConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceTfProviderConfiguration) UnmarshalBinary(b []byte) error {
	var res ResourceTfProviderConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourceTfProviderConfigurationEmbedded resource tf provider configuration embedded
// swagger:model ResourceTfProviderConfigurationEmbedded
type ResourceTfProviderConfigurationEmbedded struct {

	// provider
	Provider *ResourceTfProvider `json:"provider,omitempty"`
}

// Validate validates this resource tf provider configuration embedded
func (m *ResourceTfProviderConfigurationEmbedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceTfProviderConfigurationEmbedded) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if m.Provider != nil {
		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded" + "." + "provider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceTfProviderConfigurationEmbedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceTfProviderConfigurationEmbedded) UnmarshalBinary(b []byte) error {
	var res ResourceTfProviderConfigurationEmbedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
