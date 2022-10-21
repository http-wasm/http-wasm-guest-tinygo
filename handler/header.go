package handler

import (
	"unsafe"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/mem"
)

// wasmHeader implements api.Header with imported WebAssembly functions.
type wasmHeader imports.HeaderKind

var (
	wasmRequestHeaders   api.Header = wasmHeader(imports.HeaderKindRequest)
	wasmRequestTrailers  api.Header = wasmHeader(imports.HeaderKindRequestTrailers)
	wasmHeaders          api.Header = wasmHeader(imports.HeaderKindResponse)
	wasmResponseTrailers api.Header = wasmHeader(imports.HeaderKindResponseTrailers)
)

// Names implements the same method as documented on api.Request.
func (w wasmHeader) Names() (names []string) {
	countLen := imports.GetHeaderNames(imports.HeaderKind(w), mem.ReadBufPtr, mem.ReadBufLimit)
	if countLen == 0 {
		return
	}
	size := uint32(countLen)
	if size <= mem.ReadBufLimit { // then re-use the mutable buffer.
		return mem.GetNULTerminated(mem.ReadBuf[:size])
	}
	// Otherwise, we have to allocate a new buffer for the large entry.
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = imports.GetHeaderNames(imports.HeaderKind(w), ptr, size)
	return mem.GetNULTerminated(buf)
}

// Get implements the same method as documented on api.Request.
func (w wasmHeader) Get(name string) (value string, ok bool) {
	values := w.GetAll(name)
	if len(values) > 0 {
		value = values[0]
		ok = true
	}
	return
}

// GetAll implements the same method as documented on api.Request.
func (w wasmHeader) GetAll(name string) (names []string) {
	namePtr, nameSize := mem.StringToPtr(name)
	countLen := imports.GetHeaderValues(imports.HeaderKind(w), namePtr, nameSize, mem.ReadBufPtr, mem.ReadBufLimit)
	if countLen == 0 {
		return
	}
	size := uint32(countLen)
	if size == 0 {
		return
	}
	if size <= mem.ReadBufLimit { // then re-use the mutable buffer.
		return mem.GetNULTerminated(mem.ReadBuf[:size])
	}
	// Otherwise, we have to allocate a new buffer for the large entry.
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = imports.GetHeaderValues(imports.HeaderKind(w), namePtr, nameSize, ptr, size)
	return mem.GetNULTerminated(buf)
}

// Set implements the same method as documented on api.Request.
func (w wasmHeader) Set(name, value string) {
	namePtr, nameSize := mem.StringToPtr(name)
	valuePtr, valueSize := mem.StringToPtr(value)
	imports.SetHeaderValue(imports.HeaderKind(w), namePtr, nameSize, valuePtr, valueSize)
}
