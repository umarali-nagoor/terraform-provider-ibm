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

// GetNetworkAclsNetworkACLIDRulesOKBodyFirst A reference to the first page of resources
// swagger:model getNetworkAclsNetworkAclIdRulesOKBodyFirst
type GetNetworkAclsNetworkACLIDRulesOKBodyFirst struct {

	// The URL for the first page of resources
	// Required: true
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href *string `json:"href"`
}

// Validate validates this get network acls network Acl Id rules o k body first
func (m *GetNetworkAclsNetworkACLIDRulesOKBodyFirst) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetNetworkAclsNetworkACLIDRulesOKBodyFirst) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	if err := validate.Pattern("href", "body", string(*m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetNetworkAclsNetworkACLIDRulesOKBodyFirst) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNetworkAclsNetworkACLIDRulesOKBodyFirst) UnmarshalBinary(b []byte) error {
	var res GetNetworkAclsNetworkACLIDRulesOKBodyFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
