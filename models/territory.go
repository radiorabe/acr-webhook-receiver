// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Territory territory
//
// swagger:model Territory
type Territory struct {

	// release date
	ReleaseDate string `json:"release_date,omitempty"`

	// territories
	Territories []string `json:"territories"`
}

// Validate validates this territory
func (m *Territory) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this territory based on context it is used
func (m *Territory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Territory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Territory) UnmarshalBinary(b []byte) error {
	var res Territory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
