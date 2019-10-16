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

// IpsecPolicyTemplateResourceGroup ResourceGroupIdentity
// swagger:model ipsecPolicyTemplateResourceGroup
type IpsecPolicyTemplateResourceGroup struct {

	// The unique identifier for this resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this ipsec policy template resource group
func (m *IpsecPolicyTemplateResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpsecPolicyTemplateResourceGroup) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpsecPolicyTemplateResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpsecPolicyTemplateResourceGroup) UnmarshalBinary(b []byte) error {
	var res IpsecPolicyTemplateResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
