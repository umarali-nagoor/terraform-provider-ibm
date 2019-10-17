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

// PostFloatingIpsParamsBodyResourceGroup idreference
// swagger:model postFloatingIpsParamsBodyResourceGroup
type PostFloatingIpsParamsBodyResourceGroup struct {

	// The unique identifier for this resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post floating ips params body resource group
func (m *PostFloatingIpsParamsBodyResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostFloatingIpsParamsBodyResourceGroup) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostFloatingIpsParamsBodyResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostFloatingIpsParamsBodyResourceGroup) UnmarshalBinary(b []byte) error {
	var res PostFloatingIpsParamsBodyResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
