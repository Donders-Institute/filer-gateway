// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseBody500 JSON object containing server side error.
//
// swagger:model responseBody500
type ResponseBody500 struct {

	// server-side error message.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// server-side exit code.
	ExitCode int64 `json:"exitCode,omitempty"`
}

// Validate validates this response body500
func (m *ResponseBody500) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response body500 based on context it is used
func (m *ResponseBody500) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseBody500) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseBody500) UnmarshalBinary(b []byte) error {
	var res ResponseBody500
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
