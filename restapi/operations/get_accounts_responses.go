// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	//"cassandragenrest/models"
	"github.com/stevef1uk/cassandragenrest/models"
)

// GetAccountsOKCode is the HTTP code returned for type GetAccountsOK
const GetAccountsOKCode int = 200

/*GetAccountsOK A list of Accounts

swagger:response getAccountsOK
*/
type GetAccountsOK struct {

	/*
	  In: Body
	*/
	Payload models.GetAccountsOKBody `json:"body,omitempty"`
}

// NewGetAccountsOK creates GetAccountsOK with default headers values
func NewGetAccountsOK() *GetAccountsOK {
	return &GetAccountsOK{}
}

// WithPayload adds the payload to the get accounts o k response
func (o *GetAccountsOK) WithPayload(payload models.GetAccountsOKBody) *GetAccountsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get accounts o k response
func (o *GetAccountsOK) SetPayload(payload models.GetAccountsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetAccountsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
