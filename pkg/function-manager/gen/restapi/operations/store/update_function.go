///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateFunctionHandlerFunc turns a function with the right signature into a update function handler
type UpdateFunctionHandlerFunc func(UpdateFunctionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateFunctionHandlerFunc) Handle(params UpdateFunctionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateFunctionHandler interface for that can handle valid update function params
type UpdateFunctionHandler interface {
	Handle(UpdateFunctionParams, interface{}) middleware.Responder
}

// NewUpdateFunction creates a new http.Handler for the update function operation
func NewUpdateFunction(ctx *middleware.Context, handler UpdateFunctionHandler) *UpdateFunction {
	return &UpdateFunction{Context: ctx, Handler: handler}
}

/*UpdateFunction swagger:route PUT /{functionName} Store updateFunction

Update a function

*/
type UpdateFunction struct {
	Context *middleware.Context
	Handler UpdateFunctionHandler
}

func (o *UpdateFunction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateFunctionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}