package handler

import (
	"unsafe"

	"github.com/tetratelabs/tinymem"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/mem"
)

// wasmHeader implements api.Header with imported WebAssembly functions.
type wasmHeader struct {
	getNames func(ptr uintptr, limit uint32) (len uint32)
	get      func(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64)
	set      func(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32)
}

// compile-time check to ensure wasmHeaders implements api.Header.
var _ api.Header = (*wasmHeader)(nil)

// Names implements the same method as documented on api.Request.
func (w *wasmHeader) Names() (names []string) {
	size := w.getNames(mem.ReadBufPtr, mem.ReadBufLimit)
	if size == 0 {
		return
	}
	if size <= mem.ReadBufLimit { // then re-use the mutable buffer.
		return mem.GetNULTerminated(mem.ReadBuf[:size])
	}
	// Otherwise, we have to allocate a new buffer for the large entry.
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = w.getNames(ptr, size)
	return mem.GetNULTerminated(buf)
}

// Get implements the same method as documented on api.Request.
func (w *wasmHeader) Get(name string) (value string, ok bool) {
	namePtr, nameSize := tinymem.StringToPtr(name)
	okLen := w.get(namePtr, nameSize, 0, 0)
	if okLen == 0 {
		return "", false
	}
	size := uint32(okLen)
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = w.get(namePtr, nameSize, ptr, size)
	return string(buf), true
}

// Set implements the same method as documented on api.Request.
func (w *wasmHeader) Set(name, value string) {
	namePtr, nameSize := tinymem.StringToPtr(name)
	valuePtr, valueSize := tinymem.StringToPtr(value)
	w.set(namePtr, nameSize, valuePtr, valueSize)
}
