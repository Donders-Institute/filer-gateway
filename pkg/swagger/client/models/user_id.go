// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// UserID user identifier.
//
// swagger:model userID
type UserID string

// Validate validates this user ID
func (m UserID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user ID based on context it is used
func (m UserID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
