// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dagu-dev/dagu/service/frontend/models"
)

// ListDagsOKCode is the HTTP code returned for type ListDagsOK
const ListDagsOKCode int = 200

/*
ListDagsOK A successful response.

swagger:response listDagsOK
*/
type ListDagsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListDagsResponse `json:"body,omitempty"`
}

// NewListDagsOK creates ListDagsOK with default headers values
func NewListDagsOK() *ListDagsOK {

	return &ListDagsOK{}
}

// WithPayload adds the payload to the list dags o k response
func (o *ListDagsOK) WithPayload(payload *models.ListDagsResponse) *ListDagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list dags o k response
func (o *ListDagsOK) SetPayload(payload *models.ListDagsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListDagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListDagsDefault Generic error response.

swagger:response listDagsDefault
*/
type ListDagsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewListDagsDefault creates ListDagsDefault with default headers values
func NewListDagsDefault(code int) *ListDagsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListDagsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list dags default response
func (o *ListDagsDefault) WithStatusCode(code int) *ListDagsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list dags default response
func (o *ListDagsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list dags default response
func (o *ListDagsDefault) WithPayload(payload *models.APIError) *ListDagsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list dags default response
func (o *ListDagsDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListDagsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
