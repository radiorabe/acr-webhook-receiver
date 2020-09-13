// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddResultReader is a Reader for the AddResult structure.
type AddResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddResultCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddResultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddResultCreated creates a AddResultCreated with default headers values
func NewAddResultCreated() *AddResultCreated {
	return &AddResultCreated{}
}

/*AddResultCreated handles this case with default header values.

Created
*/
type AddResultCreated struct {
}

func (o *AddResultCreated) Error() string {
	return fmt.Sprintf("[POST /v1/_webhooks/results][%d] addResultCreated ", 201)
}

func (o *AddResultCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddResultBadRequest creates a AddResultBadRequest with default headers values
func NewAddResultBadRequest() *AddResultBadRequest {
	return &AddResultBadRequest{}
}

/*AddResultBadRequest handles this case with default header values.

Bad Request
*/
type AddResultBadRequest struct {
}

func (o *AddResultBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/_webhooks/results][%d] addResultBadRequest ", 400)
}

func (o *AddResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddResultInternalServerError creates a AddResultInternalServerError with default headers values
func NewAddResultInternalServerError() *AddResultInternalServerError {
	return &AddResultInternalServerError{}
}

/*AddResultInternalServerError handles this case with default header values.

Server Error
*/
type AddResultInternalServerError struct {
}

func (o *AddResultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/_webhooks/results][%d] addResultInternalServerError ", 500)
}

func (o *AddResultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
