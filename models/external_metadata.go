// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExternalMetadata external metadata
//
// swagger:model ExternalMetadata
type ExternalMetadata struct {

	// deezer
	Deezer interface{} `json:"deezer,omitempty"`

	// isrc
	Isrc interface{} `json:"isrc,omitempty"`

	// itunes
	Itunes interface{} `json:"itunes,omitempty"`

	// lyricfind
	Lyricfind interface{} `json:"lyricfind,omitempty"`

	// musicstory
	Musicstory interface{} `json:"musicstory,omitempty"`

	// spotify
	Spotify interface{} `json:"spotify,omitempty"`

	// upc
	Upc interface{} `json:"upc,omitempty"`

	// youtube
	Youtube interface{} `json:"youtube,omitempty"`
}

// Validate validates this external metadata
func (m *ExternalMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this external metadata based on context it is used
func (m *ExternalMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalMetadata) UnmarshalBinary(b []byte) error {
	var res ExternalMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
