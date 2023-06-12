package mem

import (
	"runtime"
	"unsafe"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler/internal/imports"
)

var (
	// ReadBuf is sharable because there is no parallelism in wasm.
	ReadBuf = make([]byte, ReadBufLimit)
	// ReadBufPtr is used to avoid duplicate host function calls.
	ReadBufPtr = uint32(uintptr(unsafe.Pointer(&ReadBuf[0])))
	// ReadBufLimit is constant memory overhead for reading fields.
	ReadBufLimit = uint32(2048)
)

// SliceToPtr returns a pointer and size pair for the given slice in a way
// compatible with WebAssembly numeric types.
// The returned pointer aliases the slice hence the slice must be kept alive
// until ptr is no longer needed.
func SliceToPtr(b []byte) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.SliceData(b))
	return uint32(uintptr(ptr)), uint32(len(b))
}

// StringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
// The returned pointer aliases the string hence the string must be kept alive
// until ptr is no longer needed.
func StringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}

// GetString copies a string from the bytes returned by fn, so that it can
// safely be used without risk of corruption.
func GetString(fn func(ptr uint32, limit imports.BufLimit) (len uint32)) (result string) {
	size := fn(ReadBufPtr, ReadBufLimit)
	if size == 0 {
		return // If nothing was read, return an empty string.
	} else if size <= ReadBufLimit {
		return string(ReadBuf[:size]) // string will copy the buffer.
	}

	// Otherwise, allocate a new string
	buf := make([]byte, size)
	ptr := unsafe.Pointer(unsafe.SliceData(buf))
	_ = fn(uint32(uintptr(ptr)), size)
	result = *(*string)(ptr) // don't return string(buf) as that copies buf.
	runtime.KeepAlive(buf)   // keep buf alive until ptr is no longer needed.
	return
}

// GetBytes copies the bytes returned by fn, so that they can safely be used
// without risk of corruption.
func GetBytes(fn func(ptr uint32, limit imports.BufLimit) (len uint32)) (result []byte) {
	size := fn(ReadBufPtr, ReadBufLimit)
	if size == 0 {
		return // If nothing was read, return a nil slice.
	} else if size <= ReadBufLimit {
		// copy to avoid passing out our read buffer
		result = make([]byte, size)
		copy(result, ReadBuf)
		return
	}

	result = make([]byte, size)
	ptr := unsafe.Pointer(unsafe.SliceData(result))
	_ = fn(uint32(uintptr(ptr)), size)
	return
}

func GetNULTerminated(b []byte) (entries []string) {
	if len(b) == 0 {
		return
	}
	entry, entryPos := 0, 0
	for size := len(b); size > 0; size-- {
		if b[entryPos] == 0 { // then, we reached the end of the field.
			entries = append(entries, string(b[entry:entryPos]))
			entryPos++
			entry = entryPos
		} else {
			entryPos++
		}
	}
	return
}
