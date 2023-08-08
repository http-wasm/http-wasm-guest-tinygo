// Package pool uses one api.Handler for multiple requests.
package pool

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal"
)

var currentHandler api.Handler = api.UnimplementedHandler{} //nolint

// SetHandler sets the handler used for all requests.
func SetHandler(handler api.Handler) {
	currentHandler = handler
}

// handleRequest is only exported to the host.
//
//go:export handle_request
func handleRequest() (ctxNext uint64) { //nolint
	next, reqCtx := currentHandler.HandleRequest(internal.WasmRequest, internal.WasmResponse)
	ctxNext = uint64(reqCtx) << 32
	if next {
		ctxNext |= 1
	}
	return
}

// handleResponse is only exported to the host.
//
//go:export handle_response
func handleResponse(reqCtx uint32, isError uint32) { //nolint
	isErrorB := false
	if isError == 1 {
		isErrorB = true
	}
	currentHandler.HandleResponse(reqCtx, internal.WasmRequest, internal.WasmResponse, isErrorB)
}
