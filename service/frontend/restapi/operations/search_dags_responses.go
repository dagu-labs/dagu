// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dagu-dev/dagu/service/frontend/models"
)

// SearchDagsOKCode is the HTTP code returned for type SearchDagsOK
const SearchDagsOKCode int = 200

/*
SearchDagsOK A successful response.

swagger:response searchDagsOK
*/
type SearchDagsOK struct {

	/*
	  In: Body
	*/
	Payload *models.SearchDagsResponse `json:"body,omitempty"`
}

// NewSearchDagsOK creates SearchDagsOK with default headers values
func NewSearchDagsOK() *SearchDagsOK {

	return &SearchDagsOK{}
}

// WithPayload adds the payload to the search dags o k response
func (o *SearchDagsOK) WithPayload(payload *models.SearchDagsResponse) *SearchDagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search dags o k response
func (o *SearchDagsOK) SetPayload(payload *models.SearchDagsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchDagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
SearchDagsDefault Generic error response.

swagger:response searchDagsDefault
*/
type SearchDagsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewSearchDagsDefault creates SearchDagsDefault with default headers values
func NewSearchDagsDefault(code int) *SearchDagsDefault {
	if code <= 0 {
		code = 500
	}

	return &SearchDagsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the search dags default response
func (o *SearchDagsDefault) WithStatusCode(code int) *SearchDagsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the search dags default response
func (o *SearchDagsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the search dags default response
func (o *SearchDagsDefault) WithPayload(payload *models.APIError) *SearchDagsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search dags default response
func (o *SearchDagsDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchDagsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
