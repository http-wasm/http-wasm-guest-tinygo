package handler

import (
	"github.com/tetratelabs/tinymem"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
)

var (
	wasmRequestHeaders = &wasmHeader{
		getNames: imports.GetRequestHeaderNames,
		get:      imports.GetRequestHeader,
		set:      imports.SetRequestHeader,
	}
	wasmRequestTrailers = &wasmHeader{
		getNames: imports.GetRequestTrailerNames,
		get:      imports.GetRequestTrailer,
		set:      imports.SetRequestTrailer,
	}
	wasmRequestBody = &wasmBody{
		read:  imports.ReadRequestBody,
		write: imports.WriteRequestBody,
	}
)

// wasmRequest implements api.Request with imported WebAssembly functions.
type wasmRequest struct{}

// compile-time check to ensure wasmRequest implements api.Request.
var _ api.Request = wasmRequest{}

// GetMethod implements the same method as documented on api.Request.
func (wasmRequest) GetMethod() string {
	return getString(imports.GetMethod)
}

// SetMethod implements the same method as documented on api.Request.
func (wasmRequest) SetMethod(method string) {
	ptr, size := tinymem.StringToPtr(method)
	imports.SetMethod(ptr, size)
}

// GetURI implements the same method as documented on api.Request.
func (wasmRequest) GetURI() string {
	return getString(imports.GetURI)
}

// SetURI implements the same method as documented on api.Request.
func (wasmRequest) SetURI(uri string) {
	ptr, size := tinymem.StringToPtr(uri)
	imports.SetURI(ptr, size)
}

// GetProtocolVersion implements the same method as documented on api.Request.
func (wasmRequest) GetProtocolVersion() string {
	return getString(imports.GetProtocolVersion)
}

// Headers implements the same method as documented on api.Request.
func (wasmRequest) Headers() api.Header {
	return wasmRequestHeaders
}

// Body implements the same method as documented on api.Request.
func (wasmRequest) Body() api.Body {
	return wasmRequestBody
}

// Trailers implements the same method as documented on api.Request.
func (wasmRequest) Trailers() api.Header {
	return wasmRequestTrailers
}
