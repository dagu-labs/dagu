// Code generated by go-swagger; DO NOT EDIT.

package dags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dagu-org/dagu/internal/frontend/gen/models"
)

// ListHistoryOKCode is the HTTP code returned for type ListHistoryOK
const ListHistoryOKCode int = 200

/*
ListHistoryOK A successful response.

swagger:response listHistoryOK
*/
type ListHistoryOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListHistoryResponse `json:"body,omitempty"`
}

// NewListHistoryOK creates ListHistoryOK with default headers values
func NewListHistoryOK() *ListHistoryOK {

	return &ListHistoryOK{}
}

// WithPayload adds the payload to the list history o k response
func (o *ListHistoryOK) WithPayload(payload *models.ListHistoryResponse) *ListHistoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list history o k response
func (o *ListHistoryOK) SetPayload(payload *models.ListHistoryResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHistoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListHistoryDefault Generic error response.

swagger:response listHistoryDefault
*/
type ListHistoryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewListHistoryDefault creates ListHistoryDefault with default headers values
func NewListHistoryDefault(code int) *ListHistoryDefault {
	if code <= 0 {
		code = 500
	}

	return &ListHistoryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list history default response
func (o *ListHistoryDefault) WithStatusCode(code int) *ListHistoryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list history default response
func (o *ListHistoryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list history default response
func (o *ListHistoryDefault) WithPayload(payload *models.APIError) *ListHistoryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list history default response
func (o *ListHistoryDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHistoryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}