// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Donders-Institute/filer-gateway/pkg/swagger/server/models"
)

// GetProjectsIDStorageOKCode is the HTTP code returned for type GetProjectsIDStorageOK
const GetProjectsIDStorageOKCode int = 200

/*GetProjectsIDStorageOK success

swagger:response getProjectsIdStorageOK
*/
type GetProjectsIDStorageOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBodyStorageResource `json:"body,omitempty"`
}

// NewGetProjectsIDStorageOK creates GetProjectsIDStorageOK with default headers values
func NewGetProjectsIDStorageOK() *GetProjectsIDStorageOK {

	return &GetProjectsIDStorageOK{}
}

// WithPayload adds the payload to the get projects Id storage o k response
func (o *GetProjectsIDStorageOK) WithPayload(payload *models.ResponseBodyStorageResource) *GetProjectsIDStorageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects Id storage o k response
func (o *GetProjectsIDStorageOK) SetPayload(payload *models.ResponseBodyStorageResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsIDStorageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectsIDStorageBadRequestCode is the HTTP code returned for type GetProjectsIDStorageBadRequest
const GetProjectsIDStorageBadRequestCode int = 400

/*GetProjectsIDStorageBadRequest bad request

swagger:response getProjectsIdStorageBadRequest
*/
type GetProjectsIDStorageBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody400 `json:"body,omitempty"`
}

// NewGetProjectsIDStorageBadRequest creates GetProjectsIDStorageBadRequest with default headers values
func NewGetProjectsIDStorageBadRequest() *GetProjectsIDStorageBadRequest {

	return &GetProjectsIDStorageBadRequest{}
}

// WithPayload adds the payload to the get projects Id storage bad request response
func (o *GetProjectsIDStorageBadRequest) WithPayload(payload *models.ResponseBody400) *GetProjectsIDStorageBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects Id storage bad request response
func (o *GetProjectsIDStorageBadRequest) SetPayload(payload *models.ResponseBody400) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsIDStorageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectsIDStorageNotFoundCode is the HTTP code returned for type GetProjectsIDStorageNotFound
const GetProjectsIDStorageNotFoundCode int = 404

/*GetProjectsIDStorageNotFound project not found

swagger:response getProjectsIdStorageNotFound
*/
type GetProjectsIDStorageNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetProjectsIDStorageNotFound creates GetProjectsIDStorageNotFound with default headers values
func NewGetProjectsIDStorageNotFound() *GetProjectsIDStorageNotFound {

	return &GetProjectsIDStorageNotFound{}
}

// WithPayload adds the payload to the get projects Id storage not found response
func (o *GetProjectsIDStorageNotFound) WithPayload(payload string) *GetProjectsIDStorageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects Id storage not found response
func (o *GetProjectsIDStorageNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsIDStorageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetProjectsIDStorageInternalServerErrorCode is the HTTP code returned for type GetProjectsIDStorageInternalServerError
const GetProjectsIDStorageInternalServerErrorCode int = 500

/*GetProjectsIDStorageInternalServerError failure

swagger:response getProjectsIdStorageInternalServerError
*/
type GetProjectsIDStorageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody500 `json:"body,omitempty"`
}

// NewGetProjectsIDStorageInternalServerError creates GetProjectsIDStorageInternalServerError with default headers values
func NewGetProjectsIDStorageInternalServerError() *GetProjectsIDStorageInternalServerError {

	return &GetProjectsIDStorageInternalServerError{}
}

// WithPayload adds the payload to the get projects Id storage internal server error response
func (o *GetProjectsIDStorageInternalServerError) WithPayload(payload *models.ResponseBody500) *GetProjectsIDStorageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects Id storage internal server error response
func (o *GetProjectsIDStorageInternalServerError) SetPayload(payload *models.ResponseBody500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsIDStorageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
