// Code generated by go-swagger; DO NOT EDIT.

package compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/radiorabe/acr-webhook-receiver/models"
)

// GetCustomStreamReader is a Reader for the GetCustomStream structure.
type GetCustomStreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomStreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomStreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCustomStreamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCustomStreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomStreamOK creates a GetCustomStreamOK with default headers values
func NewGetCustomStreamOK() *GetCustomStreamOK {
	return &GetCustomStreamOK{}
}

/*GetCustomStreamOK handles this case with default header values.

Results without local ID
*/
type GetCustomStreamOK struct {
	Payload []*models.Data
}

func (o *GetCustomStreamOK) Error() string {
	return fmt.Sprintf("[GET /v1/monitor-streams/{streamId}/results][%d] getCustomStreamOK  %+v", 200, o.Payload)
}

func (o *GetCustomStreamOK) GetPayload() []*models.Data {
	return o.Payload
}

func (o *GetCustomStreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomStreamBadRequest creates a GetCustomStreamBadRequest with default headers values
func NewGetCustomStreamBadRequest() *GetCustomStreamBadRequest {
	return &GetCustomStreamBadRequest{}
}

/*GetCustomStreamBadRequest handles this case with default header values.

Bad Request
*/
type GetCustomStreamBadRequest struct {
}

func (o *GetCustomStreamBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/monitor-streams/{streamId}/results][%d] getCustomStreamBadRequest ", 400)
}

func (o *GetCustomStreamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomStreamInternalServerError creates a GetCustomStreamInternalServerError with default headers values
func NewGetCustomStreamInternalServerError() *GetCustomStreamInternalServerError {
	return &GetCustomStreamInternalServerError{}
}

/*GetCustomStreamInternalServerError handles this case with default header values.

Server Error
*/
type GetCustomStreamInternalServerError struct {
}

func (o *GetCustomStreamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/monitor-streams/{streamId}/results][%d] getCustomStreamInternalServerError ", 500)
}

func (o *GetCustomStreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}