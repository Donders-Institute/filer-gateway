// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Member JSON object for a project member.
// swagger:model member
type Member struct {

	// role of the member. Use the value "none" in request data to remove user from the project membership.
	// Required: true
	// Enum: [manager contributor viewer traverse none]
	Role *string `json:"role"`

	// userid of the member.
	// Required: true
	UserID *string `json:"userID"`
}

// Validate validates this member
func (m *Member) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var memberTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manager","contributor","viewer","traverse","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberTypeRolePropEnum = append(memberTypeRolePropEnum, v)
	}
}

const (

	// MemberRoleManager captures enum value "manager"
	MemberRoleManager string = "manager"

	// MemberRoleContributor captures enum value "contributor"
	MemberRoleContributor string = "contributor"

	// MemberRoleViewer captures enum value "viewer"
	MemberRoleViewer string = "viewer"

	// MemberRoleTraverse captures enum value "traverse"
	MemberRoleTraverse string = "traverse"

	// MemberRoleNone captures enum value "none"
	MemberRoleNone string = "none"
)

// prop value enum
func (m *Member) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, memberTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Member) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userID", "body", m.UserID); err != nil {
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
