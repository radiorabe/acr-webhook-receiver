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
	"github.com/go-openapi/validate"
)

// Music music
//
// swagger:model Music
type Music struct {

	// acrid
	// Required: true
	Acrid *string `json:"acrid"`

	// album
	Album *Album `json:"album,omitempty"`

	// artists
	Artists []*Artist `json:"artists"`

	// contributors
	Contributors *Contributors `json:"contributors,omitempty"`

	// db begin time offset ms
	// Required: true
	DbBeginTimeOffsetMs *int64 `json:"db_begin_time_offset_ms"`

	// duration ms
	// Required: true
	DurationMs *int64 `json:"duration_ms"`

	// external ids
	// Required: true
	ExternalIds *ExternalIds `json:"external_ids"`

	// external metadata
	// Required: true
	ExternalMetadata *ExternalMetadata `json:"external_metadata"`

	// genres
	Genres []*Genre `json:"genres"`

	// label
	Label string `json:"label,omitempty"`

	// lyrics
	Lyrics *Lyrics `json:"lyrics,omitempty"`

	// play offset ms
	// Required: true
	PlayOffsetMs *int64 `json:"play_offset_ms"`

	// release by territories
	ReleaseByTerritories []*Territory `json:"release_by_territories"`

	// release date
	ReleaseDate string `json:"release_date,omitempty"`

	// result from
	// Required: true
	ResultFrom *int32 `json:"result_from"`

	// rights claim
	RightsClaim []*RightsClaim `json:"rights_claim"`

	// sample begin time offset ms
	// Required: true
	SampleBeginTimeOffsetMs *int64 `json:"sample_begin_time_offset_ms"`

	// sample end time offset ms
	// Required: true
	SampleEndTimeOffsetMs *int64 `json:"sample_end_time_offset_ms"`

	// score
	// Required: true
	// Maximum: 100
	// Minimum: 0
	Score *int64 `json:"score"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this music
func (m *Music) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcrid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlbum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContributors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDbBeginTimeOffsetMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDurationMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenres(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLyrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayOffsetMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseByTerritories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRightsClaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSampleBeginTimeOffsetMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSampleEndTimeOffsetMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Music) validateAcrid(formats strfmt.Registry) error {

	if err := validate.Required("acrid", "body", m.Acrid); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateAlbum(formats strfmt.Registry) error {
	if swag.IsZero(m.Album) { // not required
		return nil
	}

	if m.Album != nil {
		if err := m.Album.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("album")
			}
			return err
		}
	}

	return nil
}

func (m *Music) validateArtists(formats strfmt.Registry) error {
	if swag.IsZero(m.Artists) { // not required
		return nil
	}

	for i := 0; i < len(m.Artists); i++ {
		if swag.IsZero(m.Artists[i]) { // not required
			continue
		}

		if m.Artists[i] != nil {
			if err := m.Artists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Music) validateContributors(formats strfmt.Registry) error {
	if swag.IsZero(m.Contributors) { // not required
		return nil
	}

	if m.Contributors != nil {
		if err := m.Contributors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contributors")
			}
			return err
		}
	}

	return nil
}

func (m *Music) validateDbBeginTimeOffsetMs(formats strfmt.Registry) error {

	if err := validate.Required("db_begin_time_offset_ms", "body", m.DbBeginTimeOffsetMs); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateDurationMs(formats strfmt.Registry) error {

	if err := validate.Required("duration_ms", "body", m.DurationMs); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateExternalIds(formats strfmt.Registry) error {

	if err := validate.Required("external_ids", "body", m.ExternalIds); err != nil {
		return err
	}

	if m.ExternalIds != nil {
		if err := m.ExternalIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_ids")
			}
			return err
		}
	}

	return nil
}

func (m *Music) validateExternalMetadata(formats strfmt.Registry) error {

	if err := validate.Required("external_metadata", "body", m.ExternalMetadata); err != nil {
		return err
	}

	if m.ExternalMetadata != nil {
		if err := m.ExternalMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *Music) validateGenres(formats strfmt.Registry) error {
	if swag.IsZero(m.Genres) { // not required
		return nil
	}

	for i := 0; i < len(m.Genres); i++ {
		if swag.IsZero(m.Genres[i]) { // not required
			continue
		}

		if m.Genres[i] != nil {
			if err := m.Genres[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("genres" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Music) validateLyrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Lyrics) { // not required
		return nil
	}

	if m.Lyrics != nil {
		if err := m.Lyrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lyrics")
			}
			return err
		}
	}

	return nil
}

func (m *Music) validatePlayOffsetMs(formats strfmt.Registry) error {

	if err := validate.Required("play_offset_ms", "body", m.PlayOffsetMs); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateReleaseByTerritories(formats strfmt.Registry) error {
	if swag.IsZero(m.ReleaseByTerritories) { // not required
		return nil
	}

	for i := 0; i < len(m.ReleaseByTerritories); i++ {
		if swag.IsZero(m.ReleaseByTerritories[i]) { // not required
			continue
		}

		if m.ReleaseByTerritories[i] != nil {
			if err := m.ReleaseByTerritories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("release_by_territories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Music) validateResultFrom(formats strfmt.Registry) error {

	if err := validate.Required("result_from", "body", m.ResultFrom); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateRightsClaim(formats strfmt.Registry) error {
	if swag.IsZero(m.RightsClaim) { // not required
		return nil
	}

	for i := 0; i < len(m.RightsClaim); i++ {
		if swag.IsZero(m.RightsClaim[i]) { // not required
			continue
		}

		if m.RightsClaim[i] != nil {
			if err := m.RightsClaim[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rights_claim" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Music) validateSampleBeginTimeOffsetMs(formats strfmt.Registry) error {

	if err := validate.Required("sample_begin_time_offset_ms", "body", m.SampleBeginTimeOffsetMs); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateSampleEndTimeOffsetMs(formats strfmt.Registry) error {

	if err := validate.Required("sample_end_time_offset_ms", "body", m.SampleEndTimeOffsetMs); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateScore(formats strfmt.Registry) error {

	if err := validate.Required("score", "body", m.Score); err != nil {
		return err
	}

	if err := validate.MinimumInt("score", "body", *m.Score, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("score", "body", *m.Score, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Music) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this music based on the context it is used
func (m *Music) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlbum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArtists(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContributors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGenres(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLyrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReleaseByTerritories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRightsClaim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Music) contextValidateAlbum(ctx context.Context, formats strfmt.Registry) error {

	if m.Album != nil {
		if err := m.Album.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("album")
			}
			return err
		}
	}

	return nil
}

func (m *Music) contextValidateArtists(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Artists); i++ {

		if m.Artists[i] != nil {
			if err := m.Artists[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("artists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Music) contextValidateContributors(ctx context.Context, formats strfmt.Registry) error {

	if m.Contributors != nil {
		if err := m.Contributors.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contributors")
			}
			return err
		}
	}

	return nil
}

func (m *Music) contextValidateExternalIds(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalIds != nil {
		if err := m.ExternalIds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_ids")
			}
			return err
		}
	}

	return nil
}

func (m *Music) contextValidateExternalMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalMetadata != nil {
		if err := m.ExternalMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *Music) contextValidateGenres(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Genres); i++ {

		if m.Genres[i] != nil {
			if err := m.Genres[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("genres" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Music) contextValidateLyrics(ctx context.Context, formats strfmt.Registry) error {

	if m.Lyrics != nil {
		if err := m.Lyrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lyrics")
			}
			return err
		}
	}

	return nil
}

func (m *Music) contextValidateReleaseByTerritories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReleaseByTerritories); i++ {

		if m.ReleaseByTerritories[i] != nil {
			if err := m.ReleaseByTerritories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("release_by_territories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Music) contextValidateRightsClaim(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RightsClaim); i++ {

		if m.RightsClaim[i] != nil {
			if err := m.RightsClaim[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rights_claim" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Music) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Music) UnmarshalBinary(b []byte) error {
	var res Music
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
