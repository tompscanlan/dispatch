///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// IngestEventOKCode is the HTTP code returned for type IngestEventOK
const IngestEventOKCode int = 200

/*IngestEventOK Event emitted

swagger:response ingestEventOK
*/
type IngestEventOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Emission `json:"body,omitempty"`
}

// NewIngestEventOK creates IngestEventOK with default headers values
func NewIngestEventOK() *IngestEventOK {

	return &IngestEventOK{}
}

// WithPayload adds the payload to the ingest event o k response
func (o *IngestEventOK) WithPayload(payload *v1.Emission) *IngestEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ingest event o k response
func (o *IngestEventOK) SetPayload(payload *v1.Emission) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IngestEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// IngestEventBadRequestCode is the HTTP code returned for type IngestEventBadRequest
const IngestEventBadRequestCode int = 400

/*IngestEventBadRequest Invalid input

swagger:response ingestEventBadRequest
*/
type IngestEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewIngestEventBadRequest creates IngestEventBadRequest with default headers values
func NewIngestEventBadRequest() *IngestEventBadRequest {

	return &IngestEventBadRequest{}
}

// WithPayload adds the payload to the ingest event bad request response
func (o *IngestEventBadRequest) WithPayload(payload *v1.Error) *IngestEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ingest event bad request response
func (o *IngestEventBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IngestEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// IngestEventUnauthorizedCode is the HTTP code returned for type IngestEventUnauthorized
const IngestEventUnauthorizedCode int = 401

/*IngestEventUnauthorized Unauthorized Request

swagger:response ingestEventUnauthorized
*/
type IngestEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewIngestEventUnauthorized creates IngestEventUnauthorized with default headers values
func NewIngestEventUnauthorized() *IngestEventUnauthorized {

	return &IngestEventUnauthorized{}
}

// WithPayload adds the payload to the ingest event unauthorized response
func (o *IngestEventUnauthorized) WithPayload(payload *v1.Error) *IngestEventUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ingest event unauthorized response
func (o *IngestEventUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IngestEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// IngestEventForbiddenCode is the HTTP code returned for type IngestEventForbidden
const IngestEventForbiddenCode int = 403

/*IngestEventForbidden access to this resource is forbidden

swagger:response ingestEventForbidden
*/
type IngestEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewIngestEventForbidden creates IngestEventForbidden with default headers values
func NewIngestEventForbidden() *IngestEventForbidden {

	return &IngestEventForbidden{}
}

// WithPayload adds the payload to the ingest event forbidden response
func (o *IngestEventForbidden) WithPayload(payload *v1.Error) *IngestEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ingest event forbidden response
func (o *IngestEventForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IngestEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*IngestEventDefault Unknown error

swagger:response ingestEventDefault
*/
type IngestEventDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewIngestEventDefault creates IngestEventDefault with default headers values
func NewIngestEventDefault(code int) *IngestEventDefault {
	if code <= 0 {
		code = 500
	}

	return &IngestEventDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the ingest event default response
func (o *IngestEventDefault) WithStatusCode(code int) *IngestEventDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the ingest event default response
func (o *IngestEventDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the ingest event default response
func (o *IngestEventDefault) WithPayload(payload *v1.Error) *IngestEventDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ingest event default response
func (o *IngestEventDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IngestEventDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
