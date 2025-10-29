package main

import (
	"net/http"
)

// extend this with any other context you want to be available to Arri procedures
// this must fullfil the arri.Context interface
type RpcEvent struct {
	W http.ResponseWriter
	R *http.Request
}

func (c RpcEvent) Request() *http.Request {
	return c.R
}

func (c RpcEvent) Writer() http.ResponseWriter {
	return c.W
}
