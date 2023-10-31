package handler

import (
	"runtime"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/mem"
)

// wasmRequest implements api.Request with imported WebAssembly functions.
type wasmRequest struct{}

// compile-time check to ensure wasmRequest implements api.Request.
var _ api.Request = wasmRequest{}

// GetMethod implements the same method as documented on api.Request.
func (wasmRequest) GetMethod() string {
	return mem.GetString(imports.GetMethod)
}

// SetMethod implements the same method as documented on api.Request.
func (wasmRequest) SetMethod(method string) {
	ptr, size := mem.StringToPtr(method)
	imports.SetMethod(ptr, size)
	runtime.KeepAlive(method) // keep method alive until ptr is no longer needed.
}

// GetURI implements the same method as documented on api.Request.
func (wasmRequest) GetURI() string {
	return mem.GetString(imports.GetURI)
}

// SetURI implements the same method as documented on api.Request.
func (wasmRequest) SetURI(uri string) {
	ptr, size := mem.StringToPtr(uri)
	imports.SetURI(ptr, size)
	runtime.KeepAlive(uri) // keep uri alive until ptr is no longer needed.
}

// GetProtocolVersion implements the same method as documented on api.Request.
func (wasmRequest) GetProtocolVersion() string {
	return mem.GetString(imports.GetProtocolVersion)
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

// GetSourceAddr implements the same method as documented on api.Request.
func (wasmRequest) GetSourceAddr() string {
	return mem.GetString(imports.GetSourceAddr)
}
