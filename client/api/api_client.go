// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetResult(params *GetResultParams) (*GetResultOK, error)

	GetResults(params *GetResultsParams) (*GetResultsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetResult ACRs cloud result

  Use this endpoint to fetch information on an exact entry.
*/
func (a *Client) GetResult(params *GetResultParams) (*GetResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResult",
		Method:             "GET",
		PathPattern:        "/v1/results/{resultId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetResults gets ACR cloud results

  This is endpoint is useful for looking into and exporting the dataset.
*/
func (a *Client) GetResults(params *GetResultsParams) (*GetResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResultsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResults",
		Method:             "GET",
		PathPattern:        "/v1/results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetResultsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}