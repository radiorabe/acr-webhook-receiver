// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/radiorabe/acr-webhook-receiver/models"
)

// GetResultsReader is a Reader for the GetResults structure.
type GetResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetResultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResultsOK creates a GetResultsOK with default headers values
func NewGetResultsOK() *GetResultsOK {
	return &GetResultsOK{}
}

/*GetResultsOK handles this case with default header values.

Returns array of results
*/
type GetResultsOK struct {
	Payload []*models.Result
}

func (o *GetResultsOK) Error() string {
	return fmt.Sprintf("[GET /v1/results][%d] getResultsOK  %+v", 200, o.Payload)
}

func (o *GetResultsOK) GetPayload() []*models.Result {
	return o.Payload
}

func (o *GetResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultsInternalServerError creates a GetResultsInternalServerError with default headers values
func NewGetResultsInternalServerError() *GetResultsInternalServerError {
	return &GetResultsInternalServerError{}
}

/*GetResultsInternalServerError handles this case with default header values.

Server Error
*/
type GetResultsInternalServerError struct {
}

func (o *GetResultsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/results][%d] getResultsInternalServerError ", 500)
}

func (o *GetResultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
