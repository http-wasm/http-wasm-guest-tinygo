package handler

import (
	"io"
	"runtime"

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
	ptr, limit := mem.SliceToPtr(bytes)
	if limit == 0 {
		return // invalid, but prevent crashing.
	}

	size, eof = read(b, ptr, limit)
	runtime.KeepAlive(bytes) // keep bytes alive until ptr is no longer needed.
	return
}

func read(b wasmBody, ptr uint32, limit imports.BufLimit) (size uint32, eof bool) {
	eofLen := imports.ReadBody(imports.BodyKind(b), ptr, limit)
	eof = (eofLen >> 32) == 1
	size = uint32(eofLen)
	return
}

// Write implements the same method as documented on api.Body.
func (b wasmBody) Write(bytes []byte) {
	ptr, size := mem.SliceToPtr(bytes)
	if size == 0 {
		return
	}

	imports.WriteBody(imports.BodyKind(b), ptr, size)
	runtime.KeepAlive(bytes) // keep bytes alive until ptr is no longer needed.
}

// WriteString implements the same method as documented on api.Body.
func (b wasmBody) WriteString(s string) {
	ptr, size := mem.StringToPtr(s)
	if size == 0 {
		return
	}

	imports.WriteBody(imports.BodyKind(b), ptr, size)
	runtime.KeepAlive(s) // keep s alive until ptr is no longer needed.
}
