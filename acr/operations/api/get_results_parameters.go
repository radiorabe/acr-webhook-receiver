// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetResultsParams creates a new GetResultsParams object
// with the default values initialized.
func NewGetResultsParams() GetResultsParams {

	var (
		// initialize parameters with default values

		limitDefault  = int64(20)
		offsetDefault = int64(0)
	)

	return GetResultsParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetResultsParams contains all the bound params for the get results operation
// typically these are obtained from a http.Request
//
// swagger:parameters getResults
type GetResultsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: header
	*/
	XRequestID *string
	/*
	  In: query
	*/
	From *strfmt.DateTime
	/*The numbers of items to return.
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*The number of items to skip before starting to collect the result set.
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*
	  In: query
	*/
	To *strfmt.DateTime
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetResultsParams() beforehand.
func (o *GetResultsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-ID")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRequestID binds and validates parameter XRequestID from header.
func (o *GetResultsParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XRequestID = &raw

	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetResultsParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("from", "query", "strfmt.DateTime", raw)
	}
	o.From = (value.(*strfmt.DateTime))

	if err := o.validateFrom(formats); err != nil {
		return err
	}

	return nil
}

// validateFrom carries on validations for parameter From
func (o *GetResultsParams) validateFrom(formats strfmt.Registry) error {

	if err := validate.FormatOf("from", "query", "date-time", o.From.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetResultsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetResultsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetResultsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", *o.Limit, 1, false); err != nil {
		return err
	}

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetResultsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetResultsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetResultsParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", *o.Offset, 0, false); err != nil {
		return err
	}

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetResultsParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("to", "query", "strfmt.DateTime", raw)
	}
	o.To = (value.(*strfmt.DateTime))

	if err := o.validateTo(formats); err != nil {
		return err
	}

	return nil
}

// validateTo carries on validations for parameter To
func (o *GetResultsParams) validateTo(formats strfmt.Registry) error {

	if err := validate.FormatOf("to", "query", "date-time", o.To.String(), formats); err != nil {
		return err
	}
	return nil
}
