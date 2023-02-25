// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/webcustomerapi1/models"
)

// PostcustomerOKCode is the HTTP code returned for type PostcustomerOK
const PostcustomerOKCode int = 200

/*
PostcustomerOK returns a greeting

swagger:response postcustomerOK
*/
type PostcustomerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Customer `json:"body,omitempty"`
}

// NewPostcustomerOK creates PostcustomerOK with default headers values
func NewPostcustomerOK() *PostcustomerOK {

	return &PostcustomerOK{}
}

// WithPayload adds the payload to the postcustomer o k response
func (o *PostcustomerOK) WithPayload(payload *models.Customer) *PostcustomerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the postcustomer o k response
func (o *PostcustomerOK) SetPayload(payload *models.Customer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostcustomerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
