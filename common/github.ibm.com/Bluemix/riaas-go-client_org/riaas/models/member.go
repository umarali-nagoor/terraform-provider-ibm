// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Member member
// swagger:model Member
type Member struct {

	// The date and time that this member was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Health of the server member in the pool.
	// Enum: [ok faulted unknown]
	Health string `json:"health,omitempty"`

	// The member's canonical URL.
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The member's unique identifier.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The port number of the application running in the server member.
	// Maximum: 65535
	// Minimum: 1
	Port int64 `json:"port,omitempty"`

	// The provisioning status of this member
	// Enum: [active create_pending update_pending delete_pending maintenance_pending]
	ProvisioningStatus string `json:"provisioning_status,omitempty"`

	// target
	Target *MemberTarget `json:"target,omitempty"`

	// Weight of the server member. This takes effect only when the load balancing algorithm of its belonging pool is `weighted_round_robin`.
	// Maximum: 100
	// Minimum: 1
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this member
func (m *Member) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisioningStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Member) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var memberTypeHealthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","faulted","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberTypeHealthPropEnum = append(memberTypeHealthPropEnum, v)
	}
}

const (

	// MemberHealthOk captures enum value "ok"
	MemberHealthOk string = "ok"

	// MemberHealthFaulted captures enum value "faulted"
	MemberHealthFaulted string = "faulted"

	// MemberHealthUnknown captures enum value "unknown"
	MemberHealthUnknown string = "unknown"
)

// prop value enum
func (m *Member) validateHealthEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, memberTypeHealthPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Member) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	// value enum
	if err := m.validateHealthEnum("health", "body", m.Health); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Member) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var memberTypeProvisioningStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","create_pending","update_pending","delete_pending","maintenance_pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberTypeProvisioningStatusPropEnum = append(memberTypeProvisioningStatusPropEnum, v)
	}
}

const (

	// MemberProvisioningStatusActive captures enum value "active"
	MemberProvisioningStatusActive string = "active"

	// MemberProvisioningStatusCreatePending captures enum value "create_pending"
	MemberProvisioningStatusCreatePending string = "create_pending"

	// MemberProvisioningStatusUpdatePending captures enum value "update_pending"
	MemberProvisioningStatusUpdatePending string = "update_pending"

	// MemberProvisioningStatusDeletePending captures enum value "delete_pending"
	MemberProvisioningStatusDeletePending string = "delete_pending"

	// MemberProvisioningStatusMaintenancePending captures enum value "maintenance_pending"
	MemberProvisioningStatusMaintenancePending string = "maintenance_pending"
)

// prop value enum
func (m *Member) validateProvisioningStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, memberTypeProvisioningStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Member) validateProvisioningStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ProvisioningStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateProvisioningStatusEnum("provisioning_status", "body", m.ProvisioningStatus); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *Member) validateWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", int64(m.Weight), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", int64(m.Weight), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Member) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Member) UnmarshalBinary(b []byte) error {
	var res Member
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
