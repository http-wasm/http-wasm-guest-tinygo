package handler

import (
	"github.com/tetratelabs/tinymem"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
)

// wasmHost implements api.Host with imported WebAssembly functions.
type wasmHost struct{}

// compile-time check to ensure wasmHost implements api.Host.
var _ api.Host = wasmHost{}

// EnableFeatures implements the same method as documented on api.Host.
func (wasmHost) EnableFeatures(features api.Features) api.Features {
	return api.Features(imports.EnableFeatures(uint64(features)))
}

// GetConfig implements the same method as documented on api.Host.
func (wasmHost) GetConfig() []byte {
	return getBytes(imports.GetConfig)
}

// Log implements the same method as documented on api.Host.
func (wasmHost) Log(message string) {
	if len(message) == 0 {
		return // don't incur host call overhead
	}
	ptr, size := tinymem.StringToPtr(message)
	imports.Log(ptr, size)
}
