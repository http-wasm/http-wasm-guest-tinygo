package handler

import (
	"io"
	"unsafe"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/mem"
)

var (
	wasmRequestBody  api.Body = wasmBody(imports.BodyKindRequest)
	wasmResponseBody api.Body = wasmBody(imports.BodyKindResponse)
)

// wasmBody implements api.Body with imported WebAssembly functions.
type wasmBody imports.BodyKind

// WriteTo implements the same method as documented on api.Body.
func (b wasmBody) WriteTo(w io.Writer) (written uint64, err error) {
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
func (b wasmBody) Read(bytes []byte) (size uint32, eof bool) {
	limit := uint32(len(bytes))
	if limit == 0 { // invalid, but prevent crashing.
		return 0, false
	}

	ptr := uintptr(unsafe.Pointer(&bytes[0])) // TODO: mem.SliceToPtr
	return read(b, ptr, limit)
}

func read(b wasmBody, ptr uintptr, limit imports.BufLimit) (size uint32, eof bool) {
	eofLen := imports.ReadBody(imports.BodyKind(b), ptr, limit)
	eof = (eofLen >> 32) == 1
	size = uint32(eofLen)
	return
}

// Write implements the same method as documented on api.Body.
func (b wasmBody) Write(bytes []byte) {
	size := uint32(len(bytes))
	if size == 0 { // invalid, but prevent crashing.
		return
	}

	ptr := uintptr(unsafe.Pointer(&bytes[0])) // TODO: mem.SliceToPtr
	imports.WriteBody(imports.BodyKind(b), ptr, size)
}

// WriteString implements the same method as documented on api.Body.
func (b wasmBody) WriteString(s string) {
	ptr, size := mem.StringToPtr(s)
	if size == 0 { // invalid, but prevent crashing.
		return
	}

	imports.WriteBody(imports.BodyKind(b), ptr, size)
}
