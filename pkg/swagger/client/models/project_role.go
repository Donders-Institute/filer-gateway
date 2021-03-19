// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectRole JSON object for a member role of a project.
//
// swagger:model projectRole
type ProjectRole struct {

	// project identifier
	// Required: true
	ProjectID *string `json:"projectID"`

	// role of the member.
	// Required: true
	// Enum: [manager contributor viewer traverse]
	Role *string `json:"role"`
}

// Validate validates this project role
func (m *ProjectRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectRole) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

var projectRoleTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manager","contributor","viewer","traverse"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectRoleTypeRolePropEnum = append(projectRoleTypeRolePropEnum, v)
	}
}

const (

	// ProjectRoleRoleManager captures enum value "manager"
	ProjectRoleRoleManager string = "manager"

	// ProjectRoleRoleContributor captures enum value "contributor"
	ProjectRoleRoleContributor string = "contributor"

	// ProjectRoleRoleViewer captures enum value "viewer"
	ProjectRoleRoleViewer string = "viewer"

	// ProjectRoleRoleTraverse captures enum value "traverse"
	ProjectRoleRoleTraverse string = "traverse"
)

// prop value enum
func (m *ProjectRole) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectRoleTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectRole) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project role based on context it is used
func (m *ProjectRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectRole) UnmarshalBinary(b []byte) error {
	var res ProjectRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
