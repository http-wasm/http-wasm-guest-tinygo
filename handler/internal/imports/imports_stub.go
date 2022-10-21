//go:build !tinygo.wasm

package imports

import "github.com/http-wasm/http-wasm-guest-tinygo/handler/api"

// enableFeatures is stubbed for compilation outside TinyGo.
func enableFeatures(features uint64) uint64 {
	return features
}

// getConfig is stubbed for compilation outside TinyGo.
func getConfig(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// log is stubbed for compilation outside TinyGo.
func log(level api.LogLevel, ptr uintptr, size uint32) {}

// getMethod is stubbed for compilation outside TinyGo.
func getMethod(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// setMethod is stubbed for compilation outside TinyGo.
func setMethod(ptr uintptr, size uint32) {}

// getURI is stubbed for compilation outside TinyGo.
func getURI(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// setURI is stubbed for compilation outside TinyGo.
func setURI(ptr uintptr, size uint32) {}

// getProtocolVersion is stubbed for compilation outside TinyGo.
func getProtocolVersion(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// getHeaderNames is stubbed for compilation outside TinyGo.
func getHeaderNames(kind HeaderKind, ptr uintptr, limit uint32) (countLen CountLen) {
	return 0
}

// getHeaderValues is stubbed for compilation outside TinyGo.
func getHeaderValues(kind HeaderKind, namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (countLen CountLen) {
	return 0
}

// setHeaderValue is stubbed for compilation outside TinyGo.
func setHeaderValue(kind HeaderKind, namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32) {
}

// readBody is stubbed for compilation outside TinyGo.
func readBody(kind BodyKind, bufPtr uintptr, bufLimit uint32) (eofLen uint64) {
	return 0
}

// writeBody is stubbed for compilation outside TinyGo.
func writeBody(kind BodyKind, bufPtr uintptr, bufLen uint32) {}

// next is stubbed for compilation outside TinyGo.
func next() {}

// getStatusCode is stubbed for compilation outside TinyGo.
func getStatusCode() uint32 {
	return 0
}

// setStatusCode is stubbed for compilation outside TinyGo.
func setStatusCode(statusCode uint32) {}
