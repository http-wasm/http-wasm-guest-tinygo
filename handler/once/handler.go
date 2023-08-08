// Package once uses an api.Handler for only one HTTP round trip.
package once

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal"
)

// Handle invokes the handler and returns
func Handle(handler api.Handler) {
	next, reqCtx := handler.HandleRequest(internal.WasmRequest, internal.WasmResponse)
	ctxNext := uint64(reqCtx) << 32
	if next {
		ctxNext |= 1
	}
	isError := awaitResponse(ctxNext)
	isErrorB := false
	if isError == 1 {
		isErrorB = true
	}
	handler.HandleResponse(reqCtx, internal.WasmRequest, internal.WasmResponse, isErrorB)
}
