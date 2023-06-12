package handler

import (
	"runtime"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/mem"
)

// wasmHost implements api.Host with imported WebAssembly functions.
type wasmHost struct{}

// compile-time check to ensure wasmHost implements api.Host.
var _ api.Host = wasmHost{}

// EnableFeatures implements the same method as documented on api.Host.
func (wasmHost) EnableFeatures(features api.Features) api.Features {
	return imports.EnableFeatures(features)
}

// GetConfig implements the same method as documented on api.Host.
func (wasmHost) GetConfig() []byte {
	return mem.GetBytes(imports.GetConfig)
}

// LogEnabled implements the same method as documented on api.Host.
func (h wasmHost) LogEnabled(level api.LogLevel) bool {
	if enabled := imports.LogEnabled(level); enabled == 1 {
		return true
	}
	return false
}

// Log implements the same method as documented on api.Host.
func (wasmHost) Log(level api.LogLevel, message string) {
	if len(message) == 0 {
		return // don't incur host call overhead
	}
	ptr, size := mem.StringToPtr(message)
	imports.Log(level, ptr, size)
	runtime.KeepAlive(message) // keep message alive until ptr is no longer needed.
}
