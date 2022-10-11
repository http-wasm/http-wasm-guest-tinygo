//go:build !tinygo.wasm

package handler

// log is stubbed for compilation outside TinyGo.
func log(ptr uintptr, size uint32) {}

// getURI is stubbed for compilation outside TinyGo.
func getURI(ptr uintptr, bufLimit uint32) uint32 { return 0 }

// setURI is stubbed for compilation outside TinyGo.
func setURI(ptr uintptr, size uint32) {}

// getRequestHeader is stubbed for compilation outside TinyGo.
func getRequestHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLimit uint32) uint64 {
	return 0
}

// next is stubbed for compilation outside TinyGo.
func next() {}

// setResponseHeader is stubbed for compilation outside TinyGo.
func setResponseHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32) {}

// setStatusCode is stubbed for compilation outside TinyGo.
func setStatusCode(statusCode uint32) {}

// setResponseBody is stubbed for compilation outside TinyGo.
func setResponseBody(bodyPtr uintptr, bodySize uint32) {}
