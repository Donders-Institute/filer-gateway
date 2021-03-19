// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Donders-Institute/filer-gateway/pkg/swagger/server/models"
)

// GetProjectsOKCode is the HTTP code returned for type GetProjectsOK
const GetProjectsOKCode int = 200

/*GetProjectsOK success

swagger:response getProjectsOK
*/
type GetProjectsOK struct {

	/*
	  In: Body
	*/
	Payload models.ResponseBodyProjects `json:"body,omitempty"`
}

// NewGetProjectsOK creates GetProjectsOK with default headers values
func NewGetProjectsOK() *GetProjectsOK {

	return &GetProjectsOK{}
}

// WithPayload adds the payload to the get projects o k response
func (o *GetProjectsOK) WithPayload(payload models.ResponseBodyProjects) *GetProjectsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects o k response
func (o *GetProjectsOK) SetPayload(payload models.ResponseBodyProjects) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ResponseBodyProjects{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetProjectsInternalServerErrorCode is the HTTP code returned for type GetProjectsInternalServerError
const GetProjectsInternalServerErrorCode int = 500

/*GetProjectsInternalServerError failure

swagger:response getProjectsInternalServerError
*/
type GetProjectsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody500 `json:"body,omitempty"`
}

// NewGetProjectsInternalServerError creates GetProjectsInternalServerError with default headers values
func NewGetProjectsInternalServerError() *GetProjectsInternalServerError {

	return &GetProjectsInternalServerError{}
}

// WithPayload adds the payload to the get projects internal server error response
func (o *GetProjectsInternalServerError) WithPayload(payload *models.ResponseBody500) *GetProjectsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get projects internal server error response
func (o *GetProjectsInternalServerError) SetPayload(payload *models.ResponseBody500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
