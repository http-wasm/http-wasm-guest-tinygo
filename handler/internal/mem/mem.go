package mem

import "unsafe"

var (
	// ReadBuf is sharable because there is no parallelism in wasm.
	ReadBuf = make([]byte, ReadBufLimit)
	// ReadBufPtr is used to avoid duplicate host function calls.
	ReadBufPtr = uintptr(unsafe.Pointer(&ReadBuf[0]))
	// ReadBufLimit is constant memory overhead for reading fields.
	ReadBufLimit = uint32(2048)
)

// StringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
func StringToPtr(s string) (uintptr, uint32) {
	if s == "" {
		return ReadBufPtr, 0
	}
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return unsafePtr, uint32(len(buf))
}

func GetString(fn func(ptr uintptr, limit uint32) (len uint32)) (result string) {
	size := fn(ReadBufPtr, ReadBufLimit)
	if size == 0 {
		return
	}
	if size > 0 && size <= ReadBufLimit {
		return string(ReadBuf[:size]) // string will copy the buffer.
	}

	// Otherwise, allocate a new string
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = fn(ptr, size)
	s := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size)
	return *(*string)(unsafe.Pointer(&s))
}

func GetBytes(fn func(ptr uintptr, limit uint32) (len uint32)) (result []byte) {
	size := fn(ReadBufPtr, ReadBufLimit)
	if size == 0 {
		return
	}
	if size > 0 && size <= ReadBufLimit {
		// copy to avoid passing a mutable buffer
		result = make([]byte, size)
		copy(result, ReadBuf)
		return
	}
	buf := make([]byte, size)
	ptr := uintptr(unsafe.Pointer(&buf[0]))
	_ = fn(ptr, size)
	return buf
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
