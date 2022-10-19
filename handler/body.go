package handler

import (
	"io"
	"unsafe"

	"github.com/tetratelabs/tinymem"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/mem"
)

// wasmBody implements api.Body with imported WebAssembly functions.
type wasmBody struct {
	read  func(bufPtr uintptr, bufLimit uint32) (eofLen uint64)
	write func(bufPtr uintptr, bufLen uint32)
}

// compile-time check to ensure wasmBody implements api.Body.
var _ api.Body = (*wasmBody)(nil)

// WriteTo implements the same method as documented on api.Body.
func (b *wasmBody) WriteTo(w io.Writer) (written uint64, err error) {
	var size uint32
	var eof bool
	var n int

	for !eof {
		size, eof = read(b, mem.ReadBufPtr, mem.ReadBufLimit)
		if size == 0 { // possible for zero length on EOF
			break
		}

		n, err = w.Write(mem.ReadBuf[:size])
		written += uint64(n)
		if err != nil {
			break
		}
	}
	return
}

// ReadN implements the same method as documented on api.Body.
func (b *wasmBody) Read(bytes []byte) (size uint32, eof bool) {
	limit := uint32(len(bytes))
	if limit == 0 { // invalid, but prevent crashing.
		return 0, false
	}

	ptr := uintptr(unsafe.Pointer(&bytes[0])) // TODO: tinymem.SliceToPtr
	return read(b, ptr, limit)
}

func read(w *wasmBody, ptr uintptr, limit uint32) (size uint32, eof bool) {
	eofLen := w.read(ptr, limit)
	eof = (eofLen >> 32) == 1
	size = uint32(eofLen)
	return
}

// Write implements the same method as documented on api.Body.
func (b *wasmBody) Write(bytes []byte) {
	size := uint32(len(bytes))
	if size == 0 { // invalid, but prevent crashing.
		return
	}

	ptr := uintptr(unsafe.Pointer(&bytes[0])) // TODO: tinymem.SliceToPtr
	b.write(ptr, size)
}

// WriteString implements the same method as documented on api.Body.
func (b *wasmBody) WriteString(s string) {
	ptr, size := tinymem.StringToPtr(s)
	if size == 0 { // invalid, but prevent crashing.
		return
	}

	b.write(ptr, size)
}
