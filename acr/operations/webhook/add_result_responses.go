// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddResultCreatedCode is the HTTP code returned for type AddResultCreated
const AddResultCreatedCode int = 201

/*AddResultCreated Created

swagger:response addResultCreated
*/
type AddResultCreated struct {
}

// NewAddResultCreated creates AddResultCreated with default headers values
func NewAddResultCreated() *AddResultCreated {

	return &AddResultCreated{}
}

// WriteResponse to the client
func (o *AddResultCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddResultBadRequestCode is the HTTP code returned for type AddResultBadRequest
const AddResultBadRequestCode int = 400

/*AddResultBadRequest Bad Request

swagger:response addResultBadRequest
*/
type AddResultBadRequest struct {
}

// NewAddResultBadRequest creates AddResultBadRequest with default headers values
func NewAddResultBadRequest() *AddResultBadRequest {

	return &AddResultBadRequest{}
}

// WriteResponse to the client
func (o *AddResultBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddResultInternalServerErrorCode is the HTTP code returned for type AddResultInternalServerError
const AddResultInternalServerErrorCode int = 500

/*AddResultInternalServerError Server Error

swagger:response addResultInternalServerError
*/
type AddResultInternalServerError struct {
}

// NewAddResultInternalServerError creates AddResultInternalServerError with default headers values
func NewAddResultInternalServerError() *AddResultInternalServerError {

	return &AddResultInternalServerError{}
}

// WriteResponse to the client
func (o *AddResultInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
