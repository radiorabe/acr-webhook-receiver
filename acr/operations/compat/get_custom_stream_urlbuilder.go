// Code generated by go-swagger; DO NOT EDIT.

package compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetCustomStreamURL generates an URL for the get custom stream operation
type GetCustomStreamURL struct {
	StreamID string

	AccessKey *string
	Date      string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCustomStreamURL) WithBasePath(bp string) *GetCustomStreamURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCustomStreamURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetCustomStreamURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/monitor-streams/{streamId}/results"

	streamID := o.StreamID
	if streamID != "" {
		_path = strings.Replace(_path, "{streamId}", streamID, -1)
	} else {
		return nil, errors.New("streamId is required on GetCustomStreamURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var accessKeyQ string
	if o.AccessKey != nil {
		accessKeyQ = *o.AccessKey
	}
	if accessKeyQ != "" {
		qs.Set("access_key", accessKeyQ)
	}

	dateQ := o.Date
	if dateQ != "" {
		qs.Set("date", dateQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetCustomStreamURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetCustomStreamURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetCustomStreamURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetCustomStreamURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetCustomStreamURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetCustomStreamURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
