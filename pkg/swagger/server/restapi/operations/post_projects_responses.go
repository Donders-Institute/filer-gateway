// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Donders-Institute/filer-gateway/pkg/swagger/server/models"
)

// PostProjectsOKCode is the HTTP code returned for type PostProjectsOK
const PostProjectsOKCode int = 200

/*PostProjectsOK success

swagger:response postProjectsOK
*/
type PostProjectsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBodyTaskResource `json:"body,omitempty"`
}

// NewPostProjectsOK creates PostProjectsOK with default headers values
func NewPostProjectsOK() *PostProjectsOK {

	return &PostProjectsOK{}
}

// WithPayload adds the payload to the post projects o k response
func (o *PostProjectsOK) WithPayload(payload *models.ResponseBodyTaskResource) *PostProjectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post projects o k response
func (o *PostProjectsOK) SetPayload(payload *models.ResponseBodyTaskResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostProjectsBadRequestCode is the HTTP code returned for type PostProjectsBadRequest
const PostProjectsBadRequestCode int = 400

/*PostProjectsBadRequest bad request

swagger:response postProjectsBadRequest
*/
type PostProjectsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody400 `json:"body,omitempty"`
}

// NewPostProjectsBadRequest creates PostProjectsBadRequest with default headers values
func NewPostProjectsBadRequest() *PostProjectsBadRequest {

	return &PostProjectsBadRequest{}
}

// WithPayload adds the payload to the post projects bad request response
func (o *PostProjectsBadRequest) WithPayload(payload *models.ResponseBody400) *PostProjectsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post projects bad request response
func (o *PostProjectsBadRequest) SetPayload(payload *models.ResponseBody400) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostProjectsInternalServerErrorCode is the HTTP code returned for type PostProjectsInternalServerError
const PostProjectsInternalServerErrorCode int = 500

/*PostProjectsInternalServerError failure

swagger:response postProjectsInternalServerError
*/
type PostProjectsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody500 `json:"body,omitempty"`
}

// NewPostProjectsInternalServerError creates PostProjectsInternalServerError with default headers values
func NewPostProjectsInternalServerError() *PostProjectsInternalServerError {

	return &PostProjectsInternalServerError{}
}

// WithPayload adds the payload to the post projects internal server error response
func (o *PostProjectsInternalServerError) WithPayload(payload *models.ResponseBody500) *PostProjectsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post projects internal server error response
func (o *PostProjectsInternalServerError) SetPayload(payload *models.ResponseBody500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
