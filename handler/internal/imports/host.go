// Package imports re-exports imported wasm functions to work around an error
// compiling TinyGo to wasm "cannot use an exported function as value".
package imports

func EnableFeatures(features uint64) uint64 {
	return enableFeatures(features)
}

func GetConfig(ptr uintptr, limit uint32) (len uint32) {
	return getConfig(ptr, limit)
}

func Log(ptr uintptr, size uint32) {
	log(ptr, size)
}

func GetMethod(ptr uintptr, limit uint32) (len uint32) {
	return getMethod(ptr, limit)
}

func SetMethod(ptr uintptr, size uint32) {
	setMethod(ptr, size)
}

func GetURI(ptr uintptr, limit uint32) (len uint32) {
	return getURI(ptr, limit)
}

func SetURI(ptr uintptr, size uint32) {
	setURI(ptr, size)
}

func GetProtocolVersion(ptr uintptr, limit uint32) (len uint32) {
	return getProtocolVersion(ptr, limit)
}

func GetRequestHeaderNames(ptr uintptr, limit uint32) (len uint32) {
	return getRequestHeaderNames(ptr, limit)
}

func GetRequestHeader(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return getRequestHeader(namePtr, nameSize, bufPtr, bufLimit)
}

func GetRequestHeaders(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return getRequestHeaders(namePtr, nameSize, bufPtr, bufLimit)
}

func SetRequestHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32) {
	setRequestHeader(namePtr, nameSize, valuePtr, valueSize)
}

func ReadRequestBody(bufPtr uintptr, bufLimit uint32) (eofLen uint64) {
	return readRequestBody(bufPtr, bufLimit)
}

func WriteRequestBody(bufPtr uintptr, bufSize uint32) {
	writeRequestBody(bufPtr, bufSize)
}

func GetRequestTrailerNames(ptr uintptr, limit uint32) (len uint32) {
	return getRequestTrailerNames(ptr, limit)
}

func GetRequestTrailer(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return getRequestTrailer(namePtr, nameSize, bufPtr, bufLimit)
}

func GetRequestTrailers(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return getRequestTrailers(namePtr, nameSize, bufPtr, bufLimit)
}

func SetRequestTrailer(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32) {
	setRequestTrailer(namePtr, nameSize, valuePtr, valueSize)
}

func Next() {
	next()
}

func GetStatusCode() uint32 {
	return getStatusCode()
}

func SetStatusCode(statusCode uint32) {
	setStatusCode(statusCode)
}

func GetResponseHeaderNames(ptr uintptr, limit uint32) (len uint32) {
	return getResponseHeaderNames(ptr, limit)
}

func GetResponseHeader(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return getResponseHeader(namePtr, nameSize, bufPtr, bufLimit)
}

func GetResponseHeaders(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return getResponseHeaders(namePtr, nameSize, bufPtr, bufLimit)
}

func SetResponseHeader(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32) {
	setResponseHeader(namePtr, nameSize, valuePtr, valueSize)
}

func ReadResponseBody(bufPtr uintptr, bufLimit uint32) (eofLen uint64) {
	return readResponseBody(bufPtr, bufLimit)
}

func WriteResponseBody(bufPtr uintptr, bufSize uint32) {
	writeResponseBody(bufPtr, bufSize)
}

func GetResponseTrailerNames(ptr uintptr, limit uint32) (len uint32) {
	return getResponseTrailerNames(ptr, limit)
}

func GetResponseTrailer(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (okLen uint64) {
	return getResponseTrailer(namePtr, nameSize, bufPtr, bufLimit)
}

func GetResponseTrailers(namePtr uintptr, nameSize uint32, bufPtr uintptr, bufLimit uint32) (len uint32) {
	return getResponseTrailers(namePtr, nameSize, bufPtr, bufLimit)
}

func SetResponseTrailer(namePtr uintptr, nameSize uint32, valuePtr uintptr, valueSize uint32) {
	setResponseTrailer(namePtr, nameSize, valuePtr, valueSize)
}
