package handler

import (
	"net/url"

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
}

func unserializeURL(u string) *url.URL {
	if u != "" {
		// we skip error parsing as the url string comes from a proper url.URL object
		// hence error is unexpected
		uri, _ := url.Parse(u)
		return uri
	}

	return &url.URL{Path: "/"}
}

// GetURI implements the same method as documented on api.Request.
func (wasmRequest) GetURI() *url.URL {
	return unserializeURL(mem.GetString(imports.GetURI))
}

// SetURI implements the same method as documented on api.Request.
func (wasmRequest) SetURI(uri *url.URL) {
	ptr, size := mem.StringToPtr(uri.String())
	imports.SetURI(ptr, size)
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
