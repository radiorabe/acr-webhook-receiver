// Code generated by go-swagger; DO NOT EDIT.

package compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetCustomStreamParams creates a new GetCustomStreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomStreamParams() *GetCustomStreamParams {
	return &GetCustomStreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomStreamParamsWithTimeout creates a new GetCustomStreamParams object
// with the ability to set a timeout on a request.
func NewGetCustomStreamParamsWithTimeout(timeout time.Duration) *GetCustomStreamParams {
	return &GetCustomStreamParams{
		timeout: timeout,
	}
}

// NewGetCustomStreamParamsWithContext creates a new GetCustomStreamParams object
// with the ability to set a context for a request.
func NewGetCustomStreamParamsWithContext(ctx context.Context) *GetCustomStreamParams {
	return &GetCustomStreamParams{
		Context: ctx,
	}
}

// NewGetCustomStreamParamsWithHTTPClient creates a new GetCustomStreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomStreamParamsWithHTTPClient(client *http.Client) *GetCustomStreamParams {
	return &GetCustomStreamParams{
		HTTPClient: client,
	}
}

/* GetCustomStreamParams contains all the parameters to send to the API endpoint
   for the get custom stream operation.

   Typically these are written to a http.Request.
*/
type GetCustomStreamParams struct {

	// XRequestID.
	XRequestID *string

	/* AccessKey.

	   Ignored but available for compatibility with upstream.
	*/
	AccessKey *string

	/* Date.

	   The UTC date in the format 'YYYYMMDD'
	*/
	Date string

	/* StreamID.

	   Stream ID, default is the non-realtime RaBe program.

	   Default: "s-qXuJARB"
	*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get custom stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomStreamParams) WithDefaults() *GetCustomStreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get custom stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomStreamParams) SetDefaults() {
	var (
		streamIDDefault = string("s-qXuJARB")
	)

	val := GetCustomStreamParams{
		StreamID: streamIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get custom stream params
func (o *GetCustomStreamParams) WithTimeout(timeout time.Duration) *GetCustomStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom stream params
func (o *GetCustomStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom stream params
func (o *GetCustomStreamParams) WithContext(ctx context.Context) *GetCustomStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom stream params
func (o *GetCustomStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom stream params
func (o *GetCustomStreamParams) WithHTTPClient(client *http.Client) *GetCustomStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom stream params
func (o *GetCustomStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get custom stream params
func (o *GetCustomStreamParams) WithXRequestID(xRequestID *string) *GetCustomStreamParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get custom stream params
func (o *GetCustomStreamParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAccessKey adds the accessKey to the get custom stream params
func (o *GetCustomStreamParams) WithAccessKey(accessKey *string) *GetCustomStreamParams {
	o.SetAccessKey(accessKey)
	return o
}

// SetAccessKey adds the accessKey to the get custom stream params
func (o *GetCustomStreamParams) SetAccessKey(accessKey *string) {
	o.AccessKey = accessKey
}

// WithDate adds the date to the get custom stream params
func (o *GetCustomStreamParams) WithDate(date string) *GetCustomStreamParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get custom stream params
func (o *GetCustomStreamParams) SetDate(date string) {
	o.Date = date
}

// WithStreamID adds the streamID to the get custom stream params
func (o *GetCustomStreamParams) WithStreamID(streamID string) *GetCustomStreamParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the get custom stream params
func (o *GetCustomStreamParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-ID
		if err := r.SetHeaderParam("X-Request-ID", *o.XRequestID); err != nil {
			return err
		}
	}

	if o.AccessKey != nil {

		// query param access_key
		var qrAccessKey string

		if o.AccessKey != nil {
			qrAccessKey = *o.AccessKey
		}
		qAccessKey := qrAccessKey
		if qAccessKey != "" {

			if err := r.SetQueryParam("access_key", qAccessKey); err != nil {
				return err
			}
		}
	}

	// query param date
	qrDate := o.Date
	qDate := qrDate
	if qDate != "" {

		if err := r.SetQueryParam("date", qDate); err != nil {
			return err
		}
	}

	// path param streamId
	if err := r.SetPathParam("streamId", o.StreamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
