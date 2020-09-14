// Code generated by go-swagger; DO NOT EDIT.

package compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetCustomStreamParams creates a new GetCustomStreamParams object
// with the default values initialized.
func NewGetCustomStreamParams() GetCustomStreamParams {

	var (
		// initialize parameters with default values

		streamIDDefault = string("s-qXuJARB")
	)

	return GetCustomStreamParams{
		StreamID: streamIDDefault,
	}
}

// GetCustomStreamParams contains all the bound params for the get custom stream operation
// typically these are obtained from a http.Request
//
// swagger:parameters getCustomStream
type GetCustomStreamParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Ignored but available for compatibility with upstream.
	  In: query
	*/
	AccessKey *string
	/*The UTC date in the format 'YYYYMMDD'
	  Required: true
	  In: query
	*/
	Date string
	/*Stream ID, default is the non-realtime RaBe program.
	  Required: true
	  In: path
	  Default: "s-qXuJARB"
	*/
	StreamID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCustomStreamParams() beforehand.
func (o *GetCustomStreamParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAccessKey, qhkAccessKey, _ := qs.GetOK("access_key")
	if err := o.bindAccessKey(qAccessKey, qhkAccessKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qDate, qhkDate, _ := qs.GetOK("date")
	if err := o.bindDate(qDate, qhkDate, route.Formats); err != nil {
		res = append(res, err)
	}

	rStreamID, rhkStreamID, _ := route.Params.GetOK("streamId")
	if err := o.bindStreamID(rStreamID, rhkStreamID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAccessKey binds and validates parameter AccessKey from query.
func (o *GetCustomStreamParams) bindAccessKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.AccessKey = &raw

	return nil
}

// bindDate binds and validates parameter Date from query.
func (o *GetCustomStreamParams) bindDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("date", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("date", "query", raw); err != nil {
		return err
	}

	o.Date = raw

	return nil
}

// bindStreamID binds and validates parameter StreamID from path.
func (o *GetCustomStreamParams) bindStreamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.StreamID = raw

	return nil
}
