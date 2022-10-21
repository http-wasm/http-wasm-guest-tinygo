package handler

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
)

// wasmResponse implements api.Response with imported WebAssembly functions.
type wasmResponse struct{}

// compile-time check to ensure wasmResponse implements api.Response.
var _ api.Response = wasmResponse{}

// GetStatusCode implements the same method as documented on api.Response.
func (r wasmResponse) GetStatusCode() uint32 {
	return imports.GetStatusCode()
}

// SetStatusCode implements the same method as documented on api.Response.
func (r wasmResponse) SetStatusCode(statusCode uint32) {
	imports.SetStatusCode(statusCode)
}

// Headers implements the same method as documented on api.Response.
func (wasmResponse) Headers() api.Header {
	return wasmHeaders
}

// Body implements the same method as documented on api.Response.
func (wasmResponse) Body() api.Body {
	return wasmResponseBody
}

// Trailers implements the same method as documented on api.Response.
func (wasmResponse) Trailers() api.Header {
	return wasmResponseTrailers
}
