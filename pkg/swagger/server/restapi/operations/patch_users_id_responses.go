// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Donders-Institute/filer-gateway/pkg/swagger/server/models"
)

// PatchUsersIDOKCode is the HTTP code returned for type PatchUsersIDOK
const PatchUsersIDOKCode int = 200

/*PatchUsersIDOK success

swagger:response patchUsersIdOK
*/
type PatchUsersIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBodyTaskResource `json:"body,omitempty"`
}

// NewPatchUsersIDOK creates PatchUsersIDOK with default headers values
func NewPatchUsersIDOK() *PatchUsersIDOK {

	return &PatchUsersIDOK{}
}

// WithPayload adds the payload to the patch users Id o k response
func (o *PatchUsersIDOK) WithPayload(payload *models.ResponseBodyTaskResource) *PatchUsersIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch users Id o k response
func (o *PatchUsersIDOK) SetPayload(payload *models.ResponseBodyTaskResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUsersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUsersIDBadRequestCode is the HTTP code returned for type PatchUsersIDBadRequest
const PatchUsersIDBadRequestCode int = 400

/*PatchUsersIDBadRequest bad request

swagger:response patchUsersIdBadRequest
*/
type PatchUsersIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody400 `json:"body,omitempty"`
}

// NewPatchUsersIDBadRequest creates PatchUsersIDBadRequest with default headers values
func NewPatchUsersIDBadRequest() *PatchUsersIDBadRequest {

	return &PatchUsersIDBadRequest{}
}

// WithPayload adds the payload to the patch users Id bad request response
func (o *PatchUsersIDBadRequest) WithPayload(payload *models.ResponseBody400) *PatchUsersIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch users Id bad request response
func (o *PatchUsersIDBadRequest) SetPayload(payload *models.ResponseBody400) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUsersIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUsersIDNotFoundCode is the HTTP code returned for type PatchUsersIDNotFound
const PatchUsersIDNotFoundCode int = 404

/*PatchUsersIDNotFound user not found

swagger:response patchUsersIdNotFound
*/
type PatchUsersIDNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPatchUsersIDNotFound creates PatchUsersIDNotFound with default headers values
func NewPatchUsersIDNotFound() *PatchUsersIDNotFound {

	return &PatchUsersIDNotFound{}
}

// WithPayload adds the payload to the patch users Id not found response
func (o *PatchUsersIDNotFound) WithPayload(payload string) *PatchUsersIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch users Id not found response
func (o *PatchUsersIDNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUsersIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchUsersIDInternalServerErrorCode is the HTTP code returned for type PatchUsersIDInternalServerError
const PatchUsersIDInternalServerErrorCode int = 500

/*PatchUsersIDInternalServerError failure

swagger:response patchUsersIdInternalServerError
*/
type PatchUsersIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody500 `json:"body,omitempty"`
}

// NewPatchUsersIDInternalServerError creates PatchUsersIDInternalServerError with default headers values
func NewPatchUsersIDInternalServerError() *PatchUsersIDInternalServerError {

	return &PatchUsersIDInternalServerError{}
}

// WithPayload adds the payload to the patch users Id internal server error response
func (o *PatchUsersIDInternalServerError) WithPayload(payload *models.ResponseBody500) *PatchUsersIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch users Id internal server error response
func (o *PatchUsersIDInternalServerError) SetPayload(payload *models.ResponseBody500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUsersIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
