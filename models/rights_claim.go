// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RightsClaim rights claim
//
// swagger:model RightsClaim
type RightsClaim struct {

	// distributor
	Distributor *Distributor `json:"distributor,omitempty"`

	// rights claim policy
	RightsClaimPolicy string `json:"rights_claim_policy,omitempty"`

	// rights owners
	RightsOwners []*RightsOwner `json:"rights_owners"`

	// start date
	StartDate string `json:"start_date,omitempty"`

	// territories
	Territories []string `json:"territories"`
}

// Validate validates this rights claim
func (m *RightsClaim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistributor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRightsOwners(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RightsClaim) validateDistributor(formats strfmt.Registry) error {
	if swag.IsZero(m.Distributor) { // not required
		return nil
	}

	if m.Distributor != nil {
		if err := m.Distributor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distributor")
			}
			return err
		}
	}

	return nil
}

func (m *RightsClaim) validateRightsOwners(formats strfmt.Registry) error {
	if swag.IsZero(m.RightsOwners) { // not required
		return nil
	}

	for i := 0; i < len(m.RightsOwners); i++ {
		if swag.IsZero(m.RightsOwners[i]) { // not required
			continue
		}

		if m.RightsOwners[i] != nil {
			if err := m.RightsOwners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rights_owners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rights claim based on the context it is used
func (m *RightsClaim) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDistributor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRightsOwners(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RightsClaim) contextValidateDistributor(ctx context.Context, formats strfmt.Registry) error {

	if m.Distributor != nil {
		if err := m.Distributor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distributor")
			}
			return err
		}
	}

	return nil
}

func (m *RightsClaim) contextValidateRightsOwners(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RightsOwners); i++ {

		if m.RightsOwners[i] != nil {
			if err := m.RightsOwners[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rights_owners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RightsClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RightsClaim) UnmarshalBinary(b []byte) error {
	var res RightsClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
