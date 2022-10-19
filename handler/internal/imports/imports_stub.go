//go:build !tinygo.wasm

package imports

// enableFeatures is stubbed for compilation outside TinyGo.
func enableFeatures(features uint64) uint64 {
	return features
}

// getConfig is stubbed for compilation outside TinyGo.
func getConfig(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// log is stubbed for compilation outside TinyGo.
func log(ptr uintptr, size uint32) {}

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

// getRequestHeaderNames is stubbed for compilation outside TinyGo.
func getRequestHeaderNames(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// getRequestHeader is stubbed for compilation outside TinyGo.
func getRequestHeader(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return 0
}

// getRequestHeaders is stubbed for compilation outside TinyGo.
func getRequestHeaders(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return 0
}

// setRequestHeader is stubbed for compilation outside TinyGo.
func setRequestHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32) {}

// readRequestBody is stubbed for compilation outside TinyGo.
func readRequestBody(bufPtr uintptr, bufLimit uint32) (eofLen uint64) {
	return 0
}

// writeRequestBody is stubbed for compilation outside TinyGo.
func writeRequestBody(bufPtr uintptr, bufLen uint32) {}

// getRequestTrailerNames is stubbed for compilation outside TinyGo.
func getRequestTrailerNames(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// getRequestTrailer is stubbed for compilation outside TinyGo.
func getRequestTrailer(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return 0
}

// getRequestTrailers is stubbed for compilation outside TinyGo.
func getRequestTrailers(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return 0
}

// setRequestTrailer is stubbed for compilation outside TinyGo.
func setRequestTrailer(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32) {}

// next is stubbed for compilation outside TinyGo.
func next() {}

// getStatusCode is stubbed for compilation outside TinyGo.
func getStatusCode() uint32 {
	return 0
}

// setStatusCode is stubbed for compilation outside TinyGo.
func setStatusCode(statusCode uint32) {}

// getResponseHeaderNames is stubbed for compilation outside TinyGo.
func getResponseHeaderNames(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// getResponseHeader is stubbed for compilation outside TinyGo.
func getResponseHeader(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return 0
}

// getResponseHeaders is stubbed for compilation outside TinyGo.
func getResponseHeaders(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return 0
}

// setResponseHeader is stubbed for compilation outside TinyGo.
func setResponseHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32) {}

// readResponseBody is stubbed for compilation outside TinyGo.
func readResponseBody(bufPtr uintptr, bufLimit uint32) (eofLen uint64) {
	return 0
}

// writeResponseBody is stubbed for compilation outside TinyGo.
func writeResponseBody(bufPtr uintptr, bufLen uint32) {}

// getResponseTrailerNames is stubbed for compilation outside TinyGo.
func getResponseTrailerNames(ptr uintptr, limit uint32) (len uint32) {
	return 0
}

// getResponseTrailer is stubbed for compilation outside TinyGo.
func getResponseTrailer(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return 0
}

// getResponseTrailers is stubbed for compilation outside TinyGo.
func getResponseTrailers(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return 0
}

// setResponseTrailer is stubbed for compilation outside TinyGo.
func setResponseTrailer(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueLen uint32) {}
