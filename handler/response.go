package handler

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
)

var (
	wasmResponseHeaders = &wasmHeader{
		getNames: imports.GetResponseHeaderNames,
		get:      imports.GetResponseHeader,
		getAll:   imports.GetResponseHeaders,
		set:      imports.SetResponseHeader,
	}
	wasmResponseTrailers = &wasmHeader{
		getNames: imports.GetResponseTrailerNames,
		get:      imports.GetResponseTrailer,
		getAll:   imports.GetResponseTrailers,
		set:      imports.SetResponseTrailer,
	}
	wasmResponseBody = &wasmBody{
		read:  imports.ReadResponseBody,
		write: imports.WriteResponseBody,
	}
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
	return wasmResponseHeaders
}

// Body implements the same method as documented on api.Response.
func (wasmResponse) Body() api.Body {
	return wasmResponseBody
}

// Trailers implements the same method as documented on api.Response.
func (wasmResponse) Trailers() api.Header {
	return wasmResponseTrailers
}
