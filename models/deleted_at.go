// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

// DeletedAt deleted at
//
// swagger:model DeletedAt
type DeletedAt struct {
	gorm.DeletedAt
}

func (m DeletedAt) Validate(formats strfmt.Registry) error {
	var f interface{} = m.DeletedAt
	if v, ok := f.(runtime.Validatable); ok {
		return v.Validate(formats)
	}
	return nil
}
