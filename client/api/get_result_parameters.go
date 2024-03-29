// Code generated by go-swagger; DO NOT EDIT.

package api

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
	"github.com/go-openapi/swag"
)

// NewGetResultParams creates a new GetResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResultParams() *GetResultParams {
	return &GetResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResultParamsWithTimeout creates a new GetResultParams object
// with the ability to set a timeout on a request.
func NewGetResultParamsWithTimeout(timeout time.Duration) *GetResultParams {
	return &GetResultParams{
		timeout: timeout,
	}
}

// NewGetResultParamsWithContext creates a new GetResultParams object
// with the ability to set a context for a request.
func NewGetResultParamsWithContext(ctx context.Context) *GetResultParams {
	return &GetResultParams{
		Context: ctx,
	}
}

// NewGetResultParamsWithHTTPClient creates a new GetResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResultParamsWithHTTPClient(client *http.Client) *GetResultParams {
	return &GetResultParams{
		HTTPClient: client,
	}
}

/* GetResultParams contains all the parameters to send to the API endpoint
   for the get result operation.

   Typically these are written to a http.Request.
*/
type GetResultParams struct {

	// XRequestID.
	XRequestID *string

	// ResultID.
	ResultID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResultParams) WithDefaults() *GetResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get result params
func (o *GetResultParams) WithTimeout(timeout time.Duration) *GetResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get result params
func (o *GetResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get result params
func (o *GetResultParams) WithContext(ctx context.Context) *GetResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get result params
func (o *GetResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get result params
func (o *GetResultParams) WithHTTPClient(client *http.Client) *GetResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get result params
func (o *GetResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get result params
func (o *GetResultParams) WithXRequestID(xRequestID *string) *GetResultParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get result params
func (o *GetResultParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithResultID adds the resultID to the get result params
func (o *GetResultParams) WithResultID(resultID int64) *GetResultParams {
	o.SetResultID(resultID)
	return o
}

// SetResultID adds the resultId to the get result params
func (o *GetResultParams) SetResultID(resultID int64) {
	o.ResultID = resultID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resultId
	if err := r.SetPathParam("resultId", swag.FormatInt64(o.ResultID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
