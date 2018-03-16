// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package server

import (
	calcsvc "goa.design/goa/examples/calc/gen/calc"
)

// NewAddAddPayload builds a calc service add endpoint payload.
func NewAddAddPayload(a int, b int) *calcsvc.AddPayload {
	return &calcsvc.AddPayload{
		A: a,
		B: b,
	}
}
