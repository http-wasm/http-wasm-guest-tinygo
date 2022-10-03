//go:build !tinygo.wasm

package handler

// log is stubbed for compilation outside TinyGo.
func log(ptr uintptr, size uint32) {}

// getPath is stubbed for compilation outside TinyGo.
func getPath(ptr uintptr, bufLimit uint32) uint32 { return 0 }

// setPath is stubbed for compilation outside TinyGo.
func setPath(ptr uintptr, size uint32) {}

// getRequestHeader is stubbed for compilation outside TinyGo.
func getRequestHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLimit uint32) uint64 {
	return 0
}

// next is stubbed for compilation outside TinyGo.
func next() {}

// setResponseHeader is stubbed for compilation outside TinyGo.
func setResponseHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32) {}

// sendResponse is stubbed for compilation outside TinyGo.
func sendResponse(statusCode uint32, bodyPtr uintptr, bodySize uint32) {}
