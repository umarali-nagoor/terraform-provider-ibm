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

// PostSecurityGroupsParamsBodyRulesItemsRemote SecurityGroupIdentityByName
//
// Uniquely identifies a security group using any one of ID, CRN, or name.
// swagger:model postSecurityGroupsParamsBodyRulesItemsRemote
type PostSecurityGroupsParamsBodyRulesItemsRemote struct {

	// A single IPv4 or IPv6 address.
	Address string `json:"address,omitempty"`

	// A range of IPv4 or IPv6 addresses, in CIDR format.
	CidrBlock string `json:"cidr_block,omitempty"`

	// The security group's CRN
	Crn string `json:"crn,omitempty"`

	// The security group's unique identifier.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The security group's user-defined name.
	// Required: true
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name *string `json:"name"`
}

// Validate validates this post security groups params body rules items remote
func (m *PostSecurityGroupsParamsBodyRulesItemsRemote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *PostSecurityGroupsParamsBodyRulesItemsRemote) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostSecurityGroupsParamsBodyRulesItemsRemote) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSecurityGroupsParamsBodyRulesItemsRemote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSecurityGroupsParamsBodyRulesItemsRemote) UnmarshalBinary(b []byte) error {
	var res PostSecurityGroupsParamsBodyRulesItemsRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
