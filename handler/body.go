package handler

import (
	"unsafe"

	"github.com/tetratelabs/tinymem"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

// wasmBody implements api.Body with imported WebAssembly functions.
type wasmBody struct {
	read  func(bufPtr uintptr, bufLimit uint32) (eofLen uint64)
	write func(bufPtr uintptr, bufLen uint32)
}

// compile-time check to ensure wasmBody implements api.Body.
var _ api.Body = (*wasmBody)(nil)

// readBuf is sharable because there is no parallelism in wasm.
var readBuf = make([]byte, 2048)

// ReadAll implements the same method as documented on api.Body.
func (w *wasmBody) ReadAll() (result []byte) {
	size, eof := w.Read(readBuf)
	if size > 0 {
		result = make([]byte, size)
		copy(result, readBuf)
	}
	for !eof {
		size, eof = w.Read(readBuf)
		result = append(result, readBuf[0:size]...)
	}
	return
}

// ReadN implements the same method as documented on api.Body.
func (w *wasmBody) Read(bytes []byte) (size uint32, eof bool) {
	limit := uint32(len(bytes))
	if limit == 0 { // invalid, but prevent crashing.
		return 0, false
	}

	ptr := uintptr(unsafe.Pointer(&bytes[0])) // TODO: tinymem.SliceToPtr
	eofLen := w.read(ptr, limit)
	eof = (eofLen >> 32) == 1
	size = uint32(eofLen)
	return
}

// Write implements the same method as documented on api.Body.
func (w *wasmBody) Write(bytes []byte) {
	size := uint32(len(bytes))
	if size == 0 { // invalid, but prevent crashing.
		return
	}

	ptr := uintptr(unsafe.Pointer(&bytes[0])) // TODO: tinymem.SliceToPtr
	w.write(ptr, size)
}

// WriteString implements the same method as documented on api.Body.
func (w *wasmBody) WriteString(s string) {
	ptr, size := tinymem.StringToPtr(s)
	if size == 0 { // invalid, but prevent crashing.
		return
	}

	w.write(ptr, size)
}
